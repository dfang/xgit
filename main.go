package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/mouuff/go-rocket-update/pkg/provider"
	"github.com/mouuff/go-rocket-update/pkg/updater"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var logger *slog.Logger
var logLevel *slog.LevelVar

var (
	version   = "dev"
	commit    = "none"
	date      = "unknown"
	goVersion = "go 1.21"
	repo      = "https://github.com/dfang/xgit"
)

func main() {
	assciLogo := `
                      #       #     
                              #     
 ##   ##   ######   ###     ######  
   # #    #     #     #       #     
    #     #     #     #       #     
   # #     ######     #       #     
 ##   ##        #   #####      ### 
           #####                    
`
	fmt.Println(assciLogo)
	// fmt.Printf("%srevision %s, built with %s at %s\n", assciLogo, xgitVersion, goVersion, buildTimestamp)

	var rootCmd = &cobra.Command{Use: "xgit"}
	var cmdA = &cobra.Command{
		Use:   "self-update",
		Short: "self update",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("self update ......")
			selfUpdate()
		},
	}

	var cmdB = &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Printf("revision %s, built with %s at %s\n", xgitVersion, goVersion, buildTimestamp)
			printVersion()
		},
	}

	var cmdC = &cobra.Command{
		Use:   "clone",
		Short: "clone",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("args", args)
			processedArgs := processArgs(args)
			fmt.Println("processedArgs", processedArgs)
			execShell("git", processedArgs)
		},
	}

	rootCmd.AddCommand(cmdA, cmdB, cmdC)
	rootCmd.Flags().Bool("debug", false, "verbose mode")

	if err := rootCmd.Execute(); err != nil {
		// Error: unknown command "clone" for "xgit"
		fmt.Println("deletegate to external command: git")

		// fmt.Println(err)
		os.Exit(1)
	}

	// // If no matching subcommand is found, delegate to an external command
	// if len(os.Args) > 1 {
	// 	fmt.Println("deletegate to external command")
	// 	externalCommand := os.Args[1]
	// 	externalArgs := os.Args[2:]
	// 	cmd := exec.Command(externalCommand, externalArgs...)
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Stderr = os.Stderr
	// 	if err := cmd.Run(); err != nil {
	// 		fmt.Println("Error running external command:", err)
	// 		os.Exit(1)
	// 	}
	// }

	os.Exit(0)

	logger.Info("gitcache client")
	args := os.Args
	if logLevel.Level() == slog.LevelDebug {
		for i, arg := range args {
			logger.Debug("arg", slog.Int("index", i), slog.String("value", arg))
		}
	}

	var isClone = false
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "clone") {
			isClone = true
			break
		}
	}

	logger.Info("debug", slog.Bool("isCloneMode", isClone))

	execShell("git", args[1:])
}

func init() {
	logLevel := &slog.LevelVar{} // INFO
	opts := slog.HandlerOptions{
		Level: logLevel,
	}
	logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
}

func processArgs(args []string) []string {
	fmt.Println("args args", args)
	var isDepth = false
	var isClone = true
	// for i := 0; i < len(args); i++ {
	// 	if strings.Contains(args[i], "-vv") || strings.Contains(args[i], "-vvv") {
	// 		// you can change the level anytime like this
	// 		// logLevel.Set(slog.LevelDebug)
	// 		args = append(args[:i], args[i+1:]...)
	// 		break
	// 	}

	// 	if strings.Contains(args[i], "-version") || strings.Contains(args[i], "-v") {
	// 		printVersion()
	// 		os.Exit(0)
	// 	}
	// }

	for i := 0; i < len(args); i++ {
		if isClone {
			if strings.Contains(args[i], "depth") || strings.Contains(args[i], "no-depth") {
				isDepth = true
			}

			// support 3 types of url
			// https://github.com/<user>/<repo>
			// github.com/<user>/<repo>
			// <user>/<repo>
			if strings.Contains(args[i], "https://github.com") {
				logger.Debug("debug", slog.String("repo", args[i]))
				args[i] = strings.Replace(args[i], "https://github.com", "https://ghproxy.com/https://github.com", -1)
				logger.Debug("debug", slog.String("repo", args[i]))
			} else if !strings.Contains(args[i], "http") && strings.Contains(args[i], "github.com") {
				logger.Debug("debug", slog.String("repo", args[i]))
				args[i] = "https://ghproxy.com/https://" + args[i]
				logger.Debug("debug", slog.String("repo", args[i]))
				break
			} else if !strings.Contains(args[i], "http") && strings.Contains(args[i], "/") {
				logger.Debug("debug", slog.String("repo", args[i]))
				args[i] = "https://ghproxy.com/https://github.com/" + args[i]
				logger.Debug("debug", slog.String("repo", args[i]))
				break
			}
		}
	}

	if isClone && (!isDepth) {
		args = append(args, "--depth=1")
	}
	return args
}

func printVersion() {
	// fmt.Printf("runtime.GOOS %s\n", runtime.GOOS)
	// fmt.Printf("runtime.GOARCH %s\n", runtime.GOARCH)
	fmt.Printf("revision: %s\n", commit)
	fmt.Printf("xgit version: %s\n", version)
	fmt.Printf("built with: %s\n", goVersion)
	fmt.Printf("built at: %s\n", date)
	fmt.Printf("repo: %s\n", repo)
	fmt.Printf("xgit %s, commit %s, built at %s\n", version, commit, date)
}

func execShell(cmd string, args []string) string {
	index := 0
	args = append(args[:index+1], args[index:]...)
	args[index] = "clone"

	// Join the strings in the args slice using a space separator
	joinedString := strings.Join(args, " ")

	toExecute := fmt.Sprintf("git clone %s", joinedString)
	logger.Info(toExecute)

	var command = exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	var err = command.Start()
	if err != nil {
		return err.Error()
	}
	err = command.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}

func selfUpdate() {
	caser := cases.Title(language.English)
	u := &updater.Updater{
		Provider: &provider.Github{
			RepositoryURL: "github.com/dfang/xgit",
			ArchiveName:   fmt.Sprintf("xgit_%s_%s.tar.gz", caser.String(runtime.GOOS), "x86_64"),
		},
		ExecutableName: "xgit",
		// Version:        "v0.0.6", // You can change this value to trigger an update
		Version: version,
	}

	log.Println("Current version: " + u.Version)
	log.Println("Looking for updates...")
	var wg sync.WaitGroup
	wg.Add(1)
	// For the example we run the update in the background
	// but you could directly call u.Update()
	var updateStatus updater.UpdateStatus
	var updateErr error
	go func() {
		updateStatus, updateErr = u.Update()
		wg.Done()
	}()

	// Here you can execute your program

	wg.Wait() // Waiting for the update process to finish before exiting
	if updateErr != nil {
		log.Println(updateErr)
	}
	if updateStatus == updater.Updated {
		log.Println("Updated!")
	}
}

package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

var logger *slog.Logger

func main() {
	logLevel := &slog.LevelVar{} // INFO
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
	fmt.Printf("%srevision %s, built with %s at %s\n", assciLogo, xgitVersion, goVersion, buildTimestamp)

	opts := slog.HandlerOptions{
		Level: logLevel,
	}
	logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))

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

	var isDepth = false
	for i := 1; i < len(args); i++ {
		if strings.Contains(args[i], "-vv") || strings.Contains(args[i], "-vvv") {
			// you can change the level anytime like this
			logLevel.Set(slog.LevelDebug)
			args = append(args[:i], args[i+1:]...)
			break
		}

		if strings.Contains(args[i], "-version") || strings.Contains(args[i], "-v") {
			printVersion()
			os.Exit(0)
		}
	}

	for i := 1; i < len(args); i++ {
		if isClone {
			if strings.Contains(args[i], "depth") || strings.Contains(args[i], "no-depth") {
				isDepth = true
			}

			if strings.Contains(args[i], "https://github.com") {
				logger.Debug("debug", slog.String("repo", args[i]))
				args[i] = strings.Replace(args[i], "https://github.com", "https://ghproxy.com/https://github.com", -1)
				logger.Debug("debug", slog.String("repo", args[i]))
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

	execShell("git", args[1:])
}

func printVersion() {
	fmt.Printf("git version: %s\n", xgitVersion)
	fmt.Printf("built with: %s\n", goVersion)
	fmt.Printf("built at: %s\n", buildTimestamp)
	fmt.Printf("repo: %s\n", repo)
}

func execShell(cmd string, args []string) string {
	var argss = ""
	for i := 0; i < len(args); i++ {
		argss = argss + args[i] + " "
	}
	logger.Debug("debug", slog.String("run", cmd+" "+string(argss)))

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

package main

// var xgitVersion string
// var goVersion string
// var buildTimestamp string
// var repo string

// func init() {
// 	repo = "https://github.com/dfang/xgit"
// 	goVersion = "go 1.21.0"
// 	// currentVersion()
// 	// latestVersion()
// }

// func currentVersion() {
// 	// Run the "git describe --tags --abbrev=0" command to get the latest tag.
// 	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")

// 	// Capture the command output.
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Printf("Error running Git command: %v\n", err)
// 		return
// 	}

// 	// Trim whitespace and newline characters from the output.
// 	tag := strings.TrimSpace(string(output))

// 	xgitVersion = tag
// 	fmt.Printf("Current version: %s\n", tag)
// }

// func latestVersion() {
// 	// Replace with your GitHub personal access token.
// 	token := os.Getenv("GITHUB_AUTH_TOKEN")
// 	var client *github.Client
// 	ctx := context.Background()
// 	if token == "" {
// 		client = github.NewClient(nil)
// 	} else {
// 		// Create an authenticated GitHub client.
// 		ts := oauth2.StaticTokenSource(
// 			&oauth2.Token{AccessToken: token},
// 		)
// 		tc := oauth2.NewClient(ctx, ts)
// 		client = github.NewClient(tc)
// 	}

// 	// Specify the repository owner and name.
// 	owner := "dfang"
// 	repo := "xgit"
// 	// Get the latest release.
// 	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
// 	if err != nil {
// 		fmt.Printf("Error getting latest release: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Latest release tag: %s\n", *release.TagName)
// 	fmt.Printf("Release name: %s\n", *release.Name)
// 	fmt.Printf("Release URL: %s\n", *release.HTMLURL)
// }

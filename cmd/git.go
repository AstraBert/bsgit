package cmd

import (
	"fmt"
	"os/exec"
)

func GitInit() error {
	cmd := exec.Command("git", "init", ".")
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("\x1b[1;31mAn error occurred while initializing the Git repository: %s", err.Error())
		return err
	}
	fmt.Println("\x1b[1;92mGit repository successfully initialized")
	return nil
}

func GitBranchRename() error {
	cmd := exec.Command("git", "branch", "-M", "main")
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("\x1b[1;31mAn error occurred while renaming the default branch to main: %s", err.Error())
		return err
	}
	fmt.Println("\x1b[1;92mDefault branch succesfully renamed to main")
	return nil
}

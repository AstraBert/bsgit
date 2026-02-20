package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ownerName string
var projectName string
var projectDescription string
var showHelp bool

var rootCmd = &cobra.Command{
	Use:   "bsgit",
	Short: "bsgit bootstraps your Git repositories from scratch.",
	Long:  "bsgit is a CLI command that initializes your git repositories and creates a README, an MIT license, a minimal CONTRIBUTING.md guide for them",
	Run: func(cmd *cobra.Command, args []string) {
		if showHelp {
			_ = cmd.Help()
			return
		}
		if projectName == "" {
			fmt.Println("\x1b[1;31mYou need to provide a project name")
			os.Exit(1)
		}
		if ownerName == "" {
			fmt.Println("\x1b[1;31mYou need to provide the name of the repository owner")
			os.Exit(1)
		}
		readme := Readme{ProjectName: projectName, Description: projectDescription}
		contributing := Contributing{ProjectName: projectName}
		license := NewLicense(ownerName)
		readmeContent := BuildReadme(readme)
		contributingContent := BuildContributing(contributing)
		licenseContent := BuildLicense(license)
		err := GitInit()
		if err != nil {
			os.Exit(1)
		}
		if err := os.WriteFile("README.md", []byte(readmeContent), 0644); err != nil {
			fmt.Printf("\x1b[1;31mAn error occurred while writing README.md: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("\x1b[1;92mREADME.md successfully written")
		if err := os.WriteFile("CONTRIBUTING.md", []byte(contributingContent), 0644); err != nil {
			fmt.Printf("\x1b[1;31mAn error occurred while writing CONTRIBUTING.md: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("\x1b[1;92mCONTRIBUTING.md successfully written")
		if err := os.WriteFile("LICENSE", []byte(licenseContent), 0644); err != nil {
			fmt.Printf("\x1b[1;31mAn error occurred while writing LICENSE: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("\x1b[1;92mLICENSE successfully written")
		if err := os.WriteFile(".gitignore", make([]byte, 0), 0644); err != nil {
			fmt.Printf("\x1b[1;31mAn error occurred while writing .gitignore: %s", err.Error())
			os.Exit(1)
		}
		fmt.Println("\x1b[1;92m.gitignore successfully written")
		fmt.Println()
		fmt.Println("\x1b[1;92mGit repository bootstrapping complete")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing bsgit '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Show the help message and exit.")
	rootCmd.Flags().StringVarP(&ownerName, "owner", "o", "", "Full name of the repostitory's owner")
	rootCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
	rootCmd.Flags().StringVarP(&projectDescription, "description", "d", "", "Description of the project")

	_ = rootCmd.MarkFlagRequired("owner")
	_ = rootCmd.MarkFlagRequired("name")
}

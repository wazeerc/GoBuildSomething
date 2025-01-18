/*
Copyright © 2025 github.com/wazeerc
*/
package cmd

import (
	"fmt"
	"os"

	"GoBuildSomething/internal/config"
	"GoBuildSomething/internal/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoBuildSomething",
	Short: "A tool to quickly start a new project 🚀",
	Long: `GoBuildSomething is a minimal CLI tool built with Go that helps you to quickly scaffold a new project.

Simply provide a project name and optional path (defaults to Desktop) to get started with an already initialized Git repository.`,
	Run: runPrompts,
}

func runPrompts(cmd *cobra.Command, args []string) {
	cmd.Println("🚀 Welcome to GoBuildSomething")
	fmt.Print("\n")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	if cfg.DefaultPath == "" {
		cmd.Println("📂 First time setup: Please set your default project path")
		cfg.DefaultPath = utils.GetInput("Default Path", os.ExpandEnv("$HOME/Desktop"))
		if err := config.SaveConfig(cfg); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			os.Exit(1)
		}
	}

	projectName := utils.GetInput("Project Name", "my-project")
	path := utils.GetInput("Path", cfg.DefaultPath)

	fmt.Printf("\n⌛ Creating project '%s'\n", projectName)

	if err := utils.InitializeProject(projectName, path); err != nil {
		fmt.Printf("Error initializing project: %v\n", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

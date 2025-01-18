package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InitializeProject(projectName, basePath string) error {
	projectPath := filepath.Join(basePath, projectName)

	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	if err := os.Chdir(projectPath); err != nil {
		return fmt.Errorf("failed to change to project directory: %w", err)
	}

	if err := initGit(); err != nil {
		return err
	}

	if err := openVSCode(); err != nil {
		return err
	}

	fmt.Printf("\n\n✅ Successfully initialized project at %s\n", projectPath)
	return nil
}

func initGit() error {
	gitCmd := exec.Command("git", "init")
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}
	return nil
}

func openVSCode() error {
	codeCmd := exec.Command("code", ".")
	if err := codeCmd.Run(); err != nil {
		return fmt.Errorf("failed to open VS Code: %w", err)
	}
	return nil
}

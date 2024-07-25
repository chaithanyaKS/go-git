/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the git repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		var path string
		if len(args) > 1 {
			return errors.New("more than one arguments found")
		}
		if len(args) == 0 {
			path = ""
		} else {
			path = args[0]
		}

		return initializeRepository(path)
	},
}

func getRepoPath(filePath string) (string, error) {
	var repoPath string
	defaultRepo := "test_repo"
	gitPath := ".git"

	if filePath == "" || filePath == "." {
		curr_path, err := os.Getwd()
		if err != nil {
			return "", err
		}
		repoPath = path.Join(curr_path, defaultRepo)
	} else {
		absPath, err := filepath.Abs(filePath)
		if err != nil {
			return "", err
		}
		repoPath = absPath
	}

	return path.Join(repoPath, gitPath), nil
}

func initializeRepository(filePath string) error {
	defaultDirs := []string{"refs", "objects"}
	repoPath, err := getRepoPath(filePath)
	if err != nil {
		return err
	}
	for _, dir := range defaultDirs {
		gitPath := path.Join(repoPath, dir)
		err := os.MkdirAll(gitPath, 0777)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Initialized empty git repository %s\n", repoPath)
	return nil

}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/chaithanyaKS/go-git/internal/blob"
	"github.com/chaithanyaKS/go-git/internal/database"
	"github.com/chaithanyaKS/go-git/internal/entry"
	"github.com/chaithanyaKS/go-git/internal/tree"
	"github.com/chaithanyaKS/go-git/internal/workspace"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		commit()
		return nil
	},
}

func commit() error {
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	rootPath := path.Join(currDir, "test_repo")
	gitPath := path.Join(rootPath, ".git")
	dbPath := path.Join(gitPath, "objects")
	workSpace := workspace.Initialize(rootPath)
	db := database.New(dbPath)
	files, err := workSpace.ListFiles()
	if err != nil {
		return err
	}
	var entries []entry.Entry
	for _, file := range files {
		fmt.Println(file)
		data, err := workspace.ReadFile(file)
		if err != nil {
			return err
		}
		blobData := blob.New(data)
		err = db.Store(blobData)
		if err != nil {
			return err
		}
		newEntry := entry.New(file, blobData.Oid)
		entries = append(entries, newEntry)
	}
	tree := tree.New(entries)
	db.Store(tree)
	return nil

}

func init() {
	rootCmd.AddCommand(commitCmd)
}

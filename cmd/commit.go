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
	"strings"
	"time"

	"github.com/chaithanyaKS/go-git/internal/author"
	"github.com/chaithanyaKS/go-git/internal/blob"
	gitCommit "github.com/chaithanyaKS/go-git/internal/commit"
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
		var commitMessage string
		if len(args) > 1 {
			return errors.New("more than one arguments found")
		}
		if len(args) == 0 {
			return errors.New("must enter a commit message")
		} else {
			commitMessage = args[0]
		}
		commit(commitMessage)
		return nil
	},
}

func commit(commitMessage string) error {
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

	name := os.Getenv("GIT_AUTHOR_NAME")
	email := os.Getenv("GIT_AUTHOR_EMAIL")
	author := author.New(name, email, time.Now())
	commit := gitCommit.New(tree.Oid, author, commitMessage)
	db.Store(commit)
	headPath := filepath.Join(gitPath, "HEAD")
	file, err := os.OpenFile(headPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(commit.Oid))
	fmt.Printf("[(root-commit %s)] %s\n", commit.Oid, strings.Split(commitMessage, "\n")[0])
	return nil

}

func init() {
	rootCmd.AddCommand(commitCmd)
}

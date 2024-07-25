package workspace

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

type Workspace struct {
	Path string
}

func Initialize(pathName string) Workspace {
	return Workspace{Path: pathName}
}

func (w *Workspace) ListFiles() ([]string, error) {
	IGNORE := []string{".git"}
	var files []string
	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	currDir = path.Join(currDir, "test_repo")

	err = filepath.Walk(w.Path, func(path string, info fs.FileInfo, err error) error {
		isFileInIgnored := slices.ContainsFunc(IGNORE, func(ignoredDir string) bool {
			return strings.Contains(path, ignoredDir)
		})
		if path != currDir && !isFileInIgnored {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

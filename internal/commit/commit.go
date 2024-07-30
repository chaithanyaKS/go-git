package commit

import (
	"fmt"
	"strings"

	"github.com/chaithanyaKS/go-git/internal/author"
)

type Commit struct {
	Oid     string
	Tree    string
	Author  *author.Author
	Message string
	Type    string
}

func New(oid string, author *author.Author, message string) *Commit {
	return &Commit{Tree: oid, Author: author, Message: message, Type: "commit"}
}

func (c *Commit) AssignOid(oid string) {
	c.Oid = oid
}
func (c *Commit) GetData() (string, error) {
	lines := []string{
		fmt.Sprintf("tree %s", c.Tree),
		fmt.Sprintf("author %s", c.Author.Name),
		fmt.Sprintf("committer %s", c.Author.Name),
		"",
		c.Message,
	}

	return strings.Join(lines, "\n"), nil
}

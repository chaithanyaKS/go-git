package tree

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/chaithanyaKS/go-git/internal/entry"
)

type Tree struct {
	Entries []entry.Entry
	Type    string
}

var ENTRY_FORMAT = "Z*H40"
var MODE = "100644"

func New(entries []entry.Entry) *Tree {
	return &Tree{Entries: entries, Type: "tree"}
}

func (t *Tree) GetData() (string, error) {
	var encodedEntries []string
	sort.SliceStable(t.Entries, func(i, j int) bool {
		return t.Entries[i].Name < t.Entries[j].Name
	})
	for _, entry := range t.Entries {
		firstElem := fmt.Sprintf("%s %s \000", MODE, entry.Name)
		secondElem, err := packObjectId(entry.Oid)
		if err != nil {
			return "", err
		}
		combinedEntry := fmt.Sprintf("%s%s", firstElem, secondElem)
		encodedEntries = append(encodedEntries, combinedEntry)
	}
	return strings.Join(encodedEntries, ""), nil
}

func packObjectId(oid string) (string, error) {
	byteArray := make([]byte, 20)

	for i := 0; i < len(oid); i += 2 {
		byteValue, err := strconv.ParseUint(oid[i:i+2], 16, 8)
		if err != nil {
			return "", err
		}
		byteArray[i/2] = byte(byteValue)
	}
	return string(byteArray), nil
}

func (t *Tree) AssignOid(oid string) {}
func (t *Tree) Len(oid string) int {
	return len(t.Entries)
}

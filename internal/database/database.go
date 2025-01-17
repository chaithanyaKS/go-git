package database

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
)

type ObjectWriter interface {
	GetData() (string, error)
	AssignOid(string)
}

type Database struct {
	Path string
	Type string
}

func New(path string) Database {
	return Database{Path: path, Type: "blob"}
}

func (db *Database) Store(blob ObjectWriter) error {
	data, err := blob.GetData()
	if err != nil {
		return err
	}
	content := fmt.Sprintf("%s %d%s%s", db.Type, len(data), "\000", data)
	objectId := createObjectId(content)
	blob.AssignOid(objectId)
	return db.writeObject(objectId, content)
}

func (db *Database) writeObject(objectId string, content string) error {
	objectPath := path.Join(db.Path, objectId[0:2], objectId[2:])
	dirname := filepath.Dir(objectPath)
	tempPath := path.Join(dirname, generateTempName(6))

	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		os.Mkdir(dirname, 0777)
	} else if err != nil {
		return err
	}

	file, err := os.OpenFile(tempPath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	writer := zlib.NewWriter(file)
	writer.Write([]byte(content))
	writer.Close()

	return os.Rename(tempPath, objectPath)

}

func generateTempName(size uint) string {
	const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	b := make([]byte, size)
	for i := range b {
		b[i] = CHARSET[rng.Intn(len(CHARSET))]
	}

	return "temp_obj_" + string(b)
}

func createObjectId(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

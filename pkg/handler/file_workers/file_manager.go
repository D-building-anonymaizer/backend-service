package files

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SplitFileName(fileName string) (string, string) {
	_, file := filepath.Split(fileName)
	ext := filepath.Ext(file)
	name := strings.TrimSuffix(file, ext)
	name += time.Now().String()
	return RemoveCyrillic(name), ext
}

func RemoveCyrillic(a string) string {
	hash := md5.Sum([]byte(a))
	return hex.EncodeToString(hash[:])
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

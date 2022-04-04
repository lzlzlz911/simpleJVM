package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			fc, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer fc.Close()
			data, err := ioutil.ReadAll(fc)
			if err != nil {
				panic(err)
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

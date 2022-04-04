package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove *
	CompositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			CompositeEntry = append(CompositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return CompositeEntry
}

package file

import "os"

// Exists returns whether the given file or directory exists or not.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

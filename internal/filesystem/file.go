package filesystem

import (
	"os"
	"path/filepath"
)

func EnsureFile(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
	}
	return nil
}

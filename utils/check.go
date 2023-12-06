package utils

import (
	"errors"
	"os"
)

/*
 * Check, checks if the .git directory exists in the specified directory.
 */

func Check(dir string) error {
	if _, err := os.Stat(dir + ".git/"); os.IsNotExist(err) {
		return errors.New(".git not found")
	}

	return nil
}

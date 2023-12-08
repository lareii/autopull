package utils

import (
	"errors"
	"os"
)

/*
 * Check is a function to check if the ".git/" directory
 * exists in the given directory.
 */

func Check(dir string) error {
	if _, err := os.Stat(dir + ".git/"); os.IsNotExist(err) {
		return errors.New(".git not found")
	}

	return nil
}

package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

/*
 * Repository is a function to get the repository URL from
 * the "./git/config" file in the specified directory.
 */

func Repository(dir string) (string, error) {
	config := dir + ".git/config"

	file, err := os.Open(config)
	if err != nil {
		return "", errors.New("git config not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var repo string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "url =") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				repo = strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("git config cannot read")
	}

	return repo, nil
}

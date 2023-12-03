package src

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func verifySignature(secret, header, payload string) error {
	// X-Hub-Signature-256: sha256=xxx
	header = strings.Split(header, "=")[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))

	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(header)) {
		return errors.New("request signatures didn't match")
	}

	return nil
}

func pullRepo(dir string) error {
	config := dir + ".git/config"

	file, err := os.Open(config)
	if err != nil {
		return errors.New("git config not found")
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
		return errors.New("git config cannot read")
	}

	cmd := exec.Command("git", "pull", repo)
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

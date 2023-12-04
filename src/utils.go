package src

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
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
	repo, err := GetRepoURL(dir)

	if err != nil {
		return errors.New(err.Error())
	}

	cmd := exec.Command("git", "pull", repo)
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func GetRepoURL(dir string) (string, error) {
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

func CheckDir(dir string) error {
	if _, err := os.Stat(dir + ".git/"); os.IsNotExist(err) {
		return errors.New(".git not found")
	}

	return nil
}

func Log(level int, msg string, newLine bool) {
	var c *color.Color

	if level == 0 {
		c = color.New(color.BgRed)
		c.Print("ERROR")
	} else if level == 1 {
		c = color.New(color.BgYellow)
		c.Print("WARNING")
	} else {
		c = color.New(color.BgGreen)
		c.Print("LOG")
	}

	print(" ")

	c = color.New(color.BgBlue)
	now := time.Now()
	c.Print(now.Format("2006/01/02 15:04:05")) // wtf

	print(" ")

	print(msg)

	if newLine {
		println()
	}
}

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

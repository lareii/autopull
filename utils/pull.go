package utils

import (
	"errors"
	"os/exec"
)

/*
 * Pull is a function for executing the "git pull" command.
 */

func Pull(dir string) error {
	repo, err := Repository(dir)

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

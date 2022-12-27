package entry

import (
	"os/exec"
)

func CreateEntry(journal, tags, body string) error {
	cmd := exec.Command("dayone2", "new", "-t", tags, "-j", journal, body)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

package edjournal

import (
	"os/user"
	"path/filepath"
)

const JOURNAL_DEFAULT_PATH = "~\\Saved Games\\Frontier Developments\\Elite Dangerous"

type Entry struct {
	timestamp string
	event string
}

func Parse(line string) (*Entry, error) {
	return &Entry{}, nil
}


func GetDefaultJournalPath() (string, error) {
	return expand(JOURNAL_DEFAULT_PATH)
}


func expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}
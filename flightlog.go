package main

import (
	"io/ioutil"
	"os"
	"strings"
)

const JOURNAL_DEFAULT_PATH = "~\\Saved Games\\Frontier Developments\\Elite Dangerous"

type FlightLog struct {
	latest_log_file *os.FileInfo
	last_entry    string
	path          string
}

type Entry struct {
	event string
	timestamp string
}

func parse(line string) (*Entry, error) {
	return &Entry{}, nil
}

func getDefaultFlightLogPath() string {
	path, err := expand(JOURNAL_DEFAULT_PATH)
	if err != nil {
		return ""
	}
	return path
}

func (fl *FlightLog) isUpdated() bool {
	return true
}

func (fl *FlightLog) updateLatestFile() *os.FileInfo {
	flight_logs, err := ioutil.ReadDir(fl.path)
	if err != nil {
		return nil
	}

	for i := len(flight_logs) - 1; i > 0; i -- {
		flight_log_file := flight_logs[i]
		if flight_log_file.IsDir() || !strings.HasPrefix(flight_log_file.Name(), "Journal.") {
			continue
		}
		if fl.latest_log_file == nil {
			fl.latest_log_file = &flight_log_file
			continue
		}
		if flight_log_file.ModTime().Before((*fl.latest_log_file).ModTime()) {
			break
		}
	}

	return fl.latest_log_file

}

func (fl *FlightLog) parseLogFile(filename string) *[]Entry {
	return &[]Entry{}
}
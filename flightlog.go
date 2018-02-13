package main

import (
	"io/ioutil"
	"os"
	"strings"
	"encoding/json"
	"path/filepath"
	"github.com/sirupsen/logrus"
)


const JOURNAL_DEFAULT_PATH = "~/Documents/ED"

type FlightLog struct {
	latest_log_file string
	last_line       string
	path            string
}

type Entry struct {
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
	Line      string
	FileName  string
}

func (e *Entry) Equals(entry *Entry) bool {
	return e.Line == entry.Line && e.FileName == e.FileName
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

func (fl *FlightLog) getLatestFile() string {
	flight_logs, err := ioutil.ReadDir(fl.path)
	if err != nil {
		return ""
	}

	var latest_log_file os.FileInfo

	for i := len(flight_logs) - 1; i > 0; i -- {
		flight_log_file := flight_logs[i]
		if flight_log_file.IsDir() || !strings.HasPrefix(flight_log_file.Name(), "Journal.") {
			continue
		}
		if latest_log_file == nil {
			latest_log_file = flight_log_file
			continue
		}
		if flight_log_file.ModTime().After(latest_log_file.ModTime()) {
			latest_log_file = flight_log_file
			continue
		} else {
			break
		}
	}

	if latest_log_file != nil {
		return filepath.Join(fl.path, latest_log_file.Name())
	} else {
		return ""
	}
}

func (fl *FlightLog) getFlightLogLines(filename string) *[]string {
	return get_lines(filename)
}

func (fl *FlightLog) parseFlightLogLines(lines *[]string, filename string) []Entry {
	var entries []Entry
	logger := Logger.WithField("filename", filepath.Base(filename))

	for _, line := range *lines {
		entry := fl.parseFlightLogLine(line, filename)
		if entry == nil {
			continue
		}
		entries = append(entries, *entry)
	}

	logger.Debugln("Parsed entries:", len(entries))
	return entries
}

func (fl *FlightLog) parseFlightLogLine(line string, filename string) *Entry {
	var e Entry
	if len(line) == 0 {
		return nil
	}

	logger := Logger.WithFields(logrus.Fields{
		"line": line,
		"filename": filepath.Base(filename),
	})
	err := json.Unmarshal([]byte(line), &e)
	if err != nil {
		logger.Debugln("Cannot parse line:", err)
		return nil
	}
	e.Line = line
	e.FileName = filename
	return &e
}


func (fl *FlightLog) getLinesToSend(lines *[]string) *[]string {
	linesToSend := []string{}

	if len(*lines) == 0 {
		return &[]string{}
	}
	last_line := (*lines)[len(*lines) - 1]
	if fl.last_line == "" {
		fl.last_line = last_line
		return &[]string{}
	}
	if fl.last_line == last_line {
		return &[]string{}
	}

	for i:= len(*lines) - 1; i > 0; i-- {
		line := (*lines)[i]
		if fl.last_line == line {
			break
		}
		linesToSend = append([]string{line}, linesToSend...)
	}

	fl.last_line = last_line

	return &linesToSend
}
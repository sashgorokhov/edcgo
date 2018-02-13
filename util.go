package main

import (
	"os/user"
	"path/filepath"
	"crypto/rand"
	"bufio"
	"os"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"strings"
)



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


func get_lines(filename string) *[]string {
	logger := Logger.WithField("filename", filepath.Base(filename))
	file, err := os.Open(filename)
	if err != nil {
		logger.Infoln("Error opening file:", err)
		return &[]string{}
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	logger.Debugln("Read lines:", len(lines))
	return &lines
}


func randToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}


func tokenToQRCode(token string) ([]byte, error) {
	return qrcode.Encode(token, qrcode.Medium, 256)
}
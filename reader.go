package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	FREQUENCY = "%f"
	CALL      = "%c"
	DATE      = "%d"
	TIME      = "%t"
	BAND      = "%b"
	MODE      = "%m"
)

func readInputFiles(files StringArray) (cs []Contact, err error) {
	cs = make([]Contact, 0)
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			return cs, err
		}
		reader := bufio.NewReader(f)
		readStructure(reader)
	}
	return cs, nil
}

func readStructure(reader *bufio.Reader) ([]string, error) {
	line := ""
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		trimmedLine := strings.TrimSpace(string(l))
		if trimmedLine != "" {
			line = strings.TrimSpace(trimmedLine[1:])
			break
		}
	}

	verbs := strings.Split(line, "\t")
	return verbs, nil
}

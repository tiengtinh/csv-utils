package main

import (
	"encoding/csv"
	"io"
	"os"
)

type record []string

func read(filepath string, comma rune, skipHead bool) (rec chan record, errChan chan error) {
	var file *os.File
	rec = make(chan record)
	errChan = make(chan error)

	go func() {
		var line int = 0
		var err error
		var lineData record

		defer close(rec)

		if file, err = os.Open(filepath); err != nil {
			errChan <- err
			return
		}

		defer file.Close()

		reader := csv.NewReader(file)
		reader.Comma = comma

		for {
			lineData, err = reader.Read()

			if err == io.EOF {
				break
			}

			defer func() {
				line++
			}()

			if err != nil {
				errChan <- err
				continue
			}

			if skipHead && line == 0 {
				continue
			}

			rec <- lineData
		}

	}()

	return rec, errChan
}

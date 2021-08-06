//
// main.go
// Copyright (C) 2021 forseason <me@forseason.vip>
//
// Distributed under terms of the MIT license.
//

package env

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

var (
	data          map[string]string
	isInitialized bool
)

func init() {
	data = make(map[string]string)
	f, err := os.Open(".env")
	if err != nil {
		log.Println("env: " + err.Error())
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if len(line) == 0 {
			continue
		}
		err = parseLine(line)
		if err != nil {
			log.Println("env: " + err.Error())
			data = make(map[string]string)
			return
		}
	}
	isInitialized = true
	log.Println("env: loaded.")
}

func parseLine(line []byte) error {
	var (
		key        string
		value      string
		i          int
		isFiltered bool
	)
	if line[0] == '#' {
		return nil
	}
	for i < len(line) {
		if !isFiltered && line[i] == '=' {
			isFiltered = true
			break
		}
		i++
	}
	if !isFiltered {
		return errors.New("invalid format")
	}
	key = string(line[:i])
	if i == len(line)-1 {
		value = ""
	} else {
		value = string(line[i+1:])
	}
	if _, ok := data[key]; ok {
		return errors.New("reduplicated keys")
	}
	data[key] = value
	return nil
}

// This function will try to get value from map with input key.
// If the map has not been initialized or there is no value for the key,
// Get() will return default value.
func Get(key, defaultValue string) string {
	if value, ok := data[key]; !ok || !isInitialized {
		return defaultValue
	} else {
		return value
	}
}

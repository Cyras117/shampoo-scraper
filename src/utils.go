package scraper

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"
)

// check if a string is inside another.
func isIn(phrase string, str string) bool {

	p := strings.ToLower(phrase)
	s := strings.ToLower(str)

	lPhr := len(phrase)
	lStr := len(str)

	if lStr < lPhr {
		return false
	}

	for i := 0; i < lStr; i++ {
		if p[0] == s[i] {
			for j := i; j < lStr; j++ {
				if s[j] == p[j-i] {
					if j-i == lPhr-1 {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

// Log errors on console//TODO apagar talvez
func errLogOutput(err error) {
	if err != nil {
		log.Output(5, err.Error())
	}
}

func CheckIfFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func DeleFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckLastModified(fileName string) time.Time {

	fileInfo, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return fileInfo.ModTime()
}

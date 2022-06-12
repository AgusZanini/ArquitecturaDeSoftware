package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return bytes, err

}

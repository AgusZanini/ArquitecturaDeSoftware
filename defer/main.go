package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if err := CreateFile("example.txt"); err != nil {
		panic(err)
	}

	if err := WriteFile("example.txt", []byte("This is an example file")); err != nil {
		panic(err)
	}

	bytes, err := ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Content of file: %s", string(bytes)))
}

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func WriteFile(path string, content []byte) error {
	if err := os.WriteFile(path, content, os.ModeAppend); err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		println("Err")
	}
	// 1回目
	err = file_scan(f)
	if err != nil {
		println("Err")
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		println("Err")
	}
	// 2回目
	err = file_scan(f)
	if err != nil {
		println("Err")
	}
}

func file_scan(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err := scanner.Err()
	return err
}

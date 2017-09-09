package main

import "fmt"

func main() {
	err := Hoge(2)
	if err != nil {
		fmt.Printf("fmt.Errorf: %#v\n", err)
	}
}

func Hoge(num int) error {
	if num%2 == 0 {
		return fmt.Errorf("Failed: %s")
	}
	return nil
}

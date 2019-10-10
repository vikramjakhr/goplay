package main

import (
	"os"
	"fmt"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error ", err)
	}
	defer file.Close()

	if _, err = file.Write([]byte("one")); err != nil {
		fmt.Println("write error", err)
	}

}

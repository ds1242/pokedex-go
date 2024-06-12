package main

import (
	"os"
	"fmt"
)

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
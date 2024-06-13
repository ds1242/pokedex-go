package pokemap

import (
	"os"
	"fmt"
)

func commandExit() error {
	fmt.Println()
	fmt.Println("Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}
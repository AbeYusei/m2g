package error

import (
	"fmt"
	"os"
)

// Handle handle error
func Handle(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

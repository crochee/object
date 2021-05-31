package main

import (
	"fmt"
	"os"

	"github.com/crochee/object/cmd"
)

func main() {
	if err := cmd.Server(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

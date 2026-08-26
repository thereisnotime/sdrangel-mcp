package main

import (
	"os"

	"github.com/thereisnotime/sdrangel-mcp/internal/version"
)

func main() {
	if err := newCLI(version.Version).Execute(); err != nil {
		os.Exit(1)
	}
}

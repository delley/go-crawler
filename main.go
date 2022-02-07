package main

import (
	"log"
	"os"

	_ "github.com/delley/go-crawler/matchers"
	"github.com/delley/go-crawler/search"
)

func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

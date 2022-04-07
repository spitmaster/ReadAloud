package main

import (
	"os"
	"read-aloud/internal/view"
)

func main() {
	view := view.CreateView()
	if err := view.Start(); err != nil {
		os.Exit(1)
	}
}

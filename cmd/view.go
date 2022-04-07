package main

import (
	"read-aloud/internal/view/biz"
)

func main() {
	app := biz.FirstCreate()
	app.Start()
}

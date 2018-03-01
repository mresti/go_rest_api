package main

import (
	"app"
)

func main() {
	app := app.App{}
	app.Initialize()
	app.Run()
}

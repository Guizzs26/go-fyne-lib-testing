package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Window")

	firstName := widget.NewEntry()
	lastName := widget.NewEntry()

	fmt.Println(firstName)
	fmt.Println(lastName)

	window.ShowAndRun()
}

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
Responsável por:

- Criar a janela principal
- Criar os campos de input (com o widget do Fyne)
- Adicionar o botão "Gerar excel"
- Organizar esses componentes com layouts do container do Fyne
- Exibir a janela
*/
func CreateWindow() {
	app := app.New()
	mainWindow := app.NewWindow("Excel generator")

	mainWindow.Resize(fyne.NewSize(400, 200))

	firstNameEntry := widget.NewEntry()
	firstNameEntry.SetPlaceHolder("Enter your first name")

	lastNameEntry := widget.NewEntry()
	lastNameEntry.SetPlaceHolder("Enter your last name")

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("Enter your age")

	birthdayEntry := widget.NewEntry()
	birthdayEntry.SetPlaceHolder("Enter your birthday date (dd/mm/aaaa)")

	randomFloatValueEntry := widget.NewEntry()
	randomFloatValueEntry.SetPlaceHolder("Enter any decimal value")

	generateExcelButton := widget.NewButton("Generate Excel Spreadsheet", func() {

	})

	content := container.NewVBox(
		widget.NewLabel("Fill in the data to generate the Excel"),
		firstNameEntry,
		lastNameEntry,
		ageEntry,
		birthdayEntry,
		randomFloatValueEntry,
		generateExcelButton,
	)

	mainWindow.SetContent(content)

	mainWindow.ShowAndRun()
}

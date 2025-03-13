package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Guizzs26/go-fyne-lib-testing/internal/excel"
)

/*
Responsável por:

- Criar os campos de input (com o widget do Fyne)
- Adicionar o botão "Gerar excel" que chama a função responsável por gerar o Excel
*/
func createInputFields() (*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Button) {
	firstNameInput := widget.NewEntry()
	firstNameInput.SetPlaceHolder("Enter your first name")

	lastNameInput := widget.NewEntry()
	lastNameInput.SetPlaceHolder("Enter your last name")

	ageInput := widget.NewEntry()
	ageInput.SetPlaceHolder("Enter your age")

	birthdayInput := widget.NewEntry()
	birthdayInput.SetPlaceHolder("Enter your birthday date (dd/mm/aaaa)")

	randomFloatValueInput := widget.NewEntry()
	randomFloatValueInput.SetPlaceHolder("Enter any decimal value")

	generateExcelButton := widget.NewButton("Generate Excel Spreadsheet", func() {
		firstName := firstNameInput.Text
		lastName := lastNameInput.Text
		ageStr := ageInput.Text
		birthdayStr := birthdayInput.Text
		randomFloatStr := randomFloatValueInput.Text

		excel.GenerateExcel(firstName, lastName, ageStr, birthdayStr, randomFloatStr)
	})

	return firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatValueInput, generateExcelButton
}

/*
Responsável por:

  - Organizar esses componentes o container do Fyne
*/
func createLayout(firstName, lastName, age, birthday, randomFloat *widget.Entry, button *widget.Button) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Fill in the data to generate the Excel"),
		firstName,
		lastName,
		age,
		birthday,
		randomFloat,
		button,
	)
}

/*
Responsável por:

  - Pela manipulação da janela principal (criação, definição da janela como principal, centralizar, modificar tamanho)
  - Setar o conteúdo
  - Exibir a janela e inicar o loop do Fyne
*/
func CreateWindow() {
	app := app.New()
	mainWindow := app.NewWindow("Excel generator")
	mainWindow.SetMaster()
	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(400, 200))

	firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatValueInput, generateExcelButton := createInputFields()
	content := createLayout(firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatValueInput, generateExcelButton)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

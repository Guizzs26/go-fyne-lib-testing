package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()                       // Cria uma nova aplicação do Fyne. Como se fosse uma nova instância do Fyne.
	window := app.NewWindow("Hello World") // Cria uma nova janela com o TÍTULO "Hello World".

	window.SetContent(widget.NewLabel("Hello World")) // Define o conteúdo dessa janela através de uma label.

	window.ShowAndRun() // Exibe a janela e roda a aplicação

	/*
		Também podemos fazer da seguinte forma:

		window.Show()
		app.Run()


		- É importante ressaltar que apenas podemos ter um único "Run()" na aplicação. Essas aplicações de GUI precisam executar uma
		espécie de "event loop" que fica "escutando" por interações do usuário com a aplicação.
		- Essas funções Run() e ShowAndRun() são as funções responsáveis por iniciar esse loop de eventos.
		- Ela deve ser chamada no final do código de configuração da aplicação Fyne, normalmente na main().

		- Qualquer código após a função Run() ou ShowAndRun() só será executado APÓS o fechamento da aplicação. Podemos usar o app.Quit() para fechar
		aplicações Desktop (não funciona para Android ou IOS).
		- Se todas as janelas forem fechadas, o programa se encerra automaticamente.
	*/
}

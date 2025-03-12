package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time 03:04:05")
	clock.SetText(formatted)
}

func main() {
	app := app.New()                 // Cria uma nova aplicação do Fyne. Como se fosse uma nova instância do Fyne.
	window := app.NewWindow("Clock") // Cria uma nova janela com o TÍTULO "Hello World".
	window.SetMaster()               // Podemos definir uma janela principal para nossa aplicação. Podemos ter uma janela principal e várias secundárias.
	// Se a "master" for fechada, toda aplicação será encerrada automaticamente.
	window.Resize(fyne.NewSize(400, 400)) // Modifica o tamanho da janela.

	// Aqui estamos aprendendo a atualizar um valor. Precisamos passar uma label (por exemplo) vazia e atribuir para uma variável, pois o retorno dessa função é uma referência
	// à label criada, então podemos atualiza-la depois.
	// formatted := time.Now().Format("Time 03:04:05")
	// clock.SetText(formatted)
	clock := widget.NewLabel("")
	window.SetContent(clock)
	updateTime(clock)

	// window.SetContent(widget.NewLabel("Hello World")) // Define o conteúdo dessa janela através de uma label fixa.

	// IIFE -> Immediately Invoked Function Expression, é uma função anônima que é automaticamente criada e executada. Usamos uma goroutine também para executar o código de forma assíncrona.
	// Aqui que mora a magia da execução em segundo plano de forma concorrente (pararelo) !!!
	go func() {
		// Isso é basicamente um for infinito que fica atualizando o valor de clock a cada segundo.
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	window.Show()

	// Criação de uma segunda janela. Como explicado anteriormente, precisamos agora chamar o Show() individualmente para cada uma.
	window2 := app.NewWindow("Second window")
	window2.SetContent(widget.NewLabel("Second - More content!"))
	window2.Show()

	// Podemos criar janelas de forma dinâmica:
	window2.SetContent(widget.NewButton("Open new", func() {
		window3 := app.NewWindow("Third window")
		window3.SetContent(widget.NewLabel("Third - More content!!!"))
		window3.Show()
	}))

	app.Run()
	// window.ShowAndRun() // Exibe a janela e roda a aplicação
	/*
		Também podemos fazer da seguinte forma:

		window.Show()
		app.Run()


		- É importante ressaltar que apenas podemos ter um único "Run()" na aplicação. Essas aplicações de GUI precisam executar uma
			espécie de "event loop" que fica "escutando" por interações do usuário com a aplicação.
		- Essas funções Run() e ShowAndRun() são as funções responsáveis por iniciar esse loop de eventos.
		- Ela deve ser chamada no final do código de configuração da aplicação Fyne, normalmente na main().

		- A principal diferença entre o ShowAndRun() e o Run() é que a primeira função é usada quando você tem APENAS UMA JANELA PRINCIPAL (inicia o loop e a janela)
			e a segunda é usada quando você tem várias janelas (apenas inicia o loop, cada janela deve ser chamada manualmente com a função Show()).

		- Qualquer código após a função Run() ou ShowAndRun() só será executado APÓS o fechamento da aplicação. Podemos usar o app.Quit() para fechar
			aplicações Desktop (não funciona para Android ou IOS).
		- Se todas as janelas forem fechadas, o programa se encerra automaticamente.

	*/
}

# 🧮 Excel Generator

O **Excel Generator** é uma aplicação **PROTÓTIPO** desenvolvida em **Golang** com a ferramenta de desenvolvimento dekstop **Fyne** que permite ao usuário gerar planilhas Excel (.xlsx) a partir de entradas de dados. 
Se a planilha já existir, novos dados serão adicionados a ela, sem sobrescrever os dados anteriores.

Não foi implementada nenhuma melhoria na interface gráfica, como layout, estilo e afins.

O projeto foi pensando para ter uma boa organização e separação de responsabilidades entre as funções.

### Funcionalidades:
- Interface gráfica simples e intuitiva.
- Validação de dados de entrada (nomes, idade, data de nascimento, etc.).
- Criação de novas planilhas ou adição de dados em planilhas existentes.

## 📋 Pré-requisitos

Certifique-se de ter instalado em sua máquina:

- [Golang](https://golang.org/dl/) (1.21 ou superior)
- [Fyne](https://developer.fyne.io/) - Siga o tutorial no "Getting Started".

## 🚀 Instalação e Execução

1. **Clone o repositório:**

```bash
git clone https://github.com/SeuUsuario/excel-generator.git](https://github.com/Guizzs26/go-fyne-lib-testing.git)
```

2. **Acesse o diretório do projeto:**

```bash
cd go-fyne-lib-testing
```

3. **Instale as dependências do Fyne:**

```bash
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
```

4. **Compile a aplicação:**

```bash  
go build -o ./bin/excel-generator.exe .\cmd\app\main.go (.exe é necessário no windows).
```

```bash  
go build -o "C:\Users\SeuUsuario\Desktop\excel-generator.exe" .\cmd\app\main.go (Exemplo de geraçao em um diretório qualquer, no caso, a área de trabalho).
```

5. **Execute a aplicação:**

Apenas um exemplo de como executar a aplicação, mas você deve executar no caminho onde o .exe foi salvo.

No Windows:

```bash
./excel-generator.exe
```

No Linux/macOS:

```bash
./excel-generator
```

## 🔨 Para Desenvolvedores

1. Execute a aplicação durante o desenvolvimento:

```bash
go run .\cmd\app\main.go
```

2. Gerar um executável otimizado:

```bash
go build -ldflags "-s -w" -o ./bin/excel-generator.exe .\cmd\app\main.go
```

## 🤝 Contribuição

Sinta-se à vontade para enviar Pull Requests ou abrir Issues para relatar bugs ou sugerir melhorias.

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.


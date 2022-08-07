package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	exibirIntroducao()
	for {
		exibirMenu()

		comando := lerComando()
		// Utilizando Switch - Case
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não encontrado")
			os.Exit(-1) //Biblioteca os
		}
	}
}

func exibirIntroducao() {
	nome, _ := os.Hostname()
	versao := 1.1
	fmt.Println("Hostname:", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// Utilizando Slice
	sites := []string{"https://hub.docker.com", "https://id.heroku.com",
		"https://www.travis-ci.com/"}
	// Utilizando For
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second) //Biblioteca time
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, _ := http.Get(site) //Biblioteca net/http

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!!!",
			resp.StatusCode)
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:",
			resp.StatusCode)
	}
}

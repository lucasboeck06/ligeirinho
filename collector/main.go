package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao puxar variáveis da .env %v", err)
	}

	apiToken := os.Getenv("URBS_API_TOKEN")
	apiUrl := os.Getenv("API_URL")

	linha := "512"

	//Requisição GET básica
	resp, err := http.Get(apiUrl + linha + "&c=" + apiToken) //Shape do Itamarati
	if err != nil {
		log.Fatalf("Erro de rede/conexão %v\n", err)
	}

	defer resp.Body.Close()

	//Lê o que recebemos do endpoint e armazena em um slice de bytes, "body"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler os dados do corpo %v\n", err)
	}
	//Converte "body" em string
	bodyString := string(body)

	switch resp.StatusCode {
	case http.StatusOK:
		if strings.Contains(bodyString, "ERRO 404") {
			log.Fatal("Erro lógico. Servidor retornou 200, mas temos uma página 404 como resposta")
		} else {
			fmt.Println(bodyString)
		}

	case http.StatusNotFound:
		log.Fatalf("Erro HTTP 404, Corpo: %s", bodyString)

	default:
		log.Fatalf("Erro HTTP inesperado %d. Corpo: %s", resp.StatusCode, bodyString)
	}
}

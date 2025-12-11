package main

import (
	"strconv"
	"strings"
)

// Struct padrão para os pontos do shape
type pontoBruto struct {
	SHP string `json:"SHP"`
	LAT string `json:"LAT"`
	LON string `json:"LON"`
	COD string `json:"COD"`
}

type ponto struct {
	SHP string
	LAT float64
	LON float64
}

func converteCoord(coord string) (float64, error) {
	//Muda a vírgula pelo ponto
	convertido := strings.Replace(coord, ",", ".", 1) //(coord, "o valor a ser substuído", "o quero colocar no lugar", quantidade máxima de alterações na string)
	//Converte a striing em um float64 já retornando
	return strconv.ParseFloat(convertido, 64)
}

func normaliza(linha []pontoBruto) []ponto {

	// Cria slice VAZIO mas com o TAMANHO PREVISTO
	// clean.len = 0
	// clean.cap = len(raw)
	clean := make([]ponto, 0, len(linha))

	for _, r := range linha {

		//Conversão com base na func de conversão
		lat, _ := converteCoord(r.LAT)
		lon, _ := converteCoord(r.LON)

		//Cria a versão final da struct tipo ponto
		p := ponto{
			SHP: r.SHP,
			LAT: lat,
			LON: lon,
		}

		//Adiciona o novo p ao slice clean, recriando o slice "limpo"
		clean = append(clean, p)
	}

	return clean
}

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

// Struct padrão para os pontos do shape no formato de dado correto
type ponto struct {
	SHP int
	LAT float64
	LON float64
}

// Troca a "," por "." e coverte a coordenada para float
func converteCoord(coord string) (float64, error) {
	//Muda a vírgula pelo ponto
	convertido := strings.Replace(coord, ",", ".", 1) //(coord, "o valor a ser substuído", "o quero colocar no lugar", quantidade máxima de alterações na string)
	//Converte a striing em um float64 já retornando
	return strconv.ParseFloat(convertido, 64)
}

// Usa a func converteCoord para criar o slice de structs limpo (com tipos corretos)
func normaliza(linha []pontoBruto) []ponto {

	// Cria slice VAZIO mas com o TAMANHO PREVISTO
	// clean.len = 0
	// clean.cap = len(raw)
	clean := make([]ponto, 0, len(linha))

	for _, r := range linha {

		//Conversão com base na func de conversão
		lat, _ := converteCoord(r.LAT)
		lon, _ := converteCoord(r.LON)

		//Converte o SHP para inteiro
		shp, _ := strconv.Atoi(r.SHP)

		//Cria a versão final da struct tipo ponto
		p := ponto{
			SHP: shp,
			LAT: lat,
			LON: lon,
		}

		//Adiciona o novo p ao slice clean, recriando o slice "limpo"
		clean = append(clean, p)
	}

	return clean
}

func deleteShp(sliceOriginal []ponto, pontoDel int) []ponto {

	i := 0

	for r := range sliceOriginal

}

// // Função que transforma o slice de structs em um map de slices de structs (ordenação por SHP)
// func agruparPorSHP(clean []ponto) map[string][]ponto {
// 	//Cria mapa de string : slice de structs (do tipo ponto)
// 	shapePorSHP := make(map[string][]ponto)

// 	//Itera sob clean
// 	//pontoAtual é a struct (ponto) da vez
// 	//Cagamos para o índice
// 	for _, pontoAtual := range clean {
// 		//codigoDoShape recebe o SHP da struct que está na vez
// 		codigoDoShape := pontoAtual.SHP

// 		//O novo mapa concatena a string de acordo com a chave, caso não exista, ele cria a chave e atribui o valor
// 		//(com a chave: SHP da vez), recebe uma concatenação do valor da struct da vez {SHP, LAT, LON}
// 		//Sim, ele usa o SHP como chave, mas ele se repete na struct [SHP{[]ponto}SHP{[]ponto}...]
// 		shapePorSHP[codigoDoShape] = append(shapePorSHP[codigoDoShape], pontoAtual)
// 	}

// 	return shapePorSHP
// }

// // Função que calcula a distância entre os ponto geográficos
// // Recebe duas structs, do tipo ponto (uma de início e outra de fim)
// func dist(a, b ponto) float64 {
// 	//Diferença das duas latitudes
// 	latDiff := a.LAT - b.LAT
// 	//Diferença das duas longitudes
// 	lonDiff := a.LON - b.LON

// 	//Retorna raiz de(latDiff ao quadrado + lonDiff ao quadrado)
// 	return math.Sqrt(latDiff*latDiff + lonDiff*lonDiff)

// 	//Utilizando o teorema de pitágoras, calculamos a hipotenusa
// 	//Assim temos a distância euclidiana (descosideramos a fórmula esférica da Terra)
// }

// func ordenaSHP(shapes map[string][]ponto) []string {
// 	//Cria um slice de strings com o tamanho exato == o número de SHPs existentes len(shapes)
// 	shpIDs := make([]string, 0, len(shapes))

// 	//Quando passado somente um valor para o range de um mapa, ele retorna apenas a chave e ignora a correspondência
// 	//Itera sobre shapes, puxando somente as strings que correspondem aos SHPs
// 	for shp := range shapes {
// 		//Concatena no slice de strings um novo SHP
// 		//Vai ser um slicec de strings, só com o SHPs
// 		shpIDs = append(shpIDs, shp)
// 	}

// 	//Cria um map para a confirmação de utilização
// 	usados := make(map[string]bool)
// 	//Cria o slice que terá a ordem definitiva dos SHPs
// 	ordem := []string{}
// 	//Pega o primeiro elemento do shpIDs e guarda (uma string)
// 	atual := shpIDs[0]
// 	//Insere a string do "atual" no slice de "ordem"
// 	ordem = append(ordem, atual)
// 	//Maca o "atual", que já foi utilizado, como true dentro do map "usados"
// 	usados[atual] = true

// 	//Loop que prossegue até que o número de itens em "ordem" seja igual ao número de chaves de "shapes"
// 	for len(ordem) < len(shapes) {
// 		//ultimoPonto recebe o valor da última Struct do slice que é referente a chave "atual"
// 		ultimoPonto := shapes[atual][len(shapes[atual])-1]

// 		menorDist := math.MaxFloat64
// 		proximo := ""

// 		//Percorr o map shapes e retorna a "chave" em shp
// 		for shp := range shapes {

// 			//Se o shp atual == false (Não estar no check de "usados"), ele continua esse for
// 			//Se satisfazer o if (== true), ele killa o for
// 			if usados[shp] {
// 				//Go imediatamente ignora o resto desse for e volta ao anterior
// 				continue
// 			}

// 			//Considera o valor SHP vindo da iteração
// 			//Para selecinoar a "chave" (shp da vez), e o [0] define que é o primeiro correspondente
// 			primeiroPonto := shapes[shp][0]

// 			d := dist(ultimoPonto, primeiroPonto)

// 			if d < menorDist {
// 				menorDist = d
// 				proximo = shp
// 			}

// 		}

// 		usados[proximo] = true
// 		ordem = append(ordem, proximo)
// 		atual = proximo
// 	}

// 	return ordem

// }

package main

import(
	"os"
	"fmt"
	"bufio"
	"encoding/csv"
	"flag"
	"time"
	"strings"
	"strconv"
	"sort"
    "io/ioutil"
	"./revista"
)

func readCSVFile(fileName string, fildNumber int)(rawData[][]string, err error){
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("Erro de I/O")
		os.Exit(1)
	}
	
	defer file.Close()
	
	reader := csv.NewReader(file)
	
	reader.Comma = ';'
	reader.FieldsPerRecord = fildNumber
	reader.TrimLeadingSpace = true
	
	rawData, err = reader.ReadAll()
	
	check(err)
	
	rawData = rawData[1:]
	return rawData, nil
}

func parseData(dataStr string)(time.Time, error){
	dataSplit := strings.Split(dataStr, "/")

	//http://golang.org/pkg/time/#Parse
	const dataLayout = "2006-Jan-02"

	data := dataSplit[2]

	switch dataSplit[1] {
	case "01":
		data = data + "-Jan-"
	case "02":
		data = data + "-Feb-"
	case "03":
		data = data + "-Mar-"
	case "04":
		data = data + "-Apr-"
	case "05":
		data = data + "-May-"
	case "06":
		data = data + "-Jun-"
	case "07":
		data = data + "-Jul-"
	case "08":
		data = data + "-Aug-"
	case "09":
		data = data + "-Sep-"
	case "10":
		data = data + "-Oct-"
	case "11":
		data = data + "-Nov-"
	case "12":
		data = data + "-Dec-"
	}

	data = data + dataSplit[0]

	return time.Parse(dataLayout, data)
}

//checa se variável de erro é diferente de nula, pra qualquer ocasião
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func escreverArquivo(nomeArquivo, conteudo string) {
	d1 := []byte(conteudo)
    err := ioutil.WriteFile(nomeArquivo, d1, 0644)
    check(err)
}

func relatorioRevisores(revisores []revista.Revisor) string {
	var retorno string

	retorno += "Revisor;Qtd. artigos revisados;Média das notas atribuídas\n"

	for _, c := range revisores {
		if(c.IsEnvolvido()){
			retorno += c.RelatorioRevisor()
			retorno += "\n"
		}
	}

	return retorno
}

func main() {
	edFileName := flag.String("e", "edicao.txt", "Nome do arquivo contendo os dados da edição")
	temFileName := flag.String("t", "temas.csv", "Nome do arquivo contendo os dados dos temas")
	pesFileName := flag.String("p", "pessoas.csv", "Nome do arquivo contendo os dados das pessoas")
	artFileName := flag.String("a", "artigos.csv", "Nome do arquivo contendo os dados dos artigos")
	revFileName := flag.String("r", "revisoes.csv", "Nome do arquivo contendo os dados das revisões")
	flag.Parse()
	
	var edicao revista.Edicao
	temas := make(map[int]string)
	autores := make([]revista.Autor, 0)
	codAutores := make(map[int]int)
	revisores := make([]revista.Revisor, 0)
	codRevisores := make(map[int]int)
	
	rawPesCSVData, _ := readCSVFile(*pesFileName, 7)
	
	for _,pessoa := range rawPesCSVData{
		cod, _ := strconv.ParseInt(strings.Trim(pessoa[0], " "), 10, 0)
		senha, _ := strconv.ParseInt(strings.Trim(pessoa[3], " "), 10, 0)
		if pessoa[6] == "A"{
			codAutores[int(cod)] = len(autores)
			autores = append(autores, revista.CriarAutor(pessoa[1], pessoa[2], int(senha), pessoa[4], pessoa[5])) 
		}else{
			codRevisores[int(cod)] = len(revisores)
			revisores = append(revisores, revista.CriarRevisor(pessoa[1], pessoa[2], int(senha), pessoa[4], pessoa[5]))
		}
	}
	
	rawTemCSVData, _ := readCSVFile(*temFileName, 3)
	
	for _,tema := range rawTemCSVData{
		index,_ := strconv.ParseInt(tema[0], 10, 0)
		temas[int(index)] = tema[1]
		revHabilitados := strings.Split(tema[2], ",")
		for _, v := range revHabilitados{
			cod,_ := strconv.ParseInt(strings.Trim(v, " "), 10, 0)
			revisores[codRevisores[int(cod)]].AddTema(tema[1])
		}
	}
	
	edFile, err := os.Open(*edFileName)
	
	if err != nil{
		fmt.Println("Erro de I/O")
		return
	}
	
	defer edFile.Close()
	
	edReader := bufio.NewReader(edFile)
	tema ,_ := edReader.ReadString('\n')
	tema = strings.Trim(tema, "\n")
	chefe,_ := edReader.ReadString('\n')
	chefe = strings.Trim(chefe, "\n")
	volStr,_ := edReader.ReadString('\n')
	volStr = strings.Trim(volStr, "\n")
	numStr,_ := edReader.ReadString('\n')
	numStr = strings.Trim(numStr, "\n")
	dataStr,_ := edReader.ReadString('\n')
	dataStr = strings.Trim(dataStr, "\n")
	data, _ := parseData(dataStr)
	var revChefe revista.Revisor
	for _, r := range revisores{
		if r.GetNome() == chefe{
			revChefe = r
		}
	}
	num, _ := strconv.ParseInt(strings.Trim(numStr, " "), 10, 0)
	vol, _ := strconv.ParseInt(strings.Trim(volStr, " "), 10, 0)
	edicao = revista.CriarEdicao(int(num), int(vol), data, revChefe)
	edicao.SetTema(tema)
	
	rawArtCSVData, _ := readCSVFile(*artFileName, 4)
	
	for _,artigo := range rawArtCSVData{
		autoresArtigo := strings.Split(artigo[2], ",")
		var codContato int64
		if len(autoresArtigo) == 1 {
			codContato, _ = strconv.ParseInt(strings.Trim(autoresArtigo[0], " "), 10, 0)
		}else{
			codContato, _ = strconv.ParseInt(strings.Trim(artigo[3], " "), 10, 0)
		}
		
		contato := autores[codAutores[int(codContato)]]
		
		art := revista.CriarArtigo(artigo[1], contato)
		
		for _, v := range autoresArtigo{
			codAutor,_ := strconv.ParseInt(strings.Trim(v, " "), 10, 0)
			art.AdicionaAutor(autores[codAutores[int(codAutor)]])
		}
		codArtigo, _ := strconv.ParseInt(strings.Trim(artigo[0], " "), 10, 0)
		edicao.SubmeterArtigo(art, int(codArtigo))
	}
	
	rawRevCSVData, _ := readCSVFile(*revFileName, 5)
	
	for _, revisao := range rawRevCSVData{
		cod, _ := strconv.ParseInt(strings.Trim(revisao[0], " "), 10, 0)
		artigo := edicao.GetArtigo(int(cod))
		cod, _ = strconv.ParseInt(strings.Trim(revisao[1], " "), 10, 0)
		revisor := &revisores[codRevisores[int(cod)]]
		notas := revisao[2:]
		var media float64
		media = 0
		for _, nota := range notas{
			nota = strings.Trim(strings.Replace(nota, ",", ".", 1), " ")
			aux, _ := strconv.ParseFloat(nota, 64)
			media += aux
		}
		media /= 3
		
		artigo.AdicionaRevisao(media, revisor)
		
	}
	
	//revOrdenados := revisoresOrdenados(revisores)
	
	sort.Sort(revista.ByName(revisores))
	
	//escreve as saidas em arquivos muito lindos
	escreverArquivo("relat-resumo.txt", edicao.Resumo(revisores))
	escreverArquivo("relat-revisoes.csv", edicao.RelatorioRevisoes())
	escreverArquivo("relat-revisores.csv", relatorioRevisores(revisores))

	
}

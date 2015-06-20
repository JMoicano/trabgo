package main

import(
	"os"
	"fmt"
	"encoding/csv"
	"flag"
	"./revista"
)

func main() {
	edFileName := flag.String("e", "edicao.txt", "Nome do arquivo contendo os dados da edição")
	temFileName := flag.String("t", "temas.csv", "Nome do arquivo contendo os dados dos temas")
	pesFileName := flag.String("p", "pessoas.csv", "Nome do arquivo contendo os dados das pessoas")
	artFileName := flag.String("a", "artigos.csv", "Nome do arquivo contendo os dados dos artigos")
	revFileName := flag.String("r", "edicao.txt", "Nome do arquivo contendo os dados das revisões")
	flag.Parse()
	
	var chefe string
	var edicao Edicao
	var temas map[int]string
	var autores map[int]Autor
	var revisores map[int]Revisor
	
	pesFile, err := os.Open(*pesFileName)
	if err != nil{
		fmt.Println("Erro de I/O")
		return
	}
	
	defer pesFile.Close()
	
	pesReader := csv.NewReader(pesFile)
	
	pesReader.Comma = ';'
	pesReader.FieldsPerRecord = 7
	pesReader.TrimLeadingSpace = true
	
	rawPesCSVData, err := pesReader.ReadAll()
	
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	rawPesCSVData = rawPesCSVData[1:]
	
	for _,pessoa := range rawPesCSVData{
		if pessoa[6] == A{
			autores[pessoa[0]] = CriarAutor(pessoa[1], pessoa[2], pessoa[3], pessoa[4], pessoa[5])
		}else{
			revisores[pessoa[0]] = CriarRevisor(pessoa[1], pessoa[2], pessoa[3], pessoa[4], pessoa[5])
		}
	}
	
	temFile, err := os.Open(*temFileName)
	if err != nil{
		fmt.Println("Erro de I/O")
		return
	}
	
	defer temFile.Close()
	
	temReader := csv.NewReader(temFile)
	
	temReader.Comma = ';'
	temReader.FieldsPerRecord = 3
	temReader.TrimLeadingSpace = true
	
	rawTemCSVData, err := temReader.ReadAll()
	
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	rawTemCSVData = rawTemCSVData[1:]
	
	for _,tema := range rawTemCSVData{
		index,_ := strconv.ParseInt(tema[0], 10, 0)
		temas[int(index)] = tema[1]
		revHabilitados := strings.Split(tema[2], ",")
		for _, v := range revHabilitados{
			cod,_ := strconv.ParseInt(strings.Trim(v, " "), 10, 0)
			revisores[int(cod)].AddTema(tema[1])
		}
	}
	
	edFile, err := os.Open(*edFileName)
	edReader := bufio.NewReader(edFile)
	tema ,_ := edReader.ReadString('\n')
	chefe,_ = edReader.ReadString('\n')
	vol,_ = edReader.ReadString('\n')
	num,_ = edReader.ReadString('\n')
	dataStr,_ = edReader.ReadString('\n') 
	dataSplit := strings.Split(dataStr, "/")//TODO: pegar a data a partir do formato da data
	
	fmt.Println(rawEdData)
	
	if err != nil{
		fmt.Println("Erro de I/O")
		return
	}
	
	defer edFile.Close()

	//fmt.Println(e)
}
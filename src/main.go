package main

import(
	"os"
	"fmt"
	"encoding/csv"
	"flag"
	"time"
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
	
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	
	rawData = rawData[1:]
	return rawData, nil
}

func parseData(dataStr string)(Time, error){
	dataSplit := strings.Split(dataStr, "/")//TODO: pegar a data a partir do formato da data

	//http://golang.org/pkg/time/#Parse
	const dataLayout = "29-Jan-1992"
	data := dataSplit[2]

	switch dataSplit[1] {
	case "01":
		data = data + "Jan"
	case "02":
		data = data + "Feb"
	case "03":
		data = data + "Mar"
	case "04":
		data = data + "Apr"
	case "05":
		data = data + "May"
	case "06":
		data = data + "Jun"
	case "07":
		data = data + "Jul"
	case "08":
		data = data + "Aug"
	case "09":
		data = data + "Sep"
	case "10":
		data = data + "Oct"
	case "11":
		data = data + "Nov"
	case "12":
		data = data + "Dec"
	}

	data = data + dataSplit[0]

	return time.ParseInLocation(dataLayout, data, &localLoc)
}

func main() {
	var edFileName, temFileName, pesFileName, artFileName, revFileName string
	edFileName = flag.String("e", "edicao.txt", "Nome do arquivo contendo os dados da edição")
	temFileName = flag.String("t", "temas.csv", "Nome do arquivo contendo os dados dos temas")
	pesFileName = flag.String("p", "pessoas.csv", "Nome do arquivo contendo os dados das pessoas")
	artFileName = flag.String("a", "artigos.csv", "Nome do arquivo contendo os dados dos artigos")
	revFileName = flag.String("r", "edicao.txt", "Nome do arquivo contendo os dados das revisões")
	flag.Parse()
	
	var edicao Edicao
	var temas map[int]string
	var autores map[int]*Autor
	var revisores map[int]*Revisor
	
	rawPesCSVData, _ := readCSVFile(pesFileName, 7)
	
	for _,pessoa := range rawPesCSVData{
		if pessoa[6] == A{
			autores[pessoa[0]] = CriarAutor(pessoa[1], pessoa[2], pessoa[3], pessoa[4], pessoa[5])
		}else{
			revisores[pessoa[0]] = CriarRevisor(pessoa[1], pessoa[2], pessoa[3], pessoa[4], pessoa[5])
		}
	}
	
	rawTemCSVData, _ := readCSVFile(temFileName, 3)
	
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
	
	if err != nil{
		fmt.Println("Erro de I/O")
		return
	}
	
	defer edFile.Close()
	
	edReader := bufio.NewReader(edFile)
	tema ,_ := edReader.ReadString('\n')
	chefe,_ := edReader.ReadString('\n')
	vol,_ = edReader.ReadString('\n')
	num,_ = edReader.ReadString('\n')
	dataStr,_ = edReader.ReadString('\n') 
	data, _ := parseData(dataStr)
	var revChefe *Revisor
	for _, r := range revisores{
		if r == chefe{
			revChefe = r
		}
	}
	edicao = CriarEdicao(num, vol, data, revChefe)
	
	rawArtCSVData, _ := readCSVFile(artFileName, 4)
	
	for _,artigo := range rawArtCSVData{
		autoresArtigo := strings.Split(artigo[2], ",")
		var codContato int
		if len(autoresArtigo) == 1 {
			codContato = strconv.ParseInt(strings.Trim(autoresArtigo[0], " "), 10, 0)
		}else{
			codContato = strconv.ParseInt(strings.Trim(artigo[3], " "), 10, 0)
		}
		
		contato := autores[int(codContato)]
		
		art := CriarArtigo(artigo[1], contato)
		
		for _, v := range autoresArtigo{
			codAutor,_ := strconv.ParseInt(strings.Trim(v, " "), 10, 0)
			art.AddAutor(autores[int(codAutor))
		}
		codArtigo, _ := strconv.ParseInt(strings.Trim(artigo[0], " "), 10, 0)
		edicao.SubmeterArtigo(art, codArtigo)
	}
	
	rawRevCSVData, _ := readCSVFile(revFileName, 5)
	
	for _, revisao := range rawRevCSVData{
		cod := strconv.ParseInt(strings.Trim(revisao[0], " "), 10, 0)
		artigo := edicao.GetArtigo(int(cod))
		cod = strconv.ParseInt(strings.Trim(revisao[1], " "), 10, 0)
		revisor := revisores[ind(cod)]
		notas := artigo[2:]
		var media float64
		media = 0
		for _, nota := range notas{
			nota = strings.Trim(strings.Replace(nota, ",", ".", 1), " ")
			media += strconv.ParseFloat(nota, 64)
		}
		media /= 3;
		artigo.AdicionaRevisao(media, revisor)
	}
	
}
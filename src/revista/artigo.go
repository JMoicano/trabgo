package revista

import (
	"sort"
	"bytes"
	"strconv"
)

type Artigo struct{
	titulo string
	contato Autor
	listaAutores[] Autor
	listaRevisores[] Revisor
	media float64
	revisoesEnviadas int
}

func CriarArtigo(titulo string, contato Autor) Artigo{
	return Artigo{titulo, contato, make([]Autor, 0), make([]Revisor, 0), 0, 0}
}

func (art *Artigo) AdicionaAutor(autor Autor){
	art.listaAutores = append(art.listaAutores, autor)
}

func (art *Artigo) AdicionaRevisao(media float64, revisor *Revisor){
	art.listaRevisores = append(art.listaRevisores, *revisor)
	revisor.AdicionaRevisao(media)
	art.media += media
	art.revisoesEnviadas ++
	if art.revisoesEnviadas == 3{
		art.media /= 3
	}
}

func (art Artigo) RelatorioRevisoes() string{
	var buffer bytes.Buffer
	sort.Sort(ByName(art.listaRevisores))
	buffer.WriteString(art.titulo)
	buffer.WriteString(";")
	buffer.WriteString(art.contato.ToString())
	buffer.WriteString(";")
	buffer.WriteString(strconv.FormatFloat(art.media, 'f', 2, 64))
	
	for _, c := range art.listaRevisores {
		buffer.WriteString(";")
		buffer.WriteString(c.GetNome())
	}

	return buffer.String()
}

func (art Artigo) GetRevisoesEnviadas() int{
	return art.revisoesEnviadas
}

func (art Artigo) GetMedia() float64{
	return art.media
}

func (art Artigo) GetTituloArtigo() string{
	return art.titulo
}

type ByMedia []Artigo

func (art ByMedia) Len() int { return len(art) }
func (art ByMedia) Less(i, j int) bool { if art[i].media == art[j].media { return art[i].titulo < art[j].titulo }else{ return art[i].media > art[j].media } }
func (art ByMedia) Swap(i, j int) { art[i], art[j] = art[j], art[i] }

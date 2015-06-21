package revista

import (
	"bytes"
	"strconv"
)

type Artigo struct{
	titulo string
	contato *Autor
	listaAutores[] *Autor
	listaRevisores[] *Revisor
	media float64
	revisoesEnviadas int
}

func CriarArtigo(titulo string, contato *Autor) *Artigo{
	var a *Artigo
	a.titulo = titulo
	a.contato = contato
	a.AdicionaAutor(contato)
	a.media = 0
	a.revisoesEnviadas = 0
	return a
}

func (art *Artigo) AdicionaAutor(autor *Autor){
	art.listaAutores = append(art.listaAutores, autor)
}

func (art *Artigo) AdicionaRevisao(media float64, revisor *Revisor){
	art.listaRevisores = append(art.listaRevisores, revisor)
	//mudei isso, conferir
	(*revisor) = AdicionaRevisao((*revisor), media)
	art.media += media
	art.revisoesEnviadas ++
	if art.revisoesEnviadas == 3{
		art.media /= 3
	}
}

func (art Artigo) relatorioRevisoes() {
	var buffer bytes.Buffer
	//TODO implementar sorting dos revisors
	//sort.Sort(art.listaRevisores)
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

func (art Artigo) Len() int {
    return len(art)
}

func (art Artigo) Less(i, j int) bool {
    return art[i].media < art[j].media;
}

func (art Artigo) Swap(i, j int) {
    art[i], art[j] = art[j], art[i]
}

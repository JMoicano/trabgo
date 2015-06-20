package revista

import(
	"sort"
)

type Artigo struct{
	titulo string
	contato *Autor
	listaAutores[] *Autor
	listaRevisores[] *Revisor
	media float64
	revisoesEnviadas int
}

func CriarArtigo(titulo string, contato *Autor) Artigo{
	var a Artigo
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
	(*revisor) = AdicionaRevisao((*revisor), media)
	art.media += media
	art.revisoesEnviadas ++
	if art.revisoesEnviadas == 3{
		art.media /= 3
	}
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

type Artigos []Artigo

func (art Artigos) Len() int {
    return len(art)
}

func (art Artigos) Less(i, j int) bool {
    return art[i].media < art[j].media;
}

func (art Artigos) Swap(i, j int) {
    art[i], art[j] = art[j], art[i]
}

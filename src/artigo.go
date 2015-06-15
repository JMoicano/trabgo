package artigo

import (
	"fmt"
//	"autor"
)

type Artigo struct{
	titulo string
//	contato Autor
//	listaAutores[] Autor
//	listaRevisores[] Revisor
	media float64
	revisoesEnviadas int
}

/*func (art *Artigo) AdicionaAutor(autor *Autor){
	art.listaAutores.append(autor);
}

func (art *Artigo) AdicionaRevisao(media float64, autor *Autor){
	
}*/

func (art Artigo) GetRevisoesEnviadas() int{
	return revisoesEnviadas
}

func (art Artigo) GetMedia() float64{
	return media
}

func (art Artigo) GetTituloArtigo() string{
	return titulo
}

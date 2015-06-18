package edicao

import(
	"fmt"
	"time"
	"revisor"
	"artigo"
)

type Edicao struct {
	volume, numero int
	dataPublicacao Time
	tema string
	chefe Revisor
	consistencia string
	artigos := map[int]Artigo{}

}

func CriarEdicao(numero int, volume int, data Time) Edicao {
	return (Edicao{volume, numero, data, "", nil, ""})
}

func GetTema(e Edicao) sting {
	return e.tema
}

func SetTema(e Edicao, tema string) Edicao{
	e.tema = tema
	return e
}

//nao entendi
// func SubmeterArtigo(a Artigo, cod int) {
// }

func SetChefe(e Edicao, r Revisor) Edicao {
	e.chefe = r
	return r
}

func GetArtigo(e Edicao, cod int){
	return e.artigos[cod]
}

//http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/
func RelatorioRevisoes(e Edicao) string {
	revisoes string
	artigos := []Artigo{}

	for _, v := range e.artigos {
		e.artigos = append(e.artigos, v)
	}

	fmt.Sort(artigos)

	for _, j range artigos {
		revisoes = revisoes + j.RelatorioRevisoes() + "\n"
	}

	return revisoes

}
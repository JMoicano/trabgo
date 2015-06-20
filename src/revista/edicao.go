package revista

import(
	"time"
	"sort"
)

type Edicao struct {
	volume, numero int
	dataPublicacao time.Time
	tema string
	chefe Revisor
	consistencia string
	artigos map[int]Artigo
}

func CriarEdicao(numero int, volume int, data time.Time, r Revisor) Edicao {
	return (Edicao{volume, numero, data, "", r, "", nil})
}

func GetTema(e Edicao) string {
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
	return e
}

func GetArtigo(e Edicao, cod int) Artigo{
	return e.artigos[cod]
}

//http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/
func RelatorioRevisoes(e Edicao) string {
	var revisoes string
	artigos := []Artigo{}

	for _, v := range e.artigos {
		artigos = append(artigos, v)
	}

	//artigo tem que implementar metodos da interface sort
	//mas da erro na hora de chamar isso aqui
	sort.Reverse(artigos)

	for _, j := range artigos {
		revisoes = revisoes + j.RelatorioRevisoes() + "\n"
	}

	return revisoes

}

func resumo
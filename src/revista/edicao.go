package revista

import(
	"time"
	"sort"
	"bytes"
	"strconv"
	"strings"
)

type Edicao struct {
	volume, numero int
	dataPublicacao time.Time
	tema string
	chefe Revisor
	consistencia string
	artigos map[int]Artigo
}

func (ed *Edicao) CriarEdicao(numero int, volume int, data time.Time, r Revisor) Edicao {
	return (Edicao{volume, numero, data, "", r, "", nil})
}

func (ed *Edicao) GetTema(e Edicao) string {
	return e.tema
}

func (ed *Edicao) SetTema(e Edicao, tema string) Edicao{
	e.tema = tema
	return e
}

//nao entendi
// func SubmeterArtigo(a Artigo, cod int) {
// }

func (ed *Edicao) SetChefe(e Edicao, r Revisor) Edicao {
	e.chefe = r
	return e
}

func (ed *Edicao) GetArtigo(e Edicao, cod int) Artigo{
	return e.artigos[cod]
}

//http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/
func (ed *Edicao) RelatorioRevisoes(e Edicao) string {
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

func (ed *Edicao) resumo(e Edicao, revisores map[int]Revisor) string {
	var buffer bytes.Buffer
	var resumo string
	var revisoresCapacitados int
	var revisoresEnvolvidos int

	var dateOut string = e.dataPublicacao.Format("ANSIC")

	buffer.WriteString("EngeSort, num. ")
	buffer.WriteString(strconv.Itoa(e.numero))
	buffer.WriteString(", volume ")
	buffer.WriteString(strconv.Itoa(e.volume))
	buffer.WriteString(" - ")
	buffer.WriteString(dateOut)
	buffer.WriteString("\n")
	buffer.WriteString("Tema: ")
	buffer.WriteString(e.tema)
	buffer.WriteString("\n")
	buffer.WriteString("Editor-chefe: ")
	buffer.WriteString(e.chefe.GetNome(e.chefe))
	buffer.WriteString("\n")
	buffer.WriteString("\n")

	for _, m := range revisores {
		for _, t := range m.temas {
			if(strings.EqualFold(e.tema, t)){
				revisoresCapacitados++
				if(m.IsEnvolvido(m)){
					revisoresEnvolvidos++
				}
				break
			}
		}
	}

	var media int

	for _, i := range e.artigos {
		media = media + i.GetRevisoesEnviadas()
	}

	media /= revisoresEnvolvidos

	buffer.WriteString("\n")
	buffer.WriteString("Artigos submetidos: ")
	buffer.WriteString(strconv.Itoa(len(e.artigos)))
	buffer.WriteString("\n")
	buffer.WriteString("Revisores Capacitados: ")
	buffer.WriteString(strconv.Itoa(revisoresCapacitados))
	buffer.WriteString("\n")
	buffer.WriteString("Revisores Envolvidos: ")
	buffer.WriteString(strconv.Itoa(revisoresEnvolvidos))
	buffer.WriteString("\n")
	buffer.WriteString("Media artigos/revisor: ")
	buffer.WriteString(strconv.Itoa(media))
	buffer.WriteString("\n")

	resumo = buffer.String()

	return resumo
}
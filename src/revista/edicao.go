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

func (ed Edicao) GetTema() string {
	return ed.tema
}

func (ed *Edicao) SetTema(tema string){
	ed.tema = tema
}


func (ed *Edicao) SubmeterArtigo(a Artigo, cod int) {
	ed.artigos[cod] = a
}

func (ed *Edicao) SetChefe(r Revisor) {
	ed.chefe = r
}

func (ed Edicao) GetArtigo(cod int) Artigo{
	return ed.artigos[cod]
}

//http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/
func (ed *Edicao) RelatorioRevisoes() string {
	var revisoes string
	artigos := []Artigo{}

	for _, v := range ed.artigos {
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

func (ed *Edicao) resumo(revisores map[int]Revisor) string {
	var buffer bytes.Buffer
	var resumo string
	var revisoresCapacitados int
	var revisoresEnvolvidos int

	var dateOut string = ed.dataPublicacao.Format("ANSIC")

	buffer.WriteString("EngeSort, num. ")
	buffer.WriteString(strconv.Itoa(ed.numero))
	buffer.WriteString(", volume ")
	buffer.WriteString(strconv.Itoa(ed.volume))
	buffer.WriteString(" - ")
	buffer.WriteString(dateOut)
	buffer.WriteString("\n")
	buffer.WriteString("Tema: ")
	buffer.WriteString(ed.tema)
	buffer.WriteString("\n")
	buffer.WriteString("Editor-chefe: ")
	buffer.WriteString(ed.chefe.GetNome())
	buffer.WriteString("\n")
	buffer.WriteString("\n")

	for _, m := range revisores {
		for _, t := range m.temas {
			if(strings.EqualFold(ed.tema, t)){
				revisoresCapacitados++
				if(m.IsEnvolvido()){
					revisoresEnvolvidos++
				}
				break
			}
		}
	}

	var media int

	for _, i := range ed.artigos {
		media = media + i.GetRevisoesEnviadas()
	}

	media /= revisoresEnvolvidos

	buffer.WriteString("\n")
	buffer.WriteString("Artigos submetidos: ")
	buffer.WriteString(strconv.Itoa(len(ed.artigos)))
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
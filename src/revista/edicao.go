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
	artigos []Artigo
	codArtigos map[int]int
}

func CriarEdicao(numero int, volume int, data time.Time, r Revisor) Edicao{
	return Edicao{volume, numero, data, "", r, make([]Artigo, 0), make(map[int]int)}
}

func (ed Edicao) GetTema() string {
	return ed.tema
}

func (ed *Edicao) SetTema(tema string){
	ed.tema = tema
}


func (ed *Edicao) SubmeterArtigo(a Artigo, cod int) {
	ed.codArtigos[cod] = len(ed.artigos)
	ed.artigos = append(ed.artigos, a)
}

func (ed *Edicao) SetChefe(r Revisor) {
	ed.chefe = r
}

func (ed Edicao) GetArtigo(cod int) *Artigo{
	return &ed.artigos[ed.codArtigos[cod]]
}

//http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/
func (ed *Edicao) RelatorioRevisoes() string {
	var revisoes string
	artigos := []Artigo{}

	revisoes = "Artigo;Autor de contato;Média das avaliações;Revisor 1; Revisor 2; Revisor 3\n"

	for _, v := range ed.artigos {
		artigos = append(artigos, v)
	}


	sort.Sort(ByMedia(artigos))

	for _, j := range artigos {
		revisoes = revisoes + j.RelatorioRevisoes() + "\n"
	}

	return revisoes

}

func gerarData(dataPublicacao time.Time) string {
	aux := dataPublicacao.Month().String()

	switch aux {
	case "January":
		aux = "Janeiro"
	case "February":
		aux = "Fevereiro"
	case "March":
		aux = "Março"
	case "April":
		aux = "Abril"
	case "May":
		aux = "Maio"
	case "June":
		aux = "Junho"
	case "July":
		aux = "Julho"
	case "August":
		aux = "Agosto"
	case "September":
		aux = "Setembro"
	case "October":
		aux = "Outubro"
	case "November":
		aux = "Novembro"
	case "December":
		aux = "Dezembro"
	}

	return aux + " de " + strconv.Itoa(dataPublicacao.Year())
}

func (ed *Edicao) Resumo(revisores []Revisor) string {
	var buffer bytes.Buffer
	var resumo string
	var revisoresCapacitados int
	var revisoresEnvolvidos int
	var media float64

	dateOut := gerarData(ed.dataPublicacao)

	buffer.WriteString("EngeSoft, num. ")
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
			if(ed.tema == t){
				revisoresCapacitados++
				if(m.IsEnvolvido()){
					revisoresEnvolvidos++
				}
				break
			}
		}
	}
	
	for _, i := range ed.artigos {
		media = media + float64(i.GetRevisoesEnviadas())
	}

	media /= float64(revisoresEnvolvidos)

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
	buffer.WriteString(strings.Replace(strconv.FormatFloat(media, 'f', 2, 64), ".", ",", -1))
	buffer.WriteString("\n")

	resumo = buffer.String()

	return resumo
}
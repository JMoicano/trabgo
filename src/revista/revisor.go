package revista

import (
	"strings"
	"strconv"
	"bytes"
)

type Revisor struct {
	nome string 
	email string
	senha int
	instituicao string
	endereco string
	temas []string
	artigosRevisados int
	notasAtribuidas float64
}

func (rev *Revisor) CriarRevisor(nome string, email string, senha int, instituicao string, endereco string) Revisor {
	return (Revisor{nome, email, senha, instituicao, endereco, nil, 0, 0.0})
}

func (rev *Revisor) GetNome(s Revisor) string {
	return s.nome
}

func (rev *Revisor) AddTema(r Revisor, tema string) {
	r.temas = append(r.temas, tema)
}

func (rev *Revisor) AdicionaRevisao(r Revisor, media float64) Revisor{
	r.notasAtribuidas = r.notasAtribuidas + media
	r.artigosRevisados++

	return r
}

func (rev *Revisor) IsEnvolvido(r Revisor) bool {
	return r.artigosRevisados > 0
}

func (rev *Revisor) IsApto(r Revisor, tema string) bool {
	for _, v := range r.temas {
		if(strings.EqualFold(tema, v)){
			return true
		}
	}

	return false
}

func (rev *Revisor) RelatorioRevisor(r Revisor) string {
	var buffer bytes.Buffer

	media := r.notasAtribuidas/float64(r.artigosRevisados)

	//essa conversao recebe o float, o 'f' eu nao sei, o 2 é numero de casas depois 
	//do ponto, e o ultimo arg é pra dizer se é float 32 ou float 64
	media_string := strconv.FormatFloat(media, 'f', 2, 64)

	buffer.WriteString(r.nome)
	buffer.WriteString(";")
	buffer.WriteString(strconv.Itoa(r.artigosRevisados))
	buffer.WriteString(";")
	buffer.WriteString(media_string)

	return buffer.String()
}

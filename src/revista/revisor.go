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

func (rev Revisor) GetNome() string {
	return rev.nome
}

func (rev *Revisor) AddTema(tema string) {
	rev.temas = append(rev.temas, tema)
}

func (rev *Revisor) AdicionaRevisao(media float64) {
	rev.notasAtribuidas = rev.notasAtribuidas + media
	rev.artigosRevisados++
}

func (rev Revisor) IsEnvolvido(r Revisor) bool {
	return r.artigosRevisados > 0
}

func (rev Revisor) IsApto(tema string) bool {
	for _, v := range rev.temas {
		if(strings.EqualFold(tema, v)){
			return true
		}
	}

	return false
}

func (rev *Revisor) RelatorioRevisor() string {
	var buffer bytes.Buffer

	media := rev.notasAtribuidas/float64(rev.artigosRevisados)

	//essa conversao recebe o float, o 'f' eu nao sei, o 2 é numero de casas depois 
	//do ponto, e o ultimo arg é pra dizer se é float 32 ou float 64
	media_string := strconv.FormatFloat(media, 'f', 2, 64)

	buffer.WriteString(rev.nome)
	buffer.WriteString(";")
	buffer.WriteString(strconv.Itoa(rev.artigosRevisados))
	buffer.WriteString(";")
	buffer.WriteString(media_string)

	return buffer.String()
}

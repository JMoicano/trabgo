package main

import (
	"fmt"
	"strings"
	//"strconv"
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

func criarRevisor(nome string, email string, senha int, instituicao string, endereco string) Revisor {
	return (Revisor{nome, email, senha, instituicao, endereco, nil, 0, 0.0})
}

func toString(s Revisor) string {
	return s.nome
}

func addTema(r Revisor, tema string) {
	r.temas = append(r.temas, tema)
}

func adicionaRevisao(r Revisor, media float64) Revisor{
	r.notasAtribuidas = r.notasAtribuidas + media
	r.artigosRevisados++

	return r
}

func isEnvolvido(r Revisor) bool {
	return r.artigosRevisados > 0
}

func isApto(r Revisor, tema string) bool {
	for _, v := range r.temas {
		if(strings.EqualFold(tema, v)){
			return true
		}
	}

	return false
}

// func relatorioRevisor(r Revisor) string {
// 	media := r.notasAtribuidas/float64(r.artigosRevisados)
// 	temp := r.nome + ";" + string(r.artigosRevisados) + ";" + strconv.FormatFloat(media, 'E', 2, 0)

// 	return temp
// }

func main(){

	r := criarRevisor("Cesar", "cesar@email", 123, "ufes", "endereco")
	r = adicionaRevisao(r, 9)

	fmt.Println(r)
}

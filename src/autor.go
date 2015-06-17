package autor

type Autor struct {
	nome string 
	email string
	senha int
	instituicao string
	endereco string
}

func criarAutor(nome string, email string, senha int, instituicao string, endereco string) Autor {
	return (Autor{nome, email, senha, instituicao, endereco})
}

func toString(s Autor) string {
	return s.nome
}

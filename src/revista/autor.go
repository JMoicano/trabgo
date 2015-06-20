package revista

type Autor struct {
	nome string 
	email string
	senha int
	instituicao string
	endereco string
}

func CriarAutor(nome string, email string, senha int, instituicao string, endereco string) Autor {
	return (Autor{nome, email, senha, instituicao, endereco})
}

func ToString(s Autor) string {
	return s.nome
}

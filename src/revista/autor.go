package revista

type Autor struct {
	nome string 
	email string
	senha int
	instituicao string
	endereco string
}

func (aut *Autor) CriarAutor(nome string, email string, senha int, instituicao string, endereco string) Autor {
	return (Autor{nome, email, senha, instituicao, endereco})
}

func (aut *Autor) ToString() string {
	return aut.nome
}

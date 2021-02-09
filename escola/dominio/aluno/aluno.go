package aluno

type Aluno struct {
	cpf       *CPF
	nome      string
	email     *Email
	telefones []*Telefone
}

func (a *Aluno) AdicionarTelefone(ddd, numero string) error {
	telefone, err := NewTelefone(ddd, numero)
	a.telefones = append(a.telefones, telefone)
	return err
}

func NewAluno(nome string, cpf *CPF, email *Email) *Aluno {
	return &Aluno{
		nome:  nome,
		email: email,
		cpf:   cpf,
	}
}

func (a *Aluno) GetCpf() *CPF {
	return a.cpf
}

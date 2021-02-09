package aluno

type RepositorioDeAlunos interface {
	Matricular(aluno *Aluno) error

	BuscaPorCPF(cpf CPF) (*Aluno, error)

	ListarTodosAlunosMatriculados() ([]*Aluno, error)
}

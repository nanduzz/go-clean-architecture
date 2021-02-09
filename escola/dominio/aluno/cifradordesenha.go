package aluno

type CifradorDeSenha interface {
	CrifrarSenha(senha string) string
	ValidarSenhaCifrada(senhaCifrada, senha string) bool
}

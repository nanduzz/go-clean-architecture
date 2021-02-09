package aluno

import (
	"crypto/md5"
	"encoding/hex"
)

type CifradorDeSenhaMD5 struct{}

func (c *CifradorDeSenhaMD5) CrifrarSenha(senha string) string {

	senhaMD5 := md5.Sum([]byte(senha))
	return hex.EncodeToString(senhaMD5[:])

}
func (c *CifradorDeSenhaMD5) ValidarSenhaCifrada(senhaCifrada, senha string) bool {
	return c.CrifrarSenha(senha) == senhaCifrada
}

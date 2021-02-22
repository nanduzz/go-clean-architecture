package dominio

import "log"

type Ouvinte interface {
	DeveProcessar(evento Evento) bool
	ReageAo(evento Evento)
}

type AbstractOuvinte struct {
	Ouvinte
}

func (o *AbstractOuvinte) Processa(evento Evento) {
	if o.DeveProcessar(evento) {
		log.Println("na struct abstrata")
		o.ReageAo(evento)
	}
}

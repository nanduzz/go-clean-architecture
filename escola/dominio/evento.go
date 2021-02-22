package dominio

import "time"

type Evento interface {
	Momento() time.Time
}

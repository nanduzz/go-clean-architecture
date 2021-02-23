package evento

import "time"

type Evento interface {
	Momento() time.Time
	Tipo() TipoDeEvento
	Informacoes() map[string]interface{}
}

package configuration

import (
	"crypto/tls"

	"github.com/go-pg/pg"
)

// PostgresDatabaseCfg - конфигурация подключений (массив) к разным базам данных Postgres
type PostgresDatabaseCfg []struct {
	Host      string      `json:"host"`
	Port      string      `json:"port"`
	Database  string      `json:"database"`
	Name      string      `json:"name"`
	User      string      `json:"user"`
	Password  string      `json:"password"`
	EnableTLS *tls.Config `json:"enable_tls"`
}

// Connections - подключения к базам данных Postgres
type Connections map[string]*pg.DB

// EstablishPostgresConnections - подключение ко всем указанным в JSON базам данных Postgres
func (cons *Connections) EstablishPostgresConnections(cfg *PostgresDatabaseCfg) {
	(*cons) = make(map[string]*pg.DB)
	for i := range *cfg {
		(*cons)[(*cfg)[i].Name] = pg.Connect(&pg.Options{
			Addr:      (*cfg)[i].Host + ":" + (*cfg)[i].Port,
			User:      (*cfg)[i].User,
			Password:  (*cfg)[i].Password,
			Database:  (*cfg)[i].Database,
			TLSConfig: (*cfg)[i].EnableTLS,
		})
	}
}

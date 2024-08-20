package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	User     string
	Password string
	Port     int
	Dbname   string
}

func (s *Sql) Connect() {
	sqlData := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", s.Host, s.User, s.Password, s.Port, s.Dbname)
	s.Db = sqlx.MustConnect("postgres", sqlData)

	if err := s.Db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect database ok")
	}
}

func (s *Sql) Close() {
	s.Db.Close()
}

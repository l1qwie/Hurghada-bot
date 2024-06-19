package apptype

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Common struct {
	Request     string
	Id          int
	Name        string
	Lastname    string
	Phone       string
	Nickname    string
	Language    string
	ExMessageId int
	Action      string
	DvijId      int
	Level       int
	TitleRu     string
	TitleEn     string
	DiscrpRu    string
	DiscrpEn    string
	DelActivity int
	Changeable  string
	ChangesRu   string
	ChangesEn   string
}

type Answer struct {
	Userid int    `json:"userid"`
	Actid  int    `json:"actid"`
	Answer string `json:"answer"`
	First  bool   `json:"first"`
	Second bool   `json:"second"`
}

func docConnect() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		docHost,
		docPort,
		docUsername,
		docPass,
		docDbname,
		docSslmode)
}

func ConnectToDatabase(doc bool) *sql.DB {
	var (
		db  *sql.DB
		err error
	)
	if doc {
		db, err = sql.Open("postgres", docConnect())
	}
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

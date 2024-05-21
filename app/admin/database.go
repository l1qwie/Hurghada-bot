package admin

import (
	"InfoBot/apptype"

	_ "github.com/lib/pq"
)

func saveInDB(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Buttons (ru, en, status) VALUES ($1, $2, 1)", req.TitleRu, req.TitleEn)
	if err != nil {
		f(err)
	} else {
		_, err := apptype.DB.Exec("INSERT INTO Phrase (ru, en, status) VALUES ($1, $2, 1)", req.DiscrpRu, req.DiscrpEn)
		if err != nil {
			f(err)
		}
	}

}

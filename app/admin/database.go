package admin

import (
	"InfoBot/apptype"

	_ "github.com/lib/pq"
)

func saveInDB(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Phrases (titleru, titleen, discrpru, discrpen, status) VALUES ($1, $2, $3, $4, 1)",
		req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn)
	if err != nil {
		f(err)
	}

}

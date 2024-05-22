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

func findActivity(name int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE name = $1 AND status = 1", name).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func deleteFromDB(name int, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE Phrases SET status = 0 WHERE name = $1", name)
	if err != nil {
		f(err)
	}
}

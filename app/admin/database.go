package admin

import (
	"InfoBot/apptype"
	"fmt"

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

func selectData(which string, nameAct int, f func(error)) (string, string) {
	var ru, en string
	err := apptype.DB.QueryRow(fmt.Sprintf("SELECT %sru, %sen FROM Phrases WHERE name = $1 AND status = 1", which, which), nameAct).Scan(&ru, &en)
	if err != nil {
		f(err)
	}
	return ru, en
}

func selectAllData(nameAct int, f func(error)) (string, string, string, string) {
	var ttru, tten, disru, disen string
	err := apptype.DB.QueryRow("SELECT titleru, titleen, discrpru, discrpen FROM Phrases WHERE name = $1 AND status = 1", nameAct).Scan(&ttru, &tten, &disru, &disen)
	if err != nil {
		f(err)
	}
	return ttru, tten, disru, disen
}

func saveChangesDB(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec(fmt.Sprintf("UPDATE Phrases SET %sru = $1, %sen = $2 WHERE name = $3 AND status = 1", req.Changeable, req.Changeable),
		req.ChangesRu, req.ChangesEn, req.DelActivity)
	if err != nil {
		f(err)
	}
}

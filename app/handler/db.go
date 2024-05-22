package handler

import (
	"InfoBot/apptype"

	_ "github.com/lib/pq"
)

func wichWay(userId int, f func(error)) (bool, error) {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE isadmin = 1 AND userId = $1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0, err
}

func find(userId int, f func(error)) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE userId = $1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func dbRetrieveUser(req *apptype.Common, f func(error)) {
	dbreq := "SELECT action, level, titleru, titleen, discrpru, discrpen FROM Users WHERE userId = $1"
	err := apptype.DB.QueryRow(dbreq, req.Id).Scan(&req.Action, &req.Level, &req.TitleRu, &req.TitleEn, &req.DiscrpRu, &req.DiscrpEn)
	if err != nil {
		f(err)
	}
}

func createUser(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Users (userId, action, isadmin) VALUES ($1, 'new', 0)", req.Id)
	if err != nil {
		f(err)
	}
	req.Action = "new"
	req.Level = 0
}

func retainUser(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE Users SET action = $1, level = $2, titleru = $4, titleen = $5, discrpru = $6, discrpen = $7 WHERE userId = $3",
		req.Action, req.Level, req.Id, req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn)
	if err != nil {
		f(err)
	}
}

func changeStatus(userId int, f func(error)) {
	_, err := apptype.DB.Exec("UPDATE Users SET isadmin = 1 WHERE userId = $1", userId)
	if err != nil {
		f(err)
	}
}

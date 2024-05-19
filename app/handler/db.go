package handler

import "InfoBot/apptype"

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
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Users WHERE (isadmin = 0 OR isadmin = 1) AND userId = $1", userId).Scan(&count)
	if err != nil {
		f(err)
	}
	return count > 0
}

func dbRetrieveUser(req *apptype.Common, f func(error)) {
	dbreq := "SELECT action, level FROM Users WHERE userId = $1"
	err := apptype.DB.QueryRow(dbreq, req.Id).Scan(&req.Action, &req.Level)
	if err != nil {
		f(err)
	}
}

func createUser(req *apptype.Common, f func(error)) {
	_, err := apptype.DB.Exec("INSERT INTO Users (userId, isadmin) VALUES ($1, 0)", req.Id)
	if err != nil {
		f(err)
	}
	req.Action = "new"
	req.Level = 0
}

func dbRetainUser(req *apptype.Common, f func(error)) {
	dbreq := "UPDATE Users SET action = $1, level = $2 WHERE userId = $3"
	_, err := apptype.DB.Exec(dbreq, req.Action, req.Level, req.Id)
	if err != nil {
		f(err)
	}
}

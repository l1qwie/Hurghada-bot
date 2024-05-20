package settest

import "InfoBot/apptype"

func UpdateAction() {
	_, err := apptype.DB.Exec("UPDATE Users SET action = $1 WHERE userId = $2", "new", 999)
	if err != nil {
		panic(err)
	}
}

func UpdateLanguage() {
	_, err := apptype.DB.Exec("UPDATE Users SET language = $1 WHERE userId = $2", "ru", 999)
	if err != nil {
		panic(err)
	}
}

func UpdateLevel(level int) {
	_, err := apptype.DB.Exec("UPDATE Users SET level = $1 WHERE userId = $2", level, 999)
	if err != nil {
		panic(err)
	}
}

func DeleteUser() {
	_, err := apptype.DB.Exec("DELETE FROM Users WHERE userId = $1", 999)
	if err != nil {
		panic(err)
	}
}

func CreateUser() {
	_, err := apptype.DB.Exec("INSERT INTO Users (userId, isadmin) VALUES ($1, $2)", 999, 0)
	if err != nil {
		panic(err)
	}
}

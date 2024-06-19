package notification

import (
	"InfoBot/apptype"
)

func notifone() bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE userid = 999 AND id = 1 AND status = 0 AND notifeone = true").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func notiftwo() bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE userid = 999 AND id = 1 AND status = 0 AND notifeone = true AND notiftwo = true").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func selectDateTime() string {
	var datetime string
	err := apptype.DB.QueryRow("SELECT datetime FROM Dvij WHERE id = 1").Scan(&datetime)
	if err != nil {
		panic(err)
	}
	return datetime
}

func changedstatus1() bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE notifeone = true AND answerone = 1 AND notiftwo = false AND answertwo = 0 AND id = 1 AND userid = 999 AND status != -1").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func changedstatus2() bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE notifeone = true AND answerone = 1 AND notiftwo = true AND answertwo = 1 AND id = 1 AND userid = 999 AND status != -1").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

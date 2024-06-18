package notification

import (
	"InfoBot/apptype"
)

func deleteDvij() {
	_, err := apptype.DB.Exec("DELETE FROM Dvij WHERE id = 1")
	if err != nil {
		panic(err)
	}
}

func deleteClientDvij() {
	_, err := apptype.DB.Exec("DELETE FROM DvijClients WHERE id = 1 AND userid = 999")
	if err != nil {
		panic(err)
	}
}

func createDvij() {
	query := `
        INSERT INTO Dvij (caption, description, datetime, link, status)
        VALUES ('Купаться!', 'Го купаться на пляже!', 
                CURRENT_TIMESTAMP + INTERVAL '12 second', 'https://www.google.com/maps/place/', 0)`
	_, err := apptype.DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func createClientDvij() {
	_, err := apptype.DB.Exec("INSERT INTO DvijClients (id, userid, notiftime, status) VALUES (1, 999, CURRENT_TIMESTAMP, 0)")
	if err != nil {
		panic(err)
	}
}

func reSetSeqDvij() {
	_, err := apptype.DB.Exec("ALTER SEQUENCE dvij_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}

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

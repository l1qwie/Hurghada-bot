package channel

import (
	"InfoBot/apptype"
)

func newActivity() bool {
	var count int
	err := apptype.DB.QueryRow(`SELECT COUNT(*) FROM Dvij WHERE caption = 'Волейбол' AND description = 'Собираемся в зале поиграть в волейбол приглашаем абсолютно всех желающий (ну может быть только без совсем уж маленьких)' 
								AND date = '14.09.2024' AND time = '12:12' AND link = 'https://www.google.com/maps/place/El+Dahar+Bazars/@27.2562144,33.8211347,15.37z/data=!4m6!3m5!1s0x145287e83c67c555:0xa7c4c6807da2a614!8m2!3d27.2518573!4d33.8178499!16s%2Fg%2F11fzb6n0cq?entry=ttu' AND status = 0`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func newRegClient(id int) bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM DvijClients WHERE id = $1 AND userid = 999", id).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func deleteActivity() {
	_, err := apptype.DB.Exec("DELETE FROM Dvij")
	if err != nil {
		panic(err)
	}
	_, err = apptype.DB.Exec("ALTER SEQUENCE dvij_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}

func deleteClientAcivity() {
	_, err := apptype.DB.Exec("DELETE FROM DvijClients")
	if err != nil {
		panic(err)
	}
}

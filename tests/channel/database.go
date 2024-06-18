package channel

import (
	"InfoBot/apptype"
)

func newActivity() bool {
	var count int
	err := apptype.DB.QueryRow(`SELECT COUNT(*) FROM Dvij WHERE caption = 'Купаться!' AND description = 'Го купаться на пляже!' 
								AND datetime = '2024-08-12 12:00:00' AND link = 'https://www.google.com/maps/place/%D0%9F%D0%B0%D1%80%D0%BA+%D0%9C%D0%B5%D0%BD%D0%B0%D1%85%D0%B5%D0%BC+%D0%91%D0%B5%D0%B3%D0%B8%D0%BD/@32.0398261,34.7853328,14.54z/data=!4m6!3m5!1s0x151d4b3c3ec5ba41:0x3ffa8eae4c5afcd1!8m2!3d32.0416723!4d34.8025098!16s%2Fg%2F122j3m_n?entry=ttu' AND status = 0`).Scan(&count)
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

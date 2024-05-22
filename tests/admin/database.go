package admin

import "InfoBot/apptype"

func checkNewActivity() bool {
	var count int
	err := apptype.DB.QueryRow(`SELECT COUNT(*) FROM Phrases WHERE titleru = 'Богослужение' AND titleen = 'Worship Service' AND name = 1 AND status = 1`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

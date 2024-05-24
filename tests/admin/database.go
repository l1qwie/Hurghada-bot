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

func checkDelActivity() bool {
	count := -1
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE name = 1 AND status = 1").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count == 0
}

func checkChangeActivity() bool {
	var count int
	err := apptype.DB.QueryRow("SELECT COUNT(*) FROM Phrases WHERE name = 1 AND status = 1 AND titleru = 'Церковная встреча' AND titleen = 'Church meeting'").Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

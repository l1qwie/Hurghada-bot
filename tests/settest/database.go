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

func CreateActivity() {
	_, err := apptype.DB.Exec(`INSERT INTO Phrases (titleru, titleen, discrpru, discrpen, status) VALUES ('Богослужение', 'Worship Service', 'Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/', 'Every Saturday at 1:00 PM, our worship service starts in the main church building.
	For the first 20-25 minutes, we play and sing songs, then spend some time on prayers, followed by the main sermon. The service usually ends around 3:00 PM, but no one leaves right away, as there is an opportunity to socialize and have some tea. We are looking forward to seeing you!
	
	Start time: Every Saturday at 1:00 PM
	Sermon by: Ahmed Abdulov
	
	We are located here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/', 1)`)
	if err != nil {
		panic(err)
	}
}

func DeleteActivity() {
	_, err := apptype.DB.Exec("DELETE FROM Phrases")
	if err != nil {
		panic(err)
	}
}

func RestartSeq() {
	_, err := apptype.DB.Exec("ALTER SEQUENCE phrases_name_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}

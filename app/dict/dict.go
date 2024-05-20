package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["sayHello"] = "Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!"
	dict["Client"] = "Клиент"
	dict["Admin"] = "Админ"
	dict["lookWhatWeHave"] = "Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас! На этой неделе у нас будут кое-какие активности, и мы будем вас с нетерпением ждать! Ниже нажмите на кнопку, чтобы узнать подробнее"
	dict["worship"] = "Богослужение"
	dict["youthmeeting"] = "Молодежка"
	dict["homegroups"] = "Домашняя Группа"
	dict["prayers"] = "Молитвы"
	dict["men&women"] = "Мужское&Женское"
	dict["film"] = "Фильм"
	dict["worchipDetails"] = `Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`
	dict["MainMenu"] = "Главное Меню"
}

func en(dict map[string]string) {
	dict["sayHello"] = "Hello! This is the chatbot of Hurghada Church! We are happy to see you!"
	dict["Client"] = "Client"
	dict["Admin"] = "Admin"
	dict["lookWhatWeHave"] = "Hello! This is the chatbot of Hurghada Church! We are happy to see you! This week, we will have some activities, and we are eagerly waiting for you! Click the button below to learn more"
	dict["worship"] = "Worship Service"
	dict["youthmeeting"] = "Youth Group"
	dict["homegroups"] = "Home Group"
	dict["prayers"] = "Prayers"
	dict["men&women"] = "Men's & Women's"
	dict["film"] = "Movie"
	dict["worchipDetails"] = `Every Saturday at 1:00 PM, our worship service starts in the main church building.
	For the first 20-25 minutes, we play and sing songs, then spend some time on prayers, followed by the main sermon. The service usually ends around 3:00 PM, but no one leaves right away, as there is an opportunity to socialize and have some tea. We are looking forward to seeing you!
	
	Start time: Every Saturday at 1:00 PM
	Sermon by: Ahmed Abdulov
	
	We are located here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`
	dict["MainMenu"] = "Main Menu"
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
}

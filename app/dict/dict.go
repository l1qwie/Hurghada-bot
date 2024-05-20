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
	dict["youthDetails"] = `Каждое воскресенье у нас проходят молодежные встречи на дому (очень похоже на домашнюю группу, но только для молодых).
	На протяжении около 2 часов мы общаемся, кушаем вкусности, иногда во что-нибудь играем и обсуждаем Библию. Если Вам от 12 до 35 лет, мы очень ждем Вас!
	
	Начало: каждое воскресенье в 16:00
	Ведущий: Ислам Мусульманов
	
	Геолокация иногда меняется, поэтому лучше приходить до 16:00 в наше главное здание и спрашивать у ведущего, где сегодня будет молодежная встреча
	
	Наше главное здание тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`
	dict["homegroupsDetails"] = `Каждый четверг некоторые наши прихожане устраивают домашние группы у себя на дому.
	Если вы хотите обсудить Библию и пообщаться в приятной компании, то вам определенно к нам!
	В среднем такие домашние группы длятся около 2 часов и начинаются ближе к 6-7 часам вечера.
	Расписание может варьироваться в зависимости от того, кто ведущий
	
	Группа 1:
	Ведущий: Антон Ахмедов
	Начало: 18:00
	Адрес: Scandic Resort Hurghada
	Ссылка на адрес: https://www.google.com/maps/place/Scandic+Resort+Hurghada/@27.2429138,33.8417842,18.21z/data=!4m12!1m5!3m4!2zMzbCsDUzJzM2LjQiTiAzMMKwNDInMzQuNSJF!8m2!3d36.893445!4d30.709591!3m5!1s0x14528790e20fc35b:0xfed8fd8e22bb8a2b!8m2!3d27.2431856!4d33.8431078!16s%2Fg%2F11kjzg21ft?entry=ttu
	
	Группа 2:
	Ведущий: Авраам Израилев
	Начало: 19:00
	Адрес: Иерусалим
	Ссылка на адрес: https://www.google.com/maps/place/ZARIFFA+-+%D7%91%D7%99%D7%AA+%D7%A7%D7%A4%D7%94+%D7%99%D7%A8%D7%95%D7%A9%D7%9C%D7%9E%D7%99%E2%80%AD/@31.7556663,35.2016948,18.16z/data=!4m6!3m5!1s0x1502d7e8f28c54b9:0x3cb970274e20cd2d!8m2!3d31.755949!4d35.2027168!16s%2Fg%2F11bxf42czk?entry=ttu`
	dict["prayersDetails"] = `Каждый вторник у нас в главном здании проводится часовая молитва.
	А еще в среду у нас онлайн-молитва в Zoom. Если у вас есть нужда или же вы хотите просто помолиться, то мы ждем вас!
	
	Начало молитвы в здании (Вт): 9:00
	Начало молитвы в Zoom (Ср): 8:00
	
	Мы находимся тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu
	
	Ссылка на Zoom будет отсылаться в наш общий церковный чат. Присоединяйтесь!`
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
	dict["youthDetails"] = `Every Sunday, we have youth meetings at home (similar to home groups but just for young people).
	For about 2 hours, we chat, enjoy some treats, sometimes play games, and discuss the Bible. If you are between 12 and 35 years old, we are eagerly waiting for you!
	
	Start time: Every Sunday at 4:00 PM
	Leader: Islam Muslimov
	
	The location sometimes changes, so it's best to come to our main building before 4:00 PM and ask the leader where the youth meeting will be held today.
	
	Our main building is here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu`
	dict["homegroupsDetails"] = `Every Thursday, some of our congregation members host home groups at their houses.
	If you want to discuss the Bible and socialize in a friendly atmosphere, then you should definitely join us!
	On average, these home groups last about 2 hours and start around 6-7 PM.
	The schedule may vary depending on who the leader is
	
	Group 1:
	Leader: Anton Akhmedov
	Start time: 6:00 PM
	Address: Scandic Resort Hurghada
	Link to address: https://www.google.com/maps/place/Scandic+Resort+Hurghada/@27.2429138,33.8417842,18.21z/data=!4m12!1m5!3m4!2zMzbCsDUzJzM2LjQiTiAzMMKwNDInMzQuNSJF!8m2!3d36.893445!4d30.709591!3m5!1s0x14528790e20fc35b:0xfed8fd8e22bb8a2b!8m2!3d27.2431856!4d33.8431078!16s%2Fg%2F11kjzg21ft?entry=ttu
	
	Group 2:
	Leader: Avraam Izrayilev
	Start time: 7:00 PM
	Address: Jerusalem
	Link to address: https://www.google.com/maps/place/ZARIFFA+-+%D7%91%D7%99%D7%AA+%D7%A7%D7%A4%D7%94+%D7%99%D7%A8%D7%95%D7%A9%D7%9C%D7%9E%D7%99%E2%80%AD/@31.7556663,35.2016948,18.16z/data=!4m6!3m5!1s0x1502d7e8f28c54b9:0x3cb970274e20cd2d!8m2!3d31.755949!4d35.2027168!16s%2Fg%2F11bxf42czk?entry=ttu`
	dict["prayersDetails"] = `Every Tuesday, we have a one-hour prayer session in the main building.
	Additionally, on Wednesday, we have an online prayer meeting on Zoom. If you have a need or just want to pray, we are waiting for you!
	
	Start time for the prayer in the building (Tue): 9:00 AM
	Start time for the prayer on Zoom (Wed): 8:00 AM
	
	We are located here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/data=!4m6!3m5!1s0x145287d03816e835:0xbae794404ccd749!8m2!3d27.2508627!4d33.8319995!16s%2Fg%2F11c2ldpk6q?entry=ttu
	
	The Zoom link will be sent to our church's group chat. Join us!`
	dict["MainMenu"] = "Main Menu"
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
}

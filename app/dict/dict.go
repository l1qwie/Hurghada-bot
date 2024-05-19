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
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
}

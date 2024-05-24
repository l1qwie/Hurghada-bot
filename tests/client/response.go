package client

import (
	"InfoBot/fmtogram/formatter"
	"log"
)

func logs(fm *formatter.Formatter) {
	log.Printf(`fm.Message.Text = "%s"`, fm.Message.Text)
	log.Printf(`fm.Keyboard.Keyboard = "%v"`, fm.Keyboard.Keyboard)
	log.Printf(`fm.Message.ChatID = %d`, fm.Message.ChatID)
}

func greeteings(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Клиент", "Админ"}, []string{"client", "admin"}, []string{"cmd", "cmd"}, true)
}

func showSchedule(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас! На этой неделе у нас будут кое-какие активности, и мы будем вас с нетерпением ждать! Ниже нажмите на кнопку, чтобы узнать подробнее", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1, 1, 1, 1}, []string{"Богослужение", "Молодежка", "Домашняя Группа", "Молитвы", "Мужское&Женское", "Фильм"},
		[]string{"1", "2", "3", "4", "5", "6"}, []string{"cmd", "cmd", "cmd", "cmd", "cmd", "cmd"}, true)
}

func showWorship(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showYouth(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждое воскресенье у нас проходят молодежные встречи на дому (очень похоже на домашнюю группу, но только для молодых).
	На протяжении около 2 часов мы общаемся, кушаем вкусности, иногда во что-нибудь играем и обсуждаем Библию. Если Вам от 12 до 35 лет, мы очень ждем Вас!
	
	Начало: каждое воскресенье в 16:00
	Ведущий: Ислам Мусульманов
	
	Геолокация иногда меняется, поэтому лучше приходить до 16:00 в наше главное здание и спрашивать у ведущего, где сегодня будет молодежная встреча
	
	Наше главное здание тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showHomeGroups(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждый четверг некоторые наши прихожане устраивают домашние группы у себя на дому.
	Если вы хотите обсудить Библию и пообщаться в приятной компании, то вам определенно к нам!
	В среднем такие домашние группы длятся около 2 часов и начинаются ближе к 6-7 часам вечера.
	Расписание может варьироваться в зависимости от того, кто ведущий
	
	Группа 1:
	Ведущий: Антон Ахмедов
	Начало: 18:00
	Адрес: Scandic Resort Hurghada
	Ссылка на адрес: https://www.google.com/maps/place/Scandic+Resort+Hurghada/@27.2429138,33.8417842,18.21z/
	
	Группа 2:
	Ведущий: Авраам Израилев
	Начало: 19:00
	Адрес: Иерусалим
	Ссылка на адрес: https://www.google.com/maps/place/`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showPrayers(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждый вторник у нас в главном здании проводится часовая молитва.
	А еще в среду у нас онлайн-молитва в Zoom. Если у вас есть нужда или же вы хотите просто помолиться, то мы ждем вас!
	
	Начало молитвы в здании (Вт): 9:00
	Начало молитвы в Zoom (Ср): 8:00
	
	Мы находимся тут:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/
	
	Ссылка на Zoom будет отсылаться в наш общий церковный чат. Присоединяйтесь!`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)

}

func showMandW(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждый понедельник и четверг у нас проводятся гендерные группы.
	Для мужчин и женщин соответственно. Мы там также молимся, обсуждаем Библию,
	но на этот раз только в кругу своего же пола. Заодно мы можем обсудить некоторые свои проблемы,
	которые мы не рискнули бы обсудить в смешанной группе. Если у вас есть запрос на такую встречу, то мы ждем вас!
	
	Начало мужской группы (Пн): 9:00
	Адрес: Дом напротив Tiba Perfumes
	Ссылка на адрес: https://www.google.com/maps/place/Tiba+Perfumes/@27.2233417,33.8372355,19.79z/
	
	Начало женской группы (Чт): 10:00
	Адрес: Arabia Azur Resort
	Ссылка на адрес: https://www.google.com/maps/place/Arabia+Azur+Resort/@27.2404984,33.8424506,18.13z/`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func showFilm(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Каждую пятницу наш дорогой друг проводит показ кино на своей вилле!
	Хотите хорошо провести время, расслабиться, пообщаться и посмотреть прекрасный фильм?
	Тогда вам определенно стоит посетить данное мероприятие
	
	Хозяин виллы: Мойша Исраель
	Начало: 20:00
	Фильм: Гладиатор (2000)
	Ссылка на адрес: https://www.google.com/maps/place/
	
	Совет: лучше всего добираться на такси, так как ехать далеко и автобусы почти не ходят`, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

package admin

import (
	"InfoBot/fmtogram/formatter"
	"log"
)

func logs(fm *formatter.Formatter) {
	log.Printf(`fm.Message.Text = "%s"`, fm.Message.Text)
	log.Printf(`fm.Keyboard.Keyboard = "%v"`, fm.Keyboard.Keyboard)
	log.Printf(`fm.Message.ChatID = %d`, fm.Message.ChatID)
}

func usual(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
}

func greeteings(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Клиент", "Админ"}, []string{"client", "admin"}, []string{"cmd", "cmd"}, true)
}

func showOptions(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Выберите опцию", true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Создать", "Изменить", "Удалить", "Главное Меню"}, []string{"create", "change", "delete", "MainMenu"}, []string{"cmd", "cmd", "cmd", "cmd"}, true)
}

func chCreateOpt(fm *formatter.Formatter) {
	usual(fm)
	fm.AssertString("Введите название активности (на русском языке), которую вы добавляете. Это название будет отображаться на кнопках", true)
}

func sentNameRu(fm *formatter.Formatter) {
	usual(fm)
	fm.AssertString("Введите название активности (на английском языке), которую вы добавляете. Это название будет отображаться на кнопках", true)
}

func sentNameEn(fm *formatter.Formatter) {
	usual(fm)
	fm.AssertString("Введите описание активности (на русском языке), которую добавляете. Это описание будет отображаться после нажатия на кнопку, которую вы только что создали", true)
}

func sentTextRu(fm *formatter.Formatter) {
	usual(fm)
	fm.AssertString("Введите описание активности (на английском языке), которую добавляете. Это описание будет отображаться после нажатия на кнопку, которую вы только что создали", true)
}

func sentTextEn(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString(`Вы практически создали новую активность. Осталось ее подтвердить

	Название (ru): Богослужение
	
	Название (en): Worship Service
	
	Описание (ru): Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/
	
	Описание (en): Every Saturday at 1:00 PM, our worship service starts in the main church building.
	For the first 20-25 minutes, we play and sing songs, then spend some time on prayers, followed by the main sermon. The service usually ends around 3:00 PM, but no one leaves right away, as there is an opportunity to socialize and have some tea. We are looking forward to seeing you!
	
	Start time: Every Saturday at 1:00 PM
	Sermon by: Ahmed Abdulov
	
	We are located here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/
	
	Если все верно, то нажмите сохранить, если нет - измените, что пожелаете`, true)
	fm.AssertInlineKeyboard([]int{1, 1, 1}, []string{"Сохранить", "Изменить", "Главное Меню"}, []string{"save", "change", "MainMenu"}, []string{"cmd", "cmd", "cmd"}, true)
}

func sentSave(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Данные успешно изменены!", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Главное Меню"}, []string{"MainMenu"}, []string{"cmd"}, true)
	if !checkNewActivity() {
		panic("The new activity wasn't added in the database")
	}
}

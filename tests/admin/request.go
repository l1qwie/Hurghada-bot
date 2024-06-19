package admin

import "InfoBot/fmtogram/types"

func chList() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "list"
	return tr
}

func chNameRu() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "Богослужение"
	return tr
}

func chNameEn() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "Worship Service"
	return tr
}

func chTextRu() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = `Каждую субботу в 13:00 у нас в главном здании церкви начинается богослужение.
	На протяжении примерно 20-25 минут в начале мы играем и поем песни, дальше некоторое количество
	времени уделяется молитвам, а потом главная проповедь. Обычно богослужение заканчивается ближе к 15:00,
	но никто никуда не расходится, так как можно пообщаться и выпить чаю. Очень ждем именно Вас!
	
	Начало: каждая суббота в 13:00
	Проповедь будет вести: Ахмед Абдулов
	
	Мы находимся тут: 
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/`
	return tr
}

func chTextEn() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = `Every Saturday at 1:00 PM, our worship service starts in the main church building.
	For the first 20-25 minutes, we play and sing songs, then spend some time on prayers, followed by the main sermon. The service usually ends around 3:00 PM, but no one leaves right away, as there is an opportunity to socialize and have some tea. We are looking forward to seeing you!
	
	Start time: Every Saturday at 1:00 PM
	Sermon by: Ahmed Abdulov
	
	We are located here:
	
	https://www.google.com/maps/place/Good+Shepered+Association/@27.2508176,33.8318688,20.78z/`
	return tr
}

func chSave() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "save"
	return tr
}

func chActivity() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "1"
	return tr
}

func chPerson() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "999"
	return tr
}

func chChangeable() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "title"
	return tr
}

func chNewTitleRu() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "Церковная встреча"
	return tr
}

func chNewTitleEn() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "Church meeting"
	return tr
}

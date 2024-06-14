package channel

import (
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
)

func checkProgress(fm *formatter.Formatter) {
	if !newActivity() {
		panic("THE NEW ACIVITY WASN'T ADDED IN THE DATABASE")
	}
	fm.AssertChatName("testdvijhurghada", true)
	fm.AssertString("Название: <b>Волейбол</b>\n\n\nОписание <b>Собираемся в зале поиграть в волейбол приглашаем абсолютно всех желающий (ну может быть только без совсем уж маленьких)</b>\n\nДата и Время: <b>14.09.2024 12:12</b>\n\nСсылка на место проведения: https://www.google.com/maps/place/El+Dahar+Bazars/@27.2562144,33.8211347,15.37z/data=!4m6!3m5!1s0x145287e83c67c555:0xa7c4c6807da2a614!8m2!3d27.2518573!4d33.8178499!16s%2Fg%2F11fzb6n0cq?entry=ttu", true)
	fm.AssertParseMode("HTML", true)
	fm.AssertInlineKeyboard([]int{1}, []string{"Я приду!"}, []string{"1"}, []string{"cmd"}, true)
}

func prepareReq() *types.TelegramResponse {
	return &types.TelegramResponse{
		Result: []types.TelegramUpdate{{
			Query: types.Callback{
				TypeFrom: types.User{
					UserID:   999,
					IsBot:    false,
					Language: "ru",
				},
				Message: types.Message{
					MessageId: 1,
					Chat: types.Chat{
						Id:   111111,
						Type: "channel",
					},
				},
				Data: "1",
			}},
		},
	}
}

func checkResponse() {
	if !newRegClient(1) {
		panic("THE CLIENT WASN'T REGISTRED TO THE ACTIVITY")
	}
}

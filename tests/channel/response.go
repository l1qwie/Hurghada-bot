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
	fm.AssertString("Название: <b>Купаться!</b>\n\n\nОписание <b>Го купаться на пляже!</b>\n\nДата и Время: <b>12.08.2024 12:00:00</b>\n\nСсылка на место проведения: https://www.google.com/maps/place/%D0%9F%D0%B0%D1%80%D0%BA+%D0%9C%D0%B5%D0%BD%D0%B0%D1%85%D0%B5%D0%BC+%D0%91%D0%B5%D0%B3%D0%B8%D0%BD/@32.0398261,34.7853328,14.54z/data=!4m6!3m5!1s0x151d4b3c3ec5ba41:0x3ffa8eae4c5afcd1!8m2!3d32.0416723!4d34.8025098!16s%2Fg%2F122j3m_n?entry=ttu", true)
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

func checkResponse(fm *formatter.Formatter) {
	if !newRegClient(1) {
		panic("THE CLIENT WASN'T REGISTRED TO THE ACTIVITY")
	}
	fm.AssertChatId(999, true)
	fm.AssertString("Вы записались на событие из телеграм канала testDvijHurghada. Мы уведомим вас за сутки и за несколько часов до начала. Вам нужно будет подтвердить, что вы точно придете", true)
}

package client

import (
	"InfoBot/fmtogram/formatter"
	"fmt"
	"log"
)

func logs(fm *formatter.Formatter) {
	log.Print("fm.Message.Text =", fm.Message.Text)
	log.Print("fm.Message.ReplyMarkup =", fm.Message.ReplyMarkup)
	log.Print("fm.Message.ChatID =", fm.Message.ChatID)
}

func greeteings(fm *formatter.Formatter) {
	fmt.Println("SHALOM")
	logs(fm)
	fmt.Println("SHALOM2")
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас!", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Клиент", "Админ"}, []string{"client", "admin"}, []string{"cmd", "cmd"}, true)
}

func showSchedule(fm *formatter.Formatter) {
	logs(fm)
	fm.AssertChatId(999, true)
	fm.AssertString("Привет! Это чат-бот церкви Хургады! Мы рады видеть Вас! На этой неделе у нас будут кое-какие активности, и мы будем вас с нетерпением ждать! Ниже нажмите на кнопку, чтобы узнать подробнее", true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Богослужение", "Молодежка", "Домашняя Группа", "Молитвы", "Мужское&Женское", "Фильм"},
		[]string{"worship", "youthmeeting", "homegroups", "prayers", "men&women", "film"}, []string{"cmd", "cmd", "cmd", "cmd", "cmd", "cmd"}, true)
}

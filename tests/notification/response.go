package notification

import (
	"InfoBot/fmtogram/formatter"
	"fmt"
	"log"
	"time"
)

func timeManager() string {
	datetime := selectDateTime()
	parsedTime, err := time.Parse(time.RFC3339Nano, datetime)
	if err != nil {
		panic(err)
	}
	date := parsedTime.Format("2006-01-02")
	time := parsedTime.Format("15:04:05")
	return fmt.Sprintf("%s %s", date, time)
}

func checkProgress1(fm *formatter.Formatter) {
	if !notifone() {
		panic("The user wasn't notified once")
	}
	datetime := timeManager()
	fm.AssertChatId(999, true)
	link := "\n\nСсылка на место проведения: https://www.google.com/maps/place/"
	fm.AssertString(fmt.Sprintf("До начала активности на которую вы подписались отслось меньше суток! Вы все еще планируете поучавствовать?\n\n\nНазвание: <b>Купаться!</b>\n\n\nОписание <b>Го купаться на пляже!</b>\n\nДата и Время: <b>%s</b>%s", datetime, link), true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Да, я буду!", "Нет, я не смогу!"}, []string{`{"userid":999,"actid":1,"answer":"yes","first":true,"second":false}`, `{"userid":999,"actid":1,"answer":"no","first":true,"second":false}`}, []string{"cmd", "cmd"}, true)
	fm.AssertParseMode("HTML", true)
	log.Print("Test Notificationtest1() completed")
}

func checkProgress2(fm *formatter.Formatter) {
	if !notiftwo() {
		panic("The user wasn't notified twice")
	}
	datetime := timeManager()
	fm.AssertChatId(999, true)
	link := "\n\nСсылка на место проведения: https://www.google.com/maps/place/"
	fm.AssertString(fmt.Sprintf("До начала активности на которую вы подписались отслось меньше двух часов! Вы все еще планируете поучавствовать?\n\n\nНазвание: <b>Купаться!</b>\n\n\nОписание <b>Го купаться на пляже!</b>\n\nДата и Время: <b>%s</b>%s", datetime, link), true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Да, я буду!", "Нет, я не смогу!"}, []string{`{"userid":999,"actid":1,"answer":"yes","first":true,"second":true}`, `{"userid":999,"actid":1,"answer":"no","first":true,"second":true}`}, []string{"cmd", "cmd"}, true)
	fm.AssertParseMode("HTML", true)
	log.Print("Test Notificationtest2() completed")
}

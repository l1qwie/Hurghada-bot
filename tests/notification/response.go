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
	fm.AssertString(fmt.Sprintf("До начала активности на которую вы подписались осталось меньше суток! Вы все еще планируете поучавствовать?\n\n\nНазвание: <b>Купаться!</b>\n\n\nОписание <b>Го купаться на пляже!</b>\n\nДата и Время: <b>%s</b>%s", datetime, link), true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Да, я буду!", "Нет, я не смогу!"}, []string{"999.1.yes.true.false", "999.1.no.true.false"}, []string{"cmd", "cmd"}, true)
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
	fm.AssertString(fmt.Sprintf("До начала активности на которую вы подписались осталось меньше двух часов! Вы все еще планируете поучавствовать?\n\n\nНазвание: <b>Купаться!</b>\n\n\nОписание <b>Го купаться на пляже!</b>\n\nДата и Время: <b>%s</b>%s", datetime, link), true)
	fm.AssertInlineKeyboard([]int{1, 1}, []string{"Да, я буду!", "Нет, я не смогу!"}, []string{"999.1.yes.true.true", "999.1.no.true.true"}, []string{"cmd", "cmd"}, true)
	fm.AssertParseMode("HTML", true)
	log.Print("Test Notificationtest2() completed")
}

func checkAnswer1(fm *formatter.Formatter) {
	if !changedstatus1() {
		panic("The agree status wasn't changed in the first time")
	}
	fm.AssertChatId(999, true)
	fm.AssertString("Мы очень рады, что вы все еще хотите поучаствовать! Будем вас ждать! Мы пришлем вам еще одно уведомление за пару часов до начала", true)
	log.Print("Test checkAnswer1() completed")
}

func checkAnswer2(fm *formatter.Formatter) {
	if !changedstatus2() {
		panic("The agree status wasn't changed in the second time")
	}
	fm.AssertChatId(999, true)
	fm.AssertString("Мы очень рады, что вы все еще хотите поучаствовать! Будем вас ждать!", true)
	log.Print("Test checkAnswer2() completed")
}

package client

import (
	"InfoBot/app/dict"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"strconv"
)

const (
	START  int = 0
	LEVEL1 int = 1
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

// Checkes is there an int number or not
// If yes - returns true and the number
// Else - returns only false (and 0 in var num)
func IntCheck(phrase string) (bool, int) {
	num, err := strconv.Atoi(phrase)
	return err == nil, num
}

func mainMenu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Action = "divarication"
	crd := LengthOfNames(fm.Error)
	names := SelectNames(req.Language, len(crd), fm.Error)
	fm.WriteString(dict["lookWhatWeHave"])
	setKb(fm, crd, names, SelectValues(names, req.Language, len(crd), fm.Error))
}

func showDetails(fm *formatter.Formatter, dict map[string]string, req, lang string) {
	details := selectDetails(req, lang, fm.Error)
	fm.WriteString(details)
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func sayHello(fm *formatter.Formatter, dict map[string]string) {
	fm.WriteString(dict["sayHello"])
	setKb(fm, []int{1, 1}, []string{dict["Client"], dict["Admin"]}, []string{"client", "admin"})
}

func lookWhatWeHave(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "client" {
		mainMenu(req, fm, dict)
	} else {
		sayHello(fm, dict)
	}
}

func greeteings(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		sayHello(fm, dict)
	} else if req.Level == LEVEL1 {
		lookWhatWeHave(req, fm, dict)
	}
}

func divarication(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	det, _ := IntCheck(req.Request)
	if det {
		if findName(req.Request, fm.Error) {
			req.Action = req.Request
			showDetails(fm, dict, req.Request, req.Language)
		} else {
			mainMenu(req, fm, dict)
		}
	} else {
		mainMenu(req, fm, dict)
	}
}

func checkLanguage(req *apptype.Common) {
	if req.Language != "ru" && req.Language != "en" {
		req.Language = "en"
	}
}

func Dispatcher(req *apptype.Common, fm *formatter.Formatter) {
	checkLanguage(req)
	if req.Request == "MainMenu" {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "divarication" {
		divarication(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "new" {
		greeteings(req, fm, dict.Dictionary[req.Language])
	} else {
		showDetails(fm, dict.Dictionary[req.Language], req.Action, req.Language)
	}
	fm.WriteChatId(req.Id)
}

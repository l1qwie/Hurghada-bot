package client

import (
	"InfoBot/app/dict"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
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

func worship(fm *formatter.Formatter, dict map[string]string) {
	fm.WriteString(dict["worchipDetails"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func youthmeeting() {

}

func homegroups() {

}

func prayers() {

}

func mewAndWomen() {

}

func film() {

}

func mainMenu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Action = "divarication"
	fm.WriteString(dict["lookWhatWeHave"])
	setKb(fm, []int{1, 1, 1, 1, 1, 1}, []string{dict["worship"], dict["youthmeeting"], dict["homegroups"], dict["prayers"], dict["men&women"], dict["film"]},
		[]string{"worship", "youthmeeting", "homegroups", "prayers", "men&women", "film"})
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
	if req.Request == "worship" {
		req.Action = "worship"
		worship(fm, dict)
	} else {
		mainMenu(req, fm, dict)
	}

}

func Dispatcher(req *apptype.Common, fm *formatter.Formatter) {
	if req.Request == "MainMenu" {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "worship" {
		worship(fm, dict.Dictionary[req.Language])
	} else if req.Action == "youthmeeting" {
		youthmeeting()
	} else if req.Action == "homegroups" {
		homegroups()
	} else if req.Action == "prayers" {
		prayers()
	} else if req.Action == "men&women" {
		mewAndWomen()
	} else if req.Action == "film" {
		film()
	} else if req.Action == "divarication" {
		divarication(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "new" {
		greeteings(req, fm, dict.Dictionary[req.Language])
	}
	fm.WriteChatId(req.Id)
}

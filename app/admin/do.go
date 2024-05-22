package admin

import (
	"InfoBot/app/client"
	"InfoBot/app/dict"
	"InfoBot/apptype"
	"InfoBot/fmtogram/formatter"
	"fmt"
)

const (
	START  int = 0
	LEVEL1 int = 1
	LEVEL2 int = 2
	LEVEL3 int = 3
	LEVEL4 int = 4
	LEVEL5 int = 5
	OPT    int = 3
)

// Sets any keyboard
func setKb(fm *formatter.Formatter, crd []int, names, data []string) {
	fm.SetIkbdDim(crd)
	for i := 0; i < len(crd) && i < len(names); i++ {
		fm.WriteInlineButtonCmd(names[i], data[i])
	}
}

func mainMenu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Action = "divarication"
	req.Level = OPT
	fm.WriteString(dict["chooseOpt"])
	setKb(fm, []int{1, 1, 1}, []string{dict["create"], dict["change"], dict["delete"]}, []string{"create", "change", "delete"})
}

func checkLanguage(req *apptype.Common) {
	if req.Language != "ru" && req.Language != "en" {
		req.Language = "en"
	}
}

func titleRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 1
	fm.WriteString(dict["nameRu"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func titleEn(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 2
	req.TitleRu = req.Request
	fm.WriteString(dict["nameEn"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func discrpRu(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 3
	req.TitleEn = req.Request
	fm.WriteString(dict["textRu"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func discrpEn(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 4
	req.DiscrpRu = req.Request
	fm.WriteString(dict["textEn"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func sendAlmoseDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 5
	fm.WriteString(fmt.Sprintf(dict["almostDone"], req.TitleRu, req.TitleEn, req.DiscrpRu, req.DiscrpEn))
	setKb(fm, []int{1, 1, 1}, []string{dict["save"], dict["change"], dict["MainMenu"]}, []string{"save", "change", "MainMenu"})
}

func almostDone(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.DiscrpEn = req.Request
	sendAlmoseDone(req, fm, dict)
}

func done(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	saveInDB(req, fm.Error)
	fm.WriteString(dict["done"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func doneOrChange(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "save" {
		done(req, fm, dict)
	} else if req.Request == "change" {
		req.Action = req.Request
		//change()
	} else {
		sendAlmoseDone(req, fm, dict)
	}
}

func create(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		titleRu(req, fm, dict)
	} else if req.Level == LEVEL1 {
		titleEn(req, fm, dict)
	} else if req.Level == LEVEL2 {
		discrpRu(req, fm, dict)
	} else if req.Level == LEVEL3 {
		discrpEn(req, fm, dict)
	} else if req.Level == LEVEL4 {
		almostDone(req, fm, dict)
	} else if req.Level == LEVEL5 {
		doneOrChange(req, fm, dict)
	}
}

func chooseActivity(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	req.Level = 1
	crd := client.LengthOfNames(fm.Error)
	names := client.SelectNames(req.Language, len(crd), fm.Error)
	data := client.SelectValues(names, req.Language, len(crd), fm.Error)
	crd = append(crd, 1)
	names = append(names, dict["MainMenu"])
	data = append(data, "MainMenu")
	setKb(fm, crd, names, data)
	fm.WriteString(dict["chooseAnActiv"])
}

func whichToDel(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Request == "delete" {
		chooseActivity(req, fm, dict)
	} else {
		mainMenu(req, fm, dict)
	}
}

func goToDelete(req *apptype.Common, fm *formatter.Formatter, num int, dict map[string]string) {
	req.DelActivity = num
	deleteFromDB(req.DelActivity, fm.Error)
	fm.WriteString(dict["done"])
	setKb(fm, []int{1}, []string{dict["MainMenu"]}, []string{"MainMenu"})
}

func deleted(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	det, num := client.IntCheck(req.Request)
	if det {
		if findActivity(num, fm.Error) {
			goToDelete(req, fm, num, dict)
		} else {
			chooseActivity(req, fm, dict)
		}
	} else {
		chooseActivity(req, fm, dict)
	}
}

func delete(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	if req.Level == START {
		whichToDel(req, fm, dict)
	} else if req.Level == LEVEL1 {
		deleted(req, fm, dict)
	}
}

func changeLvlAndAct(req *apptype.Common) {
	req.Level = START
	req.Action = req.Request
}

func divarication(req *apptype.Common, fm *formatter.Formatter, dict map[string]string) {
	changeLvlAndAct(req)
	if req.Request == "create" {
		create(req, fm, dict)
	} else if req.Request == "change" {
		//change()
	} else if req.Request == "delete" {
		delete(req, fm, dict)
	} else {
		mainMenu(req, fm, dict)
	}
}

func Dispatcher(req *apptype.Common, fm *formatter.Formatter) {
	checkLanguage(req)
	if req.Request == "MainMenu" {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "divarication" {
		divarication(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "create" {
		create(req, fm, dict.Dictionary[req.Language])
	} else if req.Action == "change" {

	} else if req.Action == "delete" {
		delete(req, fm, dict.Dictionary[req.Language])
	} else {
		mainMenu(req, fm, dict.Dictionary[req.Language])
	}
	fm.WriteChatId(req.Id)
}

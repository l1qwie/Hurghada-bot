package settest

import (
	"InfoBot/app"
	"InfoBot/fmtogram/errors"
	"InfoBot/fmtogram/formatter"
	"InfoBot/fmtogram/types"
	"log"
)

const wrongAnswers int = 2

// This stuct only for Update functions to update database
type Update struct {
	Act, Lang     string
	Level, UserId int
}

// This stuct is for all tests. The main thought is you can use it anywhere
type TestStuct struct {
	TRcount, Trshcount, Round, Wcounter, inficount int
	FuncReq                                        []func() *types.TelegramResponse
	FuncRes                                        []func(*formatter.Formatter)
	FuncTrsh                                       []func() *types.TelegramResponse
	request                                        *types.TelegramResponse
	response                                       *formatter.Formatter
	UpdtLevel                                      []int
	Name                                           string
	Logf                                           string
}

// Check trash-query, call a func and return true
// else - return false
func (t *TestStuct) checkTheTrash() bool {
	if t.Trshcount < 2 {
		t.request = t.FuncTrsh[t.inficount]()
		t.inficount++
	}
	return t.Trshcount < 2
}

// The fucntion has to check what answer is and if it's worng answer - call a func and return true
// else - return false
// only if the main counter != 0
func (t *TestStuct) checkTheWorng() bool {
	if t.Wcounter < wrongAnswers {
		if t.TRcount != 0 {
			t.FuncRes[t.TRcount-1](t.response)
		}
	}
	return (t.Wcounter < wrongAnswers) && (t.TRcount != 0)
}

func (t *TestStuct) theHeadTheAdminCh() {
	if t.TRcount > 4 && t.TRcount < 7 {
		t.Trshcount = 3
		t.request = t.FuncReq[t.TRcount]()
	} else if !t.checkTheTrash() {
		t.request = t.FuncReq[t.TRcount]()
	}
}

func (t *TestStuct) theHeadTheAdmin() {
	if t.TRcount > 2 && t.TRcount < 7 {
		t.Trshcount = 3
		t.request = t.FuncReq[t.TRcount]()
	} else if !t.checkTheTrash() {
		t.request = t.FuncReq[t.TRcount]()
	}
}

func (t *TestStuct) theHeadTheClient() {
	if !t.checkTheTrash() {
		t.request = t.FuncReq[t.TRcount]()
	}
}

// Preparation the datbase. Just the routine
func (t *TestStuct) prepDatabase() {
	UpdateAction()
	UpdateLanguage()
	UpdateLevel(t.UpdtLevel[t.TRcount])
}

// Accept bot answers
// Call a func if it's not a wrong answer
func (t *TestStuct) acceptAnswers() {
	if !t.checkTheWorng() || t.TRcount > 2 && t.TRcount < 6 {
		t.FuncRes[t.TRcount](t.response)
	}
}

func (t *TestStuct) acceptAnswersTheAdminCh() {
	if t.TRcount > 4 && t.TRcount < 7 {
		t.FuncRes[t.TRcount](t.response)
	} else if !t.checkTheWorng() {
		t.FuncRes[t.TRcount](t.response)
	}
}

func (t *TestStuct) acceptAnswersTheAdmin() {
	if t.TRcount > 2 && t.TRcount < 7 {
		t.FuncRes[t.TRcount](t.response)
	} else if !t.checkTheWorng() {
		t.FuncRes[t.TRcount](t.response)
	}
}

func (t *TestStuct) acceptAnswersTheClient() {
	if !t.checkTheWorng() {
		t.FuncRes[t.TRcount](t.response)
	}
}

// The body for all tests
// Only this function could be imported
func (t *TestStuct) DoTest(ad, change bool) {
	for t.TRcount < t.Round {
		t.Trshcount = 0
		t.Wcounter = 0
		for t.Trshcount < 3 {
			if ad {
				if change {
					t.theHeadTheAdminCh()
				} else {
					t.theHeadTheAdmin()
				}
			} else {
				t.theHeadTheClient()
			}
			//t.prepDatabase()
			UpdateLevel(t.UpdtLevel[t.TRcount])
			t.response = app.Receiving(t.request)
			if t.response.Err != nil {
				errors.MadeMisstake(t.response.Err)
				panic(t.response.Err)
			}
			if ad {
				if change {
					t.acceptAnswersTheAdminCh()
				} else {
					t.acceptAnswersTheAdmin()
				}
			} else {
				t.acceptAnswersTheClient()
			}
			t.Trshcount++
			t.Wcounter++
		}
		log.Printf("%s %d has been complete\n", t.Name, t.TRcount)
		t.TRcount++
	}
}

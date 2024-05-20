package client

import "InfoBot/fmtogram/types"

func mainReq() *types.TelegramResponse {
	return &types.TelegramResponse{
		Result: []types.TelegramUpdate{{
			Query: types.Callback{
				TypeFrom: types.User{
					UserID:   999,
					IsBot:    false,
					Language: "ru",
				},
			}},
		},
	}
}

func trfunc() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "ALSDLK:AKL:DAL:KSDASD"
	return tr
}

func trfunc1() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "ALSD!!!!!!!!!!!!!:KSDASD"
	return tr
}

func trfunc2() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "ALSDLIIIIIIIIIIIIIIiDASD"
	return tr
}

func trfunc3() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "WHAT THE GELL!?"
	return tr
}

func trfunc4() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "OHJ!?"
	return tr
}

func trfunc5() *types.TelegramResponse {
	tr := mainReq()
	tr.Result[0].Query.Data = "SO, WHAT WE GONNA DO?"
	return tr
}

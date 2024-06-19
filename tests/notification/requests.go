package notification

import "InfoBot/fmtogram/types"

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
						Type: "private",
					},
				},
			}},
		},
	}
}

func createReq1() *types.TelegramResponse {
	tr := prepareReq()
	tr.Result[0].Query.Data = "999.1.yes.true.false"
	return tr
}

func createReq2() *types.TelegramResponse {
	tr := prepareReq()
	tr.Result[0].Query.Data = "999.1.yes.true.true"
	return tr
}

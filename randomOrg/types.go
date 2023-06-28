package randomOrg

type DiceResponse struct {
	ID     int        `json:"id"`
	Result DiceResult `json:"result"`
}

type DiceResult struct {
	Random         RandomResult `json:"random"`
	CompletionTime string       `json:"completionTime"`
	RequestsLeft   int          `json:"requestsLeft"`
}

type RandomResult struct {
	Data []uint8 `json:"data"`
}

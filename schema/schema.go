package schema

type RaffleEntry struct {
	Email  string `json:"email" bson:"email"`
	Raffle int64  `json:"raffle" bson:"raffle"`
	Date   string `json:"date" bson:"date"`
	Winner bool   `json:"winner" bson:"winner"`
}

type Response struct {
	Msg     string         `json:"msg"`
	Err     string         `json:"err"`
	Winners []*RaffleEntry `json:"winners"`
}

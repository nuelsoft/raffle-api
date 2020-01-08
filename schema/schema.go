package schema

type RaffleEntry struct {
	Email      string `json:"email" bson:"email"`
	Name       string `json:"name" bson:"name"`
	Phone      string `json:"phone" bson:"phone"`
	Raffle     int64  `json:"raffle" bson:"raffle"`
	Date       string `json:"date" bson:"date"`
	Winner     bool   `json:"winner" bson:"winner"`
	PaymentRef string `json:"payment_ref" bson:"payment_ref"`
}

type Response struct {
	Msg     string         `json:"msg"`
	Err     string         `json:"err"`
	Winners []*RaffleEntry `json:"winners"`
}

package entities

// RandomNumber defines the structure of the response gotten from the
// external random number service
type RandomNumber struct {
	RandomNumber int8 `json:"random_number"`
}

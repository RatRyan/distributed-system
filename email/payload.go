package main

type userPayload struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type offerPaylood struct {
	Offeror string `json:"Offeror"`
	Offeree string `json:"Offeree"`
}

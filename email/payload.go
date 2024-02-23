package main

type userPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type offerPaylood struct {
	Offeror string `json:"offeror"`
	Offeree string `json:"offeree"`
}

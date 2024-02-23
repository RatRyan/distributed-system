package main

import "fmt"

func handleOffer(key string, offer offerPaylood) {
	switch key {
	case "offer_created":
		fmt.Printf("offer created between: %v and %v", offer.Offeror, offer.Offeree)
	case "offer_accepted":
		fmt.Printf("offer accepted between: %v and %v", offer.Offeror, offer.Offeree)
	case "offer_rejected":
		fmt.Printf("offer rejected between: %v and %v", offer.Offeror, offer.Offeree)
	}
}

package main

import "fmt"

func handleUser(key string, user userPayload) {
	switch key {
	case "password_change":
		fmt.Println("password changed")
	}
}

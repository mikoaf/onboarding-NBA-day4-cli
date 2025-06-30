package server

import "fmt"

func HandleApiCallback(message string){
	fmt.Println("Callback from API:", message)
}

func HandleHardwareCallback(message string){
	fmt.Println("Callback from Hardware:", message)
}
package main

func sendDataToRoutine(ch chan string) {
	ch <- "Hello from the channel!"
}

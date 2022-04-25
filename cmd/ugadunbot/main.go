package main

import (
	"flag"
	"log"
	tgClient "ugadunbot/internal/clients/telegram"
	event_consumer "ugadunbot/internal/consumer/event-consumer"
	"ugadunbot/internal/events/telegram"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()))

	log.Printf("Cервис запущен")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Cервис остановился", err)
	}

}

func mustToken() string {

	token := flag.String("tg-bot-token", "", "тут должен быть токен")

	flag.Parse()

	if *token == "" {
		log.Fatal("токена нет ")

	}
	return *token

}

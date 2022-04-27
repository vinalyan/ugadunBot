package main

import (
	"flag"
	"log"
	tgClient "ugadunbot/internal/clients/telegram"
	event_consumer "ugadunbot/internal/consumer/event-consumer"
	"ugadunbot/internal/distributer"
	"ugadunbot/internal/events/telegram"
)

const (
	tgBotHost   = "api.telegram.org"
	googlehost  = "script.google.com"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	tgtoken, gglsheets := mustToken()
	eventsProcessor := telegram.New(tgClient.New(tgBotHost, tgtoken), *distributer.New(googlehost, gglsheets))

	log.Printf("Cервис запущен")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Cервис остановился", err)
	}

}

func mustToken() (string, string) {

	tgtoken := flag.String("tg-bot-token", "", "тут должен быть токен")
	gglsheets := flag.String("gglsheets", "", "тут должен быть токен")

	flag.Parse()

	if *tgtoken == "" {
		log.Fatal("Нет tgtoken токена")

	}

	if *gglsheets == "" {
		log.Fatal(" Нет gglsheets токена ")

	}
	return *tgtoken, *gglsheets

}

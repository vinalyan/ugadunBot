package event_consumer

import (
	"log"
	"time"
	"ugadunbot/internal/events"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

// работает в цикле
// если ошибка то бросает ее в лог
// если эвентов нет, что через секунду продолжает работать
// если
func (c Consumer) Start() error {
	for {
		gotEvent, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] Consumer Start error: %s", err.Error())
			continue
		}

		if len(gotEvent) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}
		if err := c.hadleEvents(gotEvent); err != nil {
			log.Printf(err.Error())

			continue
		}
	}
}

/*
Проблемы с фукнцией ниже
1. потеря событий: ретраи, возращение в хранилише, фолбэкб, подтверждение для фетчера,
2. ОБработка всей пачки. Остановка поле ошибки,
3. Праллельная обработка.
*/
func (c *Consumer) hadleEvents(events []events.Event) error {
	for _, event := range events {

		log.Printf("Получил новое событие %s", event)
		if err := c.processor.Process(event); err != nil {
			log.Printf("hadleEvents упал тут : %s", err.Error())

			continue
		}
	}
	return nil
}

package cards

import (
	"errors"
	"math/rand"
	"time"
	"ugadunbot/internal/lib/e"
)

type Card struct {
	Num   int    `json: "num"`
	Name  string `json: "name"`
	Value string `json: "value"`
}

type Cards struct {
	Cards []Card `json: "cards"`
}

var ErrNoSavedCards = errors.New("не сохраненнх карточек")

//выдает рандомную карточку.

func (c Cards) Random() (card *Card, err error) {
	defer func() { err = e.WrapIfErr("рандом поломался чет", err) }()

	if len(c.Cards) == 0 {
		return nil, ErrNoSavedCards
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(c.Cards))

	return &c.Cards[n], nil
}

package distributer

import (
	"encoding/json"
	cards "ugadunbot/internal/cards"
	client "ugadunbot/internal/clients/http"
	"ugadunbot/internal/lib/e"
)

//получить данные
type Distributer struct {
	host  string
	cards cards.Cards
	token string
	//client client.Client
}

func New(host string, token string) *Distributer {
	return &Distributer{
		host:  host,
		token: token,
		cards: cards.Cards{},
		//client: client.Client{},
	}
}

//получить данные
func (d Distributer) Data() (cards cards.Cards, err error) {
	defer func() { err = e.WrapIfErr("Distributer.Data", err) }()

	cl := client.New(d.host, d.token)

	data, err := cl.DoRequest("", nil)
	if err != nil {
		return d.cards, err
	}

	if err := json.Unmarshal(data, &d.cards); err != nil {
		return d.cards, err
	}

	return d.cards, nil
}

package distributer

import (
	"encoding/json"
	cards "ugadunbot/internal/cards"
	client "ugadunbot/internal/clients/http"
)

//получить данные
func Data(token string) (cards.Cards, error) {
	cl := client.New("script.google.com", token)

	data, err := cl.DoRequest("", nil)
	if err != nil {
		return cards.Cards{}, err
	}

	var cards cards.Cards
	if err := json.Unmarshal(data, &cards); err != nil {
		return cards, err
	}
	return cards, nil
}

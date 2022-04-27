package telegram

import (
	"log"
	"strings"
	"ugadunbot/internal/cards"
)

const (
	StartCmd  = "/start"
	HelpCmd   = "/help"
	NextCmd   = "/next"
	AnswerCmd = "/answer"
	ReloadCmd = "/reload"
)

//TODO: По ублюдски определяем клавиатуру

const (
	keyboardNext   = `{ "keyboard": [ [{"text": "/next"}]], "one_time_keyboard": true}`
	keyboardAnswer = `{ "keyboard": [ [{"text": "/answer"}]], "one_time_keyboard": true}`
	keyboardRemuve = `{ "remove_keyboard": true }`
	noKeyboard     = ""
)

//TODO подумать как это сделать нормальным
var rndCard = new(cards.Card)

//TODO тут надо правильно обработать ошибки
func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("Получена новая команда %s от %s", text, username)

	//TODO переделать это в более адкеватный вид.
	switch text {
	case StartCmd:
		err := p.startCmd(chatID)
		if err != nil {
			return err
		}
	case AnswerCmd:
		err := p.answerCmd(chatID)
		if err != nil {
			return err
		}
	case NextCmd:
		err := p.nextCmd(chatID)
		if err != nil {
			return err
		}

	case HelpCmd:
		p.tg.SendMessage(chatID, msgHelp, keyboardRemuve)
		err := p.nextCmd(chatID)
		if err != nil {
			return err
		}
	case ReloadCmd:
		err := p.startCmd(chatID)
		if err != nil {
			return err
		}
	default:
		p.tg.SendMessage(chatID, msgUnknownCommand, keyboardRemuve)
	}
	return nil
}

func (p *Processor) startCmd(chatID int) error {
	cards, err := p.distributer.Data()
	if err != nil {
		p.tg.SendMessage(chatID, "Что-то дико пошло не так", keyboardRemuve)
		return err
	}

	//TODO обязательно полечить эту историю с указателями

	p.cards = cards
	rndCard, err = p.cards.Random()
	if err != nil {
		p.tg.SendMessage(chatID, "Что-то дико пошло не так", keyboardRemuve)
		return err
	}
	p.tg.SendMessage(chatID, msgHello, keyboardRemuve)
	p.tg.SendMessage(chatID, rndCard.Value, keyboardAnswer)

	return nil
}

func (p *Processor) answerCmd(chatID int) error {

	p.tg.SendMessage(chatID, rndCard.Name, keyboardNext)
	return nil
}

func (p *Processor) nextCmd(chatID int) error {
	card, err := p.cards.Random()

	if err != nil {
		p.tg.SendMessage(chatID, "Что-то дико пошло не так", keyboardRemuve)
		return err
	}
	rndCard = card
	p.tg.SendMessage(chatID, rndCard.Value, keyboardAnswer)
	return nil
}

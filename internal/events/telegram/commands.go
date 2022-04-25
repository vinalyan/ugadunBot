package telegram

import (
	"log"
	"strings"
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

//TODO тут надо правильно обработать ошибки
func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("Получена новая команда %s от %s", text, username)

	//проверяем, что это комнада добавления
	switch text {
	case StartCmd:
		p.tg.SendMessage(chatID, msgHello, keyboardRemuve)
		p.tg.SendMessage(chatID, msgQuestion, keyboardAnswer)
	case AnswerCmd:
		p.tg.SendMessage(chatID, msgAnswer, keyboardNext)
	case NextCmd:
		p.tg.SendMessage(chatID, msgQuestion, keyboardAnswer)
	case HelpCmd:
		p.tg.SendMessage(chatID, msgHelp, keyboardRemuve)
		p.tg.SendMessage(chatID, msgAnswer, keyboardNext)
		//TODO тут прям затык будет надо залепить
	case ReloadCmd:
		p.tg.SendMessage(chatID, msgReload, keyboardRemuve)
		p.tg.SendMessage(chatID, msgFirstCard, noKeyboard)
		p.tg.SendMessage(chatID, msgQuestion, keyboardAnswer)
	default:
		p.tg.SendMessage(chatID, msgUnknownCommand, keyboardRemuve)
	}
	return nil
}

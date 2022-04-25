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
	keyboardNext   = `{ "keyboard": [ [{"text": "/next"}], "one_time_keyboard": true}`
	reyboardAnswer = `{ "keyboard": [ [{"text": "/answer"}], "one_time_keyboard": true}`
	noKeyboard     = ""
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("Получена новая команда %s от %s", text, username)

	//проверяем, что это комнада добавления
	switch text {
	case StartCmd:
		return p.sendHello(chatID, noKeyboard)
	case HelpCmd:
		return p.sendHelp(chatID, noKeyboard)
	case NextCmd:
		return p.sendNext(chatID, keyboardNext)
	case AnswerCmd:
		return p.sendAnswer(chatID, reyboardAnswer)
	case ReloadCmd:
		return p.sendReload(chatID, noKeyboard)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand, noKeyboard)

	}
}

//TODO если никакой допобработки не понадобится, то оставлю только одну функцию обработки собобщения
func (p *Processor) sendHello(chatID int, keyboard string) error {
	return p.tg.SendMessage(chatID, msgHello, keyboard)
}

func (p *Processor) sendHelp(chatID int, keyboard string) error {
	return p.tg.SendMessage(chatID, msgHello, keyboard)
}

func (p *Processor) sendAnswer(chatID int, keyboard string) error {
	return p.tg.SendMessage(chatID, msgHello, keyboard)
}

func (p *Processor) sendNext(chatID int, keyboard string) error {
	return p.tg.SendMessage(chatID, msgHello, keyboard)
}

func (p *Processor) sendReload(chatID int, keyboard string) error {
	return p.tg.SendMessage(chatID, msgHello, keyboard)
}

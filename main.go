package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	bot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// /start
	bot.Handle("/start", func(c tele.Context) error {
		log.Printf("[/start] user=%d username=@%s", c.Sender().ID, c.Sender().Username)
		return c.Send("Привет! 👋 Я тестовый бот на Go.\n\nДоступные команды:\n/start — приветствие\n/help — помощь\n/echo — повтори текст")
	})

	// /help
	bot.Handle("/help", func(c tele.Context) error {
		log.Printf("[/help] user=%d username=@%s", c.Sender().ID, c.Sender().Username)
		return c.Send("Список команд:\n/start — начать\n/help — помощь\n/echo <текст> — бот повторит твой текст")
	})

	// /echo
	bot.Handle("/echo", func(c tele.Context) error {
		args := c.Args()
		log.Printf("[/echo] user=%d args=%v", c.Sender().ID, args)
		if len(args) == 0 {
			return c.Send("Напиши что-нибудь после /echo, например: /echo Привет мир")
		}
		text := ""
		for _, arg := range args {
			text += arg + " "
		}
		return c.Send(text)
	})

	// любое текстовое сообщение
	bot.Handle(tele.OnText, func(c tele.Context) error {
		log.Printf("[message] user=%d username=@%s text=%q", c.Sender().ID, c.Sender().Username, c.Text())
		return c.Send("Ты написал: " + c.Text() + "\n\nИспользуй /help чтобы увидеть команды.")
	})

	log.Println("Бот запущен...")
	bot.Start()
}

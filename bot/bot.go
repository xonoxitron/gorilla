package bot

import (
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"

	"github.com/xonoxitron/gorilla/config"
	"github.com/xonoxitron/gorilla/handlers"
	"github.com/xonoxitron/gorilla/storage"
	"github.com/xonoxitron/gorilla/utils"
)

var bot *tb.Bot
var err error

func Send(receiver tb.Recipient, message string) {
	bot.Send(receiver, message)
}

func Notify(message string) {
	subscribers := strings.Split(storage.Get("subscribers"), "\r")

	for _, s := range subscribers {
		if s != "" {
			id, err := strconv.Atoi(s)
			utils.ErrorCheck(err)
			go Send(tb.Recipient(&tb.User{ID: id}), message)
		}
	}
}

func Start() {
	bot, err = tb.NewBot(tb.Settings{
		Token:  config.Get().Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	utils.ErrorCheck(err)

	bot.Handle("/start", func(m *tb.Message) {
		Send(m.Sender, handlers.Start())
	})

	bot.Handle("/subscribe", func(m *tb.Message) {
		Send(m.Sender, handlers.SubscribeUser(strconv.FormatInt(m.Chat.ID, 10)))
	})

	bot.Handle("/unsubscribe", func(m *tb.Message) {
		Send(m.Sender, handlers.UnsubscribeUser(strconv.FormatInt(m.Chat.ID, 10)))
	})

	bot.Start()
}

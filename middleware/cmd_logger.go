package middleware

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func CommandLogger(handlerFunc tele.HandlerFunc) tele.HandlerFunc {
	return func(ctx tele.Context) error {
		var command string
		var _type string
		switch {
		case ctx.Callback() != nil:
			command = ctx.Callback().Unique
			_type = "callback"
		case ctx.Message() != nil:
			command = ctx.Message().Text
			_type = "text"
		default:
			command = ctx.Text()
			_type = "default"
		}

		start := time.Now()
		err := handlerFunc(ctx)
		log.Printf("chat id: %d, username: %s, command: %s, took %v to complete, type: %s", ctx.Chat().ID, ctx.Chat().Username, command, time.Since(start), _type)
		return err
	}
}

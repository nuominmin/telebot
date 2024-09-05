package middleware

import (
	tele "gopkg.in/telebot.v3"
	"log"
)

func ErrorRespond(handlerFunc tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		resp := &tele.CallbackResponse{}
		if err := handlerFunc(c); err != nil {
			resp.Text = err.Error()
			resp.ShowAlert = true
		}

		err := c.Respond(resp)
		if err != nil {
			log.Printf("telebot/middleware/ErrorRespond, error: %s", err.Error())
		}
		return nil
	}
}

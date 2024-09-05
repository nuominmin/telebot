package middleware

import (
	tele "gopkg.in/telebot.v3"
)

func ErrorRespond(handlerFunc tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		if err := handlerFunc(c); err != nil {
			return c.Respond(&tele.CallbackResponse{
				Text:      err.Error(),
				ShowAlert: true,
			})
		}

		return c.Respond(&tele.CallbackResponse{})
	}
}

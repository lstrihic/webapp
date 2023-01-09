package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lstrihic/webapp/port/http/middleware/localization"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// LocalizeMessage localize message.
func LocalizeMessage(ctx *fiber.Ctx, messageID, defaultMessage string) string {
	localizer, ok := ctx.Locals(localization.DefaultContextKey).(*i18n.Localizer)
	if !ok {
		return defaultMessage
	}
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
	if err != nil {
		return defaultMessage
	}
	return message
}

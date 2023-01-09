package localization

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/fx"
	"golang.org/x/text/language"
)

var Provider = fx.Provide(InitI18nService)

type Service interface {
	GetBundle() *i18n.Bundle
}

type service struct {
	bundle *i18n.Bundle
}

//go:embed messages/*.toml
var LocaleFS embed.FS

// InitI18nService initialize localization.
func InitI18nService() Service {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, _ = bundle.LoadMessageFileFS(LocaleFS, "messages/active.en.toml")
	_, _ = bundle.LoadMessageFileFS(LocaleFS, "messages/active.hr.toml")
	return &service{bundle: bundle}
}

func (s *service) GetBundle() *i18n.Bundle {
	return s.bundle
}

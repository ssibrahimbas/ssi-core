package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type I18n struct {
	b  *i18n.Bundle
	fb string
}

func New(fb string) *I18n {
	b := i18n.NewBundle(language.English)
	return &I18n{
		b:  b,
		fb: fb,
	}
}

func (i *I18n) LoadLanguages(ld string, languages ...string) {
	i.b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, lang := range languages {
		i.b.MustLoadMessageFile(ld + "/" + lang + ".toml")
	}
}

func (i *I18n) translate(c *i18n.LocalizeConfig, languages ...string) string {
	localizer := i18n.NewLocalizer(i.b, languages...)
	return localizer.MustLocalize(c)
}

func (i *I18n) Translate(key string, languages ...string) string {
	return i.translate(&i18n.LocalizeConfig{
		MessageID: key,
	}, languages...)
}

func (i *I18n) TranslateWithParams(key string, params interface{}, languages ...string) string {
	return i.translate(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: params,
	}, languages...)
}

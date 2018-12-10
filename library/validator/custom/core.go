package custom

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/gosagawa/realize_sample/library/log"
	"go.uber.org/zap"
	validator "gopkg.in/go-playground/validator.v9"
)

// Validation カスタムバリデーション処理インタフェース
type Validation interface {
	GetTagName() string
	Register(val *validator.Validate, trans ut.Translator) error
}

// customMap カスタムバリデーション
var customMap = make(map[string]Validation)

// RegisterCustomValidations カスタムバリデーションをまとめて登録する
func RegisterCustomValidations(val *validator.Validate, trans ut.Translator) error {

	for _, v := range customMap {
		if err := v.Register(val, trans); err != nil {
			return err
		}
		log.Logger.Debug("registered custom validation", zap.String("tagName", v.GetTagName()))
	}

	return nil
}

// registrationFunc カスタムバリデーションのメッセージ変換ロジック登録用の汎用関数
// github.com/go-playground/locales/en.go の RegisterDefaultTranslations の実装を参考に
func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

// translateFunc カスタムバリデーションのメッセージ変換ロジック登録用の汎用関数
// github.com/go-playground/locales/en.go の RegisterDefaultTranslations の実装を参考に
func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}

	return t
}

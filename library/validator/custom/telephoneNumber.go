package custom

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

const tagName = "tel"

var reg *regexp.Regexp

func init() {

	var err error
	reg, err = regexp.Compile(`^[0-9]+-[0-9]+-[0-9]+$`)

	if err != nil {
		panic(err)
	}

	// カスタムバリデーションを登録
	customMap[tagName] = &TelephoneNumber{}
}

// TelephoneNumber カスタムバリデーション：電話番号
type TelephoneNumber struct {
}

// GetTagName カスタムバリデーションのタグを取得する
func (v *TelephoneNumber) GetTagName() string {
	return tagName
}

// Register カスタムバリデーションやメッセージ変換ロジックを登録する
func (v *TelephoneNumber) Register(val *validator.Validate, trans ut.Translator) error {

	if err := val.RegisterValidation(tagName, isTelephoneNumber); err != nil {
		return err
	}

	// カスタム定義のtranslationを追加
	// https://github.com/go-playground/validator/blob/v9/translations/en/en.go の RegisterDefaultTranslations の実装を参考に
	if err := val.RegisterTranslation(tagName, trans, registrationFunc("tel", "{0} must be a telephone number", false), translateFunc); err != nil {
		return err
	}

	return nil
}

// isTelephoneNumber 入力が電話番号かどうかを検証する
func isTelephoneNumber(fl validator.FieldLevel) bool {

	v := fl.Field().String()
	if v == "" {
		return true
	}

	return reg.MatchString(v)
}

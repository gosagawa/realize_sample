package validator

import (
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/gosagawa/realize_sample/library/log"
	"github.com/gosagawa/realize_sample/library/validator/custom"
	validator "gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/en"
)

// FieldError 検証NGの情報を持った構造体
type FieldError struct {
	validator.FieldError
	ErrorMessage string
}

var val *validator.Validate
var trans ut.Translator

func init() {

	log.Logger.Debug("begin init validation")

	// Translatorは、公式のテストコードを参考にして実装
	// https://github.com/go-playground/validator/blob/v9/translations/en/en_test.go
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ = uni.GetTranslator("en")

	val = validator.New()

	if err := en.RegisterDefaultTranslations(val, trans); err != nil {
		panic("デフォルトのtranslation情報登録失敗: " + err.Error())
	}

	if err := custom.RegisterCustomValidations(val, trans); err != nil {
		panic("カスタムvalidation登録失敗: " + err.Error())
	}

	log.Logger.Debug("end init validation")
}

// Validate 引数の構造体のデータ検証を実施する
func Validate(v interface{}) ([]FieldError, error) {

	if err := val.Struct(v); err != nil {

		// 検証NGだった場合は、ValidationErrors構造体が返ってくるはず
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			var fieldErrors []FieldError
			for _, v := range ve {
				fe := FieldError{
					FieldError:   v,
					ErrorMessage: v.Translate(trans),
				}
				fieldErrors = append(fieldErrors, fe)
			}

			// 検証NG
			return fieldErrors, nil
		}

		// 検証中にエラーが発生
		return nil, err
	}

	// 検証OK
	return nil, nil
}

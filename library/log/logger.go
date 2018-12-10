package log

import (
	"io/ioutil"

	"github.com/gosagawa/realize_sample/library/env"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

// Logger ロガー
var Logger *zap.Logger

func init() {

	configYaml, err := ioutil.ReadFile(env.GetConfigPath("log"))
	if err != nil {
		panic("logger設定読み込み失敗: " + err.Error())
	}

	var myConfig zap.Config
	if err := yaml.Unmarshal(configYaml, &myConfig); err != nil {
		panic("loggerのyaml変換失敗: " + err.Error())
	}

	Logger, err = myConfig.Build()
	if err != nil {
		panic("loggerビルド失敗: " + err.Error())
	}
}

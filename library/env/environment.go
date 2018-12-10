package env

import (
	"os"
)

var appEnv, configPath string

func init() {

	appEnv = os.Getenv("APP_ENV")

	if IsLocal() || IsTravis() {
		configPath = os.Getenv("GOPATH") + "/src/github.com/gosagawa/" + GetAppName() + "/config"
	} else {
		configPath = "/home/lapisum/config/" + GetAppName()
	}
}

// GetAppName アプリ名を取得する
func GetAppName() string {
	return "realize_sample"
}

// GetConfigPath 設定ファイル置き場所のパスを取得する
func GetConfigPath(libraryName string) string {
	return configPath + "/" + libraryName + "/" + GetAppEnv() + ".yaml"
}

// GetAppEnv 現在の実行環境を取得する
func GetAppEnv() string {
	return appEnv
}

// IsLocal 現在の実行環境がローカル環境かどうか
func IsLocal() bool {
	return appEnv == "local"
}

// IsTravis 現在の実行環境がTravisCI環境であるかどうか
func IsTravis() bool {
	return appEnv == "travis"
}

// IsDevelopment 現在の実行環境が開発環境かどうか
func IsDevelopment() bool {
	return appEnv == "development"
}

// IsProduction 現在の実行環境が本番環境かどうか
func IsProduction() bool {
	return appEnv == "production"
}

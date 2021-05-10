package message

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	"github.com/line/line-bot-sdk-go/linebot"
)

// トップページ表示処理
func TopPage(w http.ResponseWriter, r *http.Request) {
	// トップページを指定
	temp, err := template.ParseFiles("resources/templates/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err = temp.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

// 構造体：LINE接続設定を管理
type LineConfig struct {
	UserId             string `json:"userId"`
	ChannelSecret      string `json:"channelSecret"`
	ChannelAccessToken string `json:"channelAccessToken"`
}

// LINEbotメッセージ送信処理
func SendMessage(w http.ResponseWriter, r *http.Request) {
	// リクエスト値を取得
	var message string = r.FormValue("message")
	// プロパティファイル読み込み
	var config LineConfig = LoadConfig()
	// LINEbotオブジェクト初期化（チャネルシークレット、第二引数にチャネルアクセストークンを指定）
	bot, _ := linebot.New(config.ChannelSecret, config.ChannelAccessToken)
	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewTextMessage(message))
	// メッセージ送信（第一引数に宛先、第二引数に送信するメッセージを指定）
	_, err := bot.PushMessage(config.UserId, messages...).Do()
	if err != nil {
		panic(err.Error())
	}
}

func LoadConfig() LineConfig {
	// 設定ファイルの読み込み
	file, err := os.Open("resources/conf/config.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var config LineConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return config
}

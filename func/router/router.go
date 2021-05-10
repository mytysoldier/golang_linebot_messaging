package router

import (
	"net/http"

	"github.com/mytysoldier/golang_linebot_messaging/func/message"
)

// パスルーティングを定義
func Router() {
	// トップページ表示のパスルーティング
	http.HandleFunc("/", message.TopPage)
	http.Handle("/resources/js/", http.FileServer(http.Dir("./")))
	// LINEbot送信のパスルーティング
	http.HandleFunc("/send", message.SendMessage)
	// 8080ポートで公開
	http.ListenAndServe(":8080", nil)
}

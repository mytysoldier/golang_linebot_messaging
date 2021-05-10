# golang_linebot_messaging

Web 画面上から入力したメッセージを、LINEbot から個人の LINE に通知する簡易お試しアプリ

### ☆ 動作確認事前準備

1. [golang 公式サイト](https://golang.org)から Go をインストールしてください
1. [Docker 公式サイト](https://www.docker.com)から Docker をインストールしてください
1. [LINE Developers](https://developers.line.biz/ja/)にて LINEbot を作成してください
1. LINEbot を作成できたら、ご自身の LINE アカウントに友達追加しておいてください
1. 「golang_linebot_messaging/resources/conf/config.json」を設定してください

   - userId：LINE Developers 上で「あなたのユーザー ID」欄から確認できます
   - channelSecret：LINE Developers 上で「チャネルシークレット」欄から確認できます
   - channelAccessToken：LINE Developers 上の「チャネルアクセストークン」欄からトークンを発行し、発行したトークンを指定してください

### ☆ 動作確認手順

以下コマンドを順に実行し、動作確認できることを確認。  
Dockerfile と同一ディレクトリにてターミナル起動。

1. `docker build -t golang_linebot_messaging .`  
   を実行し Docker イメージを作成
1. `docker image ls`  
   にて「golang_linebot_messaging」という名前の Docker イメージが作成されることを確認
1. `docker run -p 8080:8080 -td --name golang_linebot_messaging golang_linebot_messaging`  
   を実行し、Docker コンテナ起動
1. `docker ps`にて「golang_linebot_messaging」という名前のコンテナが起動していることを確認
1. http://localhost:8080 にブラウザからアクセス、トップ画面が表示されることを確認
1. テキストボックスに任意のテキストを入力し「メッセージ送信」ボタン押下、自分の LINE に LINEbot からメッセージが届いていることを確認

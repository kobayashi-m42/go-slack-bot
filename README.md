# go-slack-bot

Slackに投稿されたメッセージからGitHubの特定のリポジトリにIssueを作成するBotです。

特定のチャンネルでBot宛てのメッセージを送信してください。

投稿されたメッセージからIssueが作成されます。

なお、リポジトリは固定で事前に環境変数に設定する必要があります。

## Slack AppにBotsを追加する
下記のページより、Slack Appに Botsを追加してください。

https://slack.com/apps/A0F7YS25R-bots

追加すると、APIトークンが発行されます。

## GitHub Tokenの発行
下記のページを参考にGitHubのトークンを発行します。

https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line

## 環境変数の設定
`.env`を作成します。

```
SLACK_API_TOKEN=your_slack_api_token // 上記で取得したSlack APIのトークン 
BOT_ID=your_bot_id
CHANNEL_ID=your_channel_id
GITHUB_TOKEN=your_github_token // 上記で取得したGitHubのトークン
GITHUB_OWNER=your_github_owner
GITHUB_REPOSITORY=your_github_repository // Issueを作成するリポジトリ
```

## Dockerでの実行

### 初回起動
```
docker-compose up --build -d
```

### 2回目以降の起動
```
docker-compose up
```

## 本番環境での実行
実行時に環境変数の設定ファイルを指定する必要があります。

```
docker build -t go-slack-bot .
docker run -d --env-file ./.env --name go-slack-bot go-slack-bot
```

# go-slack-bot

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

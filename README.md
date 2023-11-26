# Стивен Кинг - бот цитатник

отвечает на ваш запрос (с именем автора), случайной цитатой из книг автора

## настройки

прописать в файле .env токен бота
если ваш бот имеет Webhook то его удалить

```
https://api.telegram.org/bot{TOKEN}/setWebhook?remove
```

## запуск

```
$ go mod init go_bot.go
$ go mod tidy
$ go run .
```


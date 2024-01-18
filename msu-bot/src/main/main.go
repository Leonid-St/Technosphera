package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
	"joke"
)

// для вендоринга используется GB
// сборка проекта осуществляется с помощью gb build
// установка зависимостей - gb vendor fetch gopkg.in/telegram-bot-api.v4
// установка зависимостей из манифеста - gb vendor restore


var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Get Joke"},
}

// При старте приложения, оно скажет телеграму ходить с обновлениями по этому URL
const WebhookURL = "https://msu-go-2017.herokuapp.com/"


func main() {
	// Heroku прокидывает порт для приложения в переменную окружения PORT
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("315298480:AAEccNOWsRU-lX00lP-oVSymNdRWhh5vVzc")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Устанавливаем вебхук
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	// получаем все обновления из канала updates
	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("received text: ", update.Message.Text)

		switch update.Message.Text {
		case "Get Joke":
			// Если пользователь нажал на кнопку, то придёт сообщение "Get Joke"
			message = tgbotapi.NewMessage(update.Message.Chat.ID, joke.GetJoke())
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Press "Get Joke" to receive joke`)
		}

		// В ответном сообщении просим показать клавиатуру
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		bot.Send(message)
	}
}

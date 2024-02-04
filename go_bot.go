package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TOKEN string

var bot *tgbotapi.BotAPI
var searchNames = [3]string{"стивен", "кинг", "стив"}
var chat_id int64
var answers = []string{
	"Люди не становятся лучше - только умнее. Они не перестают отрывать мухам крылышки, а лишь только придумывают себе гораздо более убедительные оправдания.",
	"«Никогда» - то самое слово, которое слушает Бог, когда хочет посмеяться (Темная башня)",
	"«Человек, который почувствовал ветер перемен, должен строить не щит от ветра, а ветряную мельницу» (Мертвая зона)",
	"Добрых снов, спокойной ночи, пусть приснится, что захочешь. (Позже)",
	"Если ты не будешь контролировать свой гнев, твой гнев будет контролировать тебя (Под Куполом)",
	"Плакать - это как вылить вон всю память. (Жребий Салема)",
	"Крест, хлеб и вино - только символы. Без веры крест - простое дерево, хлеб - испеченное зерно, а вино - сок винограда. (Жребий Салема)",
	"Господи, дай мне смирение принять то, что я не могу изменить, волю изменить то, что я не могу принять, и ум, что бы не слишком уж выеживаться. (Жребий Салема)",
	"не стоит беспокоиться о том, что не в твоей власти. Это верный путь к безумию. (Билли Саммерс)",
	"думать о собственном душевном здоровье было плохо. Все равно что думать о собственном сердцебиении: если ты о нем думаешь, значит, уже есть проблемы. (Чужак)",
	"Вместо того, чтобы ныть по поводу своих проблем, он продолжал спокойно жить, и старался сделать свою жизнь максимально приятной и интересной. (Рита Хейуорт, или Побег из Шоушенка)",
}

func init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "TOKEN=") {
			TOKEN = strings.ReplaceAll(s, "TOKEN=", "")
		}
	}
}

func connectWithTg() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to telegram!")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chat_id, msg)
	bot.Send(msgConfig)
}

func isMessageForTeller(update *tgbotapi.Update) bool {
	if update.Message == nil && update.Message.Text == "" {
		return false
	}

	msgUserText := strings.ToLower(update.Message.Text)
	for _, name := range searchNames {
		if strings.Contains(msgUserText, name) {
			return true
		}
	}

	return false
}

func getAnswers() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chat_id, getAnswers())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	// TODO
	connectWithTg()

	u := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(u) {
		chat_id = update.Message.Chat.ID

		if update.Message != nil && update.Message.Text == "/start" {
			sendMessage("Задай свой вопрос, назвав автора по имени." +
				" Ответом на вопрос будит цитата из его книг (\"Стивена Кинга\")")
		}

		if isMessageForTeller(&update) {
			sendAnswer(&update)
		}
	}
}

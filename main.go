package main

import (
	"github.com/semper-proficiens/devbriefs-meta/external/externalapi"
	"log"
	"log/slog"
	"os"
)

const defaultLanguage = "go"

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// StackExchange API calls
	seAllTimeQuestions, err := externalapi.SEGetTotalQuestions(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions:", err)
	}
	log.Println("StackExchange total questions to date:", seAllTimeQuestions)

	seAllQuestionsLastMonth, err := externalapi.SEGetTotalQuestionsFromLastMonth(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions last month:", err)
	}
	log.Println("StackExchange total questions last month:", seAllQuestionsLastMonth)

	seAllQuestionsThisMonth, err := externalapi.SEGetTotalQuestionsThisMonth(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions this month:", err)
	}
	log.Println("StackExchange total questions since beginning of month:", seAllQuestionsThisMonth)

	// GitHub REST API Calls
	ghReposCreatedAllTime, err := externalapi.GHGetReposCreatedAllTime(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get github repos created to date:", err)
	}
	log.Println("Github repos created in to date:", ghReposCreatedAllTime)

	ghReposCreatedThisMonth, err := externalapi.GHGetReposCreatedThisMonth(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get github repos created this month:", err)
	}
	log.Println("Github repos created this month:", ghReposCreatedThisMonth)

	ghReposCreatedLastMonth, err := externalapi.GHGetReposCreatedLastMonth(defaultLanguage)
	if err != nil {
		log.Fatal("Failed to get github repos created last month:", err)
	}
	log.Println("Github repos created last month:", ghReposCreatedLastMonth)
}

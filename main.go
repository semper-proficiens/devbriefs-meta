package main

import (
	"github.com/semper-proficiens/devbriefs-meta/external/externalapi"
	"log"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// StackExchange API calls
	seAllTimeQuestions, err := externalapi.SEGetTotalQuestions("go")
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions:", err)
	}
	log.Println("StackExchange total questions to day:", seAllTimeQuestions)

	seAllQuestionsLastMonth, err := externalapi.SEGetTotalQuestionsFromLastMonth("go")
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions last month:", err)
	}
	log.Println("StackExchange total questions last month:", seAllQuestionsLastMonth)

	seAllQuestionsThisMonth, err := externalapi.SEGetTotalQuestionsThisMonth("go")
	if err != nil {
		log.Fatal("Failed to get total StackExchange total questions this month:", err)
	}
	log.Println("StackExchange total questions since beginning of month:", seAllQuestionsThisMonth)

	// GitHub REST API Calls
	ghReposCreatedAllTime, err := externalapi.GHGetReposCreatedAllTime("go")
	if err != nil {
		log.Fatal("Failed to get github repos created all time:", err)
	}
	log.Println("Github repos created in total:", ghReposCreatedAllTime)

	ghReposCreatedThisMonth, err := externalapi.GHGetReposCreatedThisMonth("go")
	if err != nil {
		log.Fatal("Failed to get github repos created this month:", err)
	}
	log.Println("Github repos created this month:", ghReposCreatedThisMonth)

	ghReposCreatedLastMonth, err := externalapi.GHGetReposCreatedLastMonth("go")
	if err != nil {
		log.Fatal("Failed to get github repos created last month:", err)
	}
	log.Println("Github repos created last month:", ghReposCreatedLastMonth)
}

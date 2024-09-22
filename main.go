package main

import (
	"github.com/semper-proficiens/devbriefs-meta/external/externalapi"
	"log"
)

func main() {
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
}

package externalapi

import (
	"encoding/json"
	"fmt"
	utilsTime "github.com/semper-proficiens/go-utils/system/time"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseQuestionsURL = "https://api.stackexchange.com/2.3/questions"

type TotalResponse struct {
	Total int `json:"total"`
}

// SEGetTotalQuestions queries StackExchange (SE) API for the total questions in all time for a
// given programming language based on "tagged" attribute.
func SEGetTotalQuestions(language string) (int, error) {
	params := url.Values{}
	params.Add("order", "desc")
	params.Add("sort", "activity")
	params.Add("tagged", language)
	params.Add("site", "stackoverflow")
	params.Add("filter", "total")

	fullURL := fmt.Sprintf("%s?%s", baseQuestionsURL, params.Encode())
	//log.Println("Url:", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("error from get response:%s", err)
	}
	defer resp.Body.Close()

	var questions TotalResponse
	if err = json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return questions.Total, nil
}

// SEGetTotalQuestionsFromLastMonth queries StackExchange (SE) API for the total questions from last month for a
// given programming language based on "tagged" attribute.
func SEGetTotalQuestionsFromLastMonth(language string) (int, error) {
	firstDayOfLastMonth := utilsTime.GetFirstDayOfLastMonth()
	lastDayOfLastMonth := utilsTime.GetLastDayOfLastMonth()

	firstDayOfLastMonthStr := strconv.Itoa(int(firstDayOfLastMonth.Unix()))
	lastDayOfLastMonthStr := strconv.Itoa(int(lastDayOfLastMonth.Unix()))

	params := url.Values{}
	params.Add("order", "desc")
	params.Add("sort", "activity")
	params.Add("tagged", language)
	params.Add("site", "stackoverflow")
	params.Add("fromdate", firstDayOfLastMonthStr) // expects Unix format
	params.Add("todate", lastDayOfLastMonthStr)    // expects Unix format
	params.Add("filter", "total")

	fullURL := fmt.Sprintf("%s?%s", baseQuestionsURL, params.Encode())
	//log.Println("Url:", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("error from get response:%s", err)
	}
	defer resp.Body.Close()

	var questions TotalResponse
	if err = json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return questions.Total, nil
}

// SEGetTotalQuestionsThisMonth queries StackExchange (SE) API for the total questions this month for a
// given programming language based on "tagged" attribute.
func SEGetTotalQuestionsThisMonth(language string) (int, error) {
	firstDayOfThisMonth := utilsTime.GetFirstDayOfThisMonth()
	firstDayOfThisMonthStr := strconv.Itoa(int(firstDayOfThisMonth.Unix()))

	toDateStr := strconv.Itoa(int(time.Now().Unix()))

	params := url.Values{}
	params.Add("order", "desc")
	params.Add("sort", "activity")
	params.Add("tagged", language)
	params.Add("site", "stackoverflow")
	params.Add("fromdate", firstDayOfThisMonthStr) // expects Unix format
	params.Add("todate", toDateStr)                // expects Unix format
	params.Add("filter", "total")

	fullURL := fmt.Sprintf("%s?%s", baseQuestionsURL, params.Encode())
	//log.Println("Url:", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("error from get response:%s", err)
	}
	defer resp.Body.Close()

	var questions TotalResponse
	if err = json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return questions.Total, nil
}

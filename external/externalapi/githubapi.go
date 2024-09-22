package externalapi

import (
	"encoding/json"
	"fmt"
	utilTime "github.com/semper-proficiens/go-utils/system/time"
	"net/http"
	"time"
)

const baseRepositoriesURL = "https://api.github.com/search/repositories"

type GHTotalRepos struct {
	Total int `json:"total_count"`
}

func GHGetReposCreatedAllTime(language string) (int, error) {
	query := fmt.Sprintf("language:%s", language)

	fullURL := fmt.Sprintf("%s?q=%s", baseRepositoriesURL, query)
	//log.Println("Url:", fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var repos GHTotalRepos
	if err = json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return repos.Total, nil
}

func GHGetReposCreatedThisMonth(language string) (int, error) {
	firstDayOfThisMonth := utilTime.GetFirstDayOfThisMonth()
	firstDayOfThisMonthStr := firstDayOfThisMonth.Format("2006-01-02")

	todayDateStr := time.Now().Format("2006-01-02")

	query := fmt.Sprintf("language:%s+created:%s..%s", language, firstDayOfThisMonthStr, todayDateStr)

	fullURL := fmt.Sprintf("%s?q=%s&sort=stars&order=desc", baseRepositoriesURL, query)
	//log.Println("Url:", fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var repos GHTotalRepos
	if err = json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return repos.Total, nil
}

func GHGetReposCreatedLastMonth(language string) (int, error) {
	firstDayOfLastMonth := utilTime.GetFirstDayOfLastMonth()
	firstDayOfLastMonthStr := firstDayOfLastMonth.Format("2006-01-02")

	lastDayOfLastMonth := utilTime.GetLastDayOfLastMonth()
	lastDayOfLastMonthStr := lastDayOfLastMonth.Format("2006-01-02")

	query := fmt.Sprintf("language:%s+created:%s..%s", language, firstDayOfLastMonthStr, lastDayOfLastMonthStr)

	fullURL := fmt.Sprintf("%s?q=%s&sort=stars&order=desc", baseRepositoriesURL, query)
	//log.Println("Url:", fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var repos GHTotalRepos
	if err = json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return 0, fmt.Errorf("error from decode response:%s", err)
	}
	return repos.Total, nil
}

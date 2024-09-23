package externalapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GoogleTrendsResponse struct {
	InterestOverTime struct {
		TimelineData []struct {
			Date      string `json:"date"`
			Timestamp string `json:"timestamp"`
			Values    []struct {
				Query          string `json:"query"`
				Value          string `json:"value"`
				ExtractedValue int    `json:"extracted_value"`
			} `json:"values"`
		} `json:"timeline_data"`
	} `json:"interest_over_time"`
}

func GTGetGoogleSearchesFor(apiKey, language string) (int, error) {
	baseURL := "https://serpapi.com/search"
	params := url.Values{}
	params.Add("engine", "google_trends")
	params.Add("q", language)
	params.Add("hl", "en")
	params.Add("geo", "US")
	params.Add("data_type", "TIMESERIES")
	params.Add("cat", "31")
	params.Add("date", "today 1-m")
	params.Add("api_key", apiKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	//log.Println("Url:", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return 0, fmt.Errorf("error from get response: %s", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var gtResponse GoogleTrendsResponse
	if err = json.NewDecoder(resp.Body).Decode(gtResponse); err != nil {
		return 0, fmt.Errorf("error from decode response: %s", err)
	}

	// Calculate the total count of searches
	totalCount := 0
	for _, item := range gtResponse.InterestOverTime.TimelineData {
		for _, value := range item.Values {
			if value.Query == "go" {
				totalCount += value.ExtractedValue
			}
		}
	}

	return totalCount, nil
}

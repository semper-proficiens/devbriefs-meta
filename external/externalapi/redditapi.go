package externalapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type RedditPost struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
			} `json:"data"`
		} `json:"children"`
		After string `json:"after"`
	} `json:"data"`
}

func GetRedditToken() (string, error) {
	clientID := os.Getenv("REDDIT_CLIENT_ID")
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "MyRedditApp/0.1 by YourUsername")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Print the raw response body for debugging

	var tokenResponse TokenResponse
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	//log.Println("Reddit token response:", tokenResponse)

	return tokenResponse.AccessToken, nil
}

func GetRedditPosts(token string, after string) (RedditPost, error) {
	url := "https://oauth.reddit.com/r/golang/new?limit=100&raw_json=1"
	if after != "" {
		url += "&after=" + after
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RedditPost{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "DevBriefsAPI")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RedditPost{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RedditPost{}, err
	}

	var redditPost RedditPost
	if err := json.Unmarshal(body, &redditPost); err != nil {
		return RedditPost{}, fmt.Errorf("failed to decode posts response: %w", err)
	}

	// Print only the first two items for debugging
	//if len(redditPost.Data.Children) > 0 {
	//	for i, child := range redditPost.Data.Children {
	//		if i >= 2 {
	//			break
	//		}
	//		log.Printf("Post %d: %+v\n", i+1, child.Data)
	//	}
	//}

	return redditPost, nil
}

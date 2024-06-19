package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AIGateway interface {
	GetAIResponse(apiKey string, message string) (string, error)
}

type RealAIGateway struct{}

func (g *RealAIGateway) GetAIResponse(apiKey string, message string) (string, error) {

	// Using OpenAI api to get ChatGPT Response as example (It can be integrated with any AI/Model)
	requestBody, err := json.Marshal(map[string]string{
		"model":  "gpt-3.5-turbo",
		"prompt": message,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	fmt.Println("RESULT", result)

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if text, ok := choice["text"].(string); ok {
				fmt.Println("TEXT FROM CHAT GPT: ", text)
				return text, nil
			}
		}
	}

	return "", fmt.Errorf("unexpected response format")
}

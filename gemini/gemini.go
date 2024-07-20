package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"fuzztester-cli/models"
	"strconv"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateJSON(jsonData []byte) ([]models.FuzzTestCase, error) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyCEBU4MIMl2bMzBDR9ZPDjW-8k0JBVZEMM"))
	if err != nil {
		return nil, err
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-pro")
	model.GenerationConfig = genai.GenerationConfig{
		ResponseMIMEType: "application/json",
	}
	resp, err := model.GenerateContent(context.Background(), genai.Text(string(jsonData)))
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, err
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		return nil, err
	}

	unescapedData, err := strconv.Unquote(string(temp))
	if err != nil {
		return nil, err
	}

	var response []map[string]interface{}
	err = json.Unmarshal([]byte(unescapedData), &response)
	if err != nil {
		return nil, err
	}

	var cases []models.FuzzTestCase
	for _, item := range response {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		var testCase models.FuzzTestCase
		if err := json.Unmarshal(jsonData, &testCase); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		cases = append(cases, testCase)
	}

	return cases, nil
}

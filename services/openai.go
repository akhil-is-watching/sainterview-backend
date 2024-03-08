package services

import (
	"fmt"

	"github.com/akhil-is-watching/sainterview-backend/models"
	"github.com/akhil-is-watching/sainterview-backend/types"
	"github.com/go-resty/resty/v2"
)

type OpenAIService struct {
	Client *resty.Client
	APIUrl string
	APIKey string
}

var chat_service *OpenAIService

func InitOpenAIService(APIUrl string, APIKey string) {
	chat_service = &OpenAIService{
		Client: resty.New(),
		APIUrl: APIUrl,
		APIKey: APIKey,
	}
}

func ChatService() *OpenAIService {
	return chat_service
}

func (s *OpenAIService) GetQuestions(input models.Interview) (types.Questions, error) {
	response, err := s.Client.R().
		SetAuthToken(s.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": fmt.Sprintf("Create exactly %d questions for a technical interview for \n Job Role:%s\nJob Level: %s\nThe questions should be in json format [{\"question\": \"\"}]", input.QuestionCount, input.JobRole, input.JobLevel)}},
			"max_tokens": 1000,
		}).
		Post(s.APIUrl)

	if err != nil {
		return types.Questions{}, err
	}

	body := response.Body()
	openAIResponse, err := types.UnmarshalOpenAIResponse(body)
	if err != nil {
		return types.Questions{}, err
	}

	questions, err := types.UnmarshalQuestions([]byte(openAIResponse.Choices[0].Message.Content))
	if err != nil {
		return types.Questions{}, err
	}

	return questions, nil
}

package services

import (
	"github.com/akhil-is-watching/sainterview-backend/helpers"
	"github.com/akhil-is-watching/sainterview-backend/models"
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
)

type GeneratorService struct {
	HumanPalService *HumanPalService
	OpenAIService   *OpenAIService
}

var generator_service *GeneratorService

func InitGeneratorService(HumanPalService *HumanPalService, OpenAIService *OpenAIService) {
	generator_service = &GeneratorService{
		HumanPalService: HumanPalService,
		OpenAIService:   OpenAIService,
	}
}

func Generator() *GeneratorService {
	return generator_service
}

func (s *GeneratorService) Generate(input models.Interview) error {
	var questions []models.InterviewQuestion
	openai_response, err := s.OpenAIService.GetQuestions(input)
	if err != nil {
		return err
	}

	for _, qres := range openai_response {
		ID := helpers.UIDGen().GenerateID("IQ")
		question := models.InterviewQuestion{
			ID:          ID,
			Question:    qres.Question,
			InterviewID: input.ID,
		}

		err := s.HumanPalService.CreateQuestionVideo(ID, qres.Question, input.AvatarID)
		if err != nil {
			return err
		}

		questions = append(questions, question)
	}

	interviewQuestionRepo := repository.NewInterviewQuestionRepo(storage.GetDB())
	err = interviewQuestionRepo.CreateQuestion(questions)
	if err != nil {
		return err
	}

	return nil
}

// func (s *GeneratorService) Generate(input models.Interview) error {
// 	var wg sync.WaitGroup
// 	errChan := make(chan error, input.QuestionCount) // Buffered channel to prevent goroutine blocking
// 	interviewQuestionRepo := repository.NewInterviewQuestionRepo(storage.GetDB())

// 	openai_response, err := s.OpenAIService.GetQuestions(input)
// 	if err != nil {
// 		return err
// 	}

// 	for _, qres := range openai_response {
// 		wg.Add(1)
// 		go func(qres types.Question) {
// 			defer wg.Done()

// 			ID := helpers.UIDGen().GenerateID("IQ")
// 			question := models.InterviewQuestion{
// 				ID:          ID,
// 				Question:    qres.Question,
// 				InterviewID: input.ID,
// 			}

// 			if err := interviewQuestionRepo.CreateQuestion(question); err != nil {
// 				errChan <- err
// 				return
// 			}

// 			// Create question video in a Goroutine
// 			if err := s.HumanPalService.CreateQuestionVideo(ID, qres.Question, input.AvatarID); err != nil {
// 				errChan <- err
// 				return
// 			}

// 			// Append the question to the slice
// 		}(qres)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(errChan)
// 	}()

// 	for err := range errChan {
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

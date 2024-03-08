package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/akhil-is-watching/sainterview-backend/models"
	"github.com/akhil-is-watching/sainterview-backend/repository"
	"github.com/akhil-is-watching/sainterview-backend/storage"
	"github.com/akhil-is-watching/sainterview-backend/types"
	"github.com/go-resty/resty/v2"
)

type HumanPalService struct {
	Client *resty.Client
	APIUrl string
	APIKey string
}

var humanpal_service *HumanPalService

func InitHumanPalService(APIUrl string, APIKey string) {
	humanpal_service = &HumanPalService{
		Client: resty.New(),
		APIUrl: APIUrl,
		APIKey: APIKey,
	}
}

func VideoGeneratorService() *HumanPalService {
	return humanpal_service
}

func (s *HumanPalService) CreateQuestionVideo(ID string, question string, avatarID int) error {

	avatarRepo := repository.NewAvatarRepository(storage.GetDB())
	avatar, err := avatarRepo.GetByID(avatarID)
	if err != nil {
		return err
	}

	voice_data_request := types.NewVoiceDataRequest(avatar.ID, question, "")
	voice_data_response, err := s.Client.R().
		SetAuthToken(s.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(voice_data_request).
		Post(fmt.Sprintf("%s/text-to-audio", s.APIUrl))

	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := voice_data_response.Body()
	voice_data, err := types.UnmarshalVoiceDataResponse(body)
	if err != nil {
		return fmt.Errorf("%s. %s", err.Error(), string(body))
	}

	time.Sleep(12 * time.Second)
	audio_data_response, err := s.Client.R().
		SetAuthToken(s.APIKey).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/text-to-audio-chunks/%s", s.APIUrl, voice_data.Data))

	if err != nil {
		return err
	}

	body = audio_data_response.Body()
	audio_data, err := types.UnmarshalAudioDataResponse(body)
	if err != nil {
		return fmt.Errorf("%s. %s", err.Error(), string(body))
	}

	avatar_data := avatar.AvatarInput()

	create_video_request := types.NewVideoProject(ID, avatar_data, audio_data, "cheerful")
	resp, err := s.Client.R().
		SetAuthToken(s.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(create_video_request).
		Post(fmt.Sprintf("%s/videos/render-completed", s.APIUrl))

	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(string(resp.Body()))

	return nil
}

func (s *HumanPalService) SeedAvatars() {

	type AvatarAPIResponse struct {
		Data struct {
			CurrentPage int             `json:"current_page"`
			Data        []models.Avatar `json:"data"`
		} `json:"data"`
	}

	var avatar_resp AvatarAPIResponse

	resp, err := s.Client.R().
		SetAuthToken(s.APIKey).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/avatars", s.APIUrl))

	if err != nil {
		panic(err.Error())
	}

	body := resp.Body()
	err = json.Unmarshal(body, &avatar_resp)
	if err != nil {
		panic(err.Error())
	}

	avatarRepo := repository.NewAvatarRepository(storage.GetDB())
	err = avatarRepo.CreateAvatar(avatar_resp.Data.Data)
	if err != nil {
		panic(err.Error())
	}
}

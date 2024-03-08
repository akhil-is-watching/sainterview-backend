package types

import (
	"encoding/json"
)

func UnmarshalVoiceDataResponse(data []byte) (VoiceDataResp, error) {
	var r VoiceDataResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalAudioDataResponse(data []byte) (AudioData, error) {
	var r AudioData
	err := json.Unmarshal(data, &r)
	return r, err
}

type TextPartial struct {
	Title        string      `json:"title,omitempty"`
	VarName      string      `json:"var_name,omitempty"`
	Text         string      `json:"text,omitempty"`
	Index        int         `json:"index,omitempty"`
	OriginalText string      `json:"original_text,omitempty"`
	TTS          bool        `json:"tts,omitempty"`
	Source       interface{} `json:"source,omitempty"`
	Timestamps   interface{} `json:"timestamps,omitempty"`
}

type VoiceData struct {
	TTSLanguageID int           `json:"tts_language_id,omitempty"`
	TTSSpeakerID  int           `json:"tts_speaker_id,omitempty"`
	Type          int           `json:"type,omitempty"`
	Emotion       string        `json:"emotion,omitempty"`
	TextPartials  []TextPartial `json:"text_partials,omitempty"`
	Speed         int           `json:"speed,omitempty"`
	Pitch         int           `json:"pitch,omitempty"`
}

type VoiceDataResp struct {
	Data string `json:"data"`
}

type AudioChunk struct {
	Source       string        `json:"source"`
	VarName      string        `json:"var_name"`
	Duration     string        `json:"duration"`
	Timestamps   []interface{} `json:"timestamps"`
	Text         string        `json:"text"`
	OriginalText string        `json:"original_text"`
	DisplayText  interface{}   `json:"display_text"`
	Type         int           `json:"type"`
	SourcePath   string        `json:"source_path"`
}

type AudioData struct {
	Main struct {
		Source       string        `json:"source"`
		VarName      interface{}   `json:"var_name"`
		Duration     string        `json:"duration"`
		Timestamps   []interface{} `json:"timestamps"`
		Text         string        `json:"text"`
		OriginalText interface{}   `json:"original_text"`
		DisplayText  interface{}   `json:"display_text"`
		Type         int           `json:"type"`
		SourcePath   string        `json:"source_path"`
	} `json:"main"`
	Chunks struct {
		One AudioChunk `json:"1"`
	} `json:"chunks"`
}

type RenderSettings struct {
	ProjectName   string `json:"project_name"`
	Resolution    string `json:"resolution"`
	IsSocialMedia bool   `json:"is_social_media"`
	IsRenderToGIF bool   `json:"is_render_to_gif"`
}

type AvatarData struct {
	ID                int         `json:"id"`
	AvatarID          string      `json:"avatar_id"`
	OnlyShowOwned     int         `json:"only_show_owned"`
	IsFullBody        int         `json:"is_full_body"`
	MonthlyInfo       interface{} `json:"monthly_info"`
	RenderAvatarID    string      `json:"render_avatar_id"`
	IsVerticalPercent int         `json:"is_vertical_percent"`
	Name              string      `json:"name"`
	Gender            int         `json:"gender"`
	PreviewURL        string      `json:"preview_url"`
	Avatar            string      `json:"avatar"`
	NewUserOnly       int         `json:"new_user_only"`
	Thumbnail         string      `json:"thumbnail"`
	ThumbnailPath     string      `json:"thumbnail_path"`
	AvatarPath        string      `json:"avatar_path"`
}

type Avatar struct {
	Data       AvatarData  `json:"data"`
	CustomData interface{} `json:"customData"`
	Positions  []string    `json:"positions"`
}

type BackgroundGlobal struct {
	Title      string      `json:"title"`
	IsRequired bool        `json:"is_required"`
	Options    string      `json:"options"`
	Preview    string      `json:"preview"`
	IsHidden   bool        `json:"isHidden"`
	URL        string      `json:"url"`
	Default    string      `json:"default"`
	IsLoading  bool        `json:"is_loading"`
	VarName    string      `json:"var_name"`
	Value      string      `json:"value"`
	Type       string      `json:"type"`
	InputType  string      `json:"input_type"`
	IsVideo    bool        `json:"is_video"`
	IsCSV      bool        `json:"is_csv"`
	CSVTags    interface{} `json:"csv_tags"`
}

type Background struct {
	Dynamic []interface{}      `json:"dynamic"`
	Global  []BackgroundGlobal `json:"global"`
}

type Vars struct {
	Background Background `json:"background"`
	Global     []any      `json:"global"`
	Dynamic    []any      `json:"dynamic"`
	Mockup     []any      `json:"mockup"`
}

type OriginalText struct {
	Value   string `json:"value"`
	VarName string `json:"var_name"`
}

type WordPartial struct {
	Word1 []interface{} `json:"word_1"`
}

type Audio struct {
	VarName string `json:"var_name"`
	Value   string `json:"value"`
	Volume  int    `json:"volume"`
}

type AudioPartial struct {
	Audio1 string `json:"audio_1"`
}

type AudioDurationPartial struct {
	AudioDuration1 string `json:"audio_duration_1"`
}

type Voice struct {
	OriginalText          OriginalText         `json:"original_text"`
	TextPartials          []TextPartial        `json:"text_partials"`
	WordPartials          WordPartial          `json:"word_partials"`
	AudioPartials         AudioPartial         `json:"audio_partials"`
	AudioDurationPartials AudioDurationPartial `json:"audio_duration_partials"`
	Words                 []any                `json:"words"`
	Audio                 Audio                `json:"audio"`
	AudioDuration         string               `json:"audio_duration"`
	Emotion               string               `json:"emotion"`
	BackgroundAudio       Audio                `json:"background_audio"`
}

type ProjectAudio struct {
	Source        string                `json:"source"`
	VarName       interface{}           `json:"var_name"`
	Duration      string                `json:"duration"`
	Timestamps    []interface{}         `json:"timestamps"`
	Text          string                `json:"text"`
	OriginalText  interface{}           `json:"original_text"`
	DisplayText   interface{}           `json:"display_text"`
	Type          int                   `json:"type"`
	SourcePath    string                `json:"source_path"`
	Chunks        map[string]AudioChunk `json:"chunks"`
	TTSLanguageID int                   `json:"tts_language_id"`
	TTSSpeakerID  int                   `json:"tts_speaker_id"`
	IsCompleted   bool                  `json:"isCompleted"`
}

type VoiceActiveTab struct {
	Name       string `json:"name"`
	Component  string `json:"component"`
	Type       int    `json:"type"`
	IsDisabled bool   `json:"isDisabled"`
}

type StepData struct {
	ProjectAudio           ProjectAudio   `json:"projectAudio"`
	IsMulti                bool           `json:"isMulti"`
	ProjectBackgroundAudio struct{}       `json:"projectBackgroundAudio"`
	Speed                  int            `json:"speed"`
	Pitch                  int            `json:"pitch"`
	CustomText             string         `json:"customText"`
	STTLanguage            interface{}    `json:"sttLanguage"`
	IsSelectedAudio        bool           `json:"isSelectedAudio"`
	VoiceActiveTab         VoiceActiveTab `json:"voiceActiveTab"`
	IsUsedBackgroundAudio  bool           `json:"isUsedBackgroundAudio"`
}

type Data struct {
	VideoID        int            `json:"video_id"`
	IsSubtitle     bool           `json:"is_subtitle"`
	IsChangeScenes bool           `json:"is_change_scenes"`
	IsVertical     int            `json:"is_vertical"`
	RenderSettings RenderSettings `json:"render_settings"`
	Avatar         Avatar         `json:"avatar"`
	Vars           Vars           `json:"vars"`
	Voice          Voice          `json:"voice"`
	IsSend         bool           `json:"is_send"`
	StepData       StepData       `json:"stepData"`
}

func NewVoiceDataRequest(speaker_id int, text string, emotion string) VoiceData {
	return VoiceData{
		TTSLanguageID: 3,
		TTSSpeakerID:  speaker_id,
		Type:          2,
		Emotion:       emotion,
		TextPartials: []TextPartial{
			{
				Text: text,
			},
			{
				Title:        "Scene 1",
				VarName:      "text_1",
				Text:         text,
				Index:        1,
				OriginalText: text,
				TTS:          true,
			},
		},
		Speed: 0,
		Pitch: 0,
	}

}

func NewVideoProject(
	project_name string,
	avatar_data AvatarData,
	audio_data AudioData,
	emotion string,
) Data {
	var tts_speaker_id int
	if avatar_data.Gender == 1 {
		tts_speaker_id = 156
	} else {
		tts_speaker_id = 965
	}
	return Data{
		VideoID:        124,
		IsSubtitle:     false,
		IsChangeScenes: false,
		IsVertical:     0,
		RenderSettings: RenderSettings{
			ProjectName:   project_name,
			Resolution:    "jpeg-seq-720p",
			IsSocialMedia: false,
			IsRenderToGIF: false,
		},
		Avatar: Avatar{
			Data:       avatar_data,
			CustomData: nil,
			Positions:  []string{"center"},
		},
		Vars: Vars{
			Background: Background{
				Global: []BackgroundGlobal{
					{
						Title:      "Media",
						IsRequired: false,
						Options:    "",
						Preview:    "",
						IsHidden:   false,
						URL:        "https://humanavatar-live.s3.amazonaws.com/files/e3fd4d82c732fb6673315eac7d9d8200/gallery/images/2023_12_28_17_40_20__2691e45d2e42ddfaed5010920eecce3c.jpg",
						Default:    "https://humanavatar-live.s3.amazonaws.com/video-images/6290bbd22eab3/2022_09_13_18_16_16__aefd1313ce906d66f872418ad17a19a6.jpg",
						IsLoading:  false,
						VarName:    "sketch",
						Value:      "bg",
						Type:       "background_image",
						InputType:  "file",
						IsVideo:    false,
						IsCSV:      false,
						CSVTags:    nil,
					},
				},
				Dynamic: []any{},
			},
			Global:  []any{},
			Dynamic: []any{},
			Mockup:  []any{},
		},
		Voice: Voice{
			OriginalText: OriginalText{
				Value:   audio_data.Main.Text,
				VarName: "text",
			},
			TextPartials: []TextPartial{
				{
					Title:   "Scene 1",
					VarName: "text_1",
					Text:    audio_data.Main.Text,
					Index:   1,
				},
			},
			WordPartials: WordPartial{
				Word1: nil,
			},
			AudioPartials: AudioPartial{
				Audio1: audio_data.Main.SourcePath,
			},
			AudioDurationPartials: AudioDurationPartial{
				AudioDuration1: audio_data.Main.Duration,
			},
			Words: []any{},
			Audio: Audio{
				VarName: "voiceover_1",
				Value:   audio_data.Main.SourcePath,
				Volume:  50,
			},
			AudioDuration: audio_data.Main.Duration,
			Emotion:       emotion,
			BackgroundAudio: Audio{
				VarName: "background_audio_1",
				Value:   "",
				Volume:  20,
			},
		},
		IsSend: true,
		StepData: StepData{
			ProjectAudio: ProjectAudio{
				Source:       audio_data.Main.Source,
				VarName:      audio_data.Main.VarName,
				Duration:     audio_data.Main.Duration,
				Timestamps:   audio_data.Main.Timestamps,
				Text:         audio_data.Main.Text,
				OriginalText: audio_data.Main.OriginalText,
				DisplayText:  audio_data.Main.DisplayText,
				Type:         audio_data.Main.Type,
				SourcePath:   audio_data.Main.SourcePath,
				Chunks: map[string]AudioChunk{
					"1": audio_data.Chunks.One,
				},
				TTSLanguageID: 3,
				TTSSpeakerID:  tts_speaker_id,
				IsCompleted:   true,
			},
			IsMulti:         false,
			Speed:           0,
			Pitch:           0,
			CustomText:      audio_data.Main.Text,
			STTLanguage:     nil,
			IsSelectedAudio: false,
			VoiceActiveTab: VoiceActiveTab{
				Name:       "Text to Speech",
				Component:  "TTS",
				Type:       2,
				IsDisabled: false,
			},
			IsUsedBackgroundAudio: false,
		},
	}
}

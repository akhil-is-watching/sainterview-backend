package models

type InterviewQuestion struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Question    string `json:"question"`
	Audio       string `json:"audio"`
	Video       string `json:"video"`
	Status      string `gorm:"default:PENDING" json:"status"`
	InterviewID string `json:"interview_id"`
}

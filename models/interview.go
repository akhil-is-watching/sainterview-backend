package models

import "github.com/lib/pq"

type Interview struct {
	ID                 string              `gorm:"primaryKey" json:"id"`
	Industry           string              `json:"industry"`
	JobRole            string              `json:"job_role"`
	Designation        string              `json:"designation"`
	JobLevel           string              `json:"job_level"`
	AvatarID           int                 `json:"avatar_id"`
	Tags               pq.StringArray      `gorm:"type:text[]" json:"tags"`
	QuestionCount      int                 `json:"question_count"`
	InterviewQuestions []InterviewQuestion `json:"questions"`
	OrganizationID     string              `json:"organization_id"`
	Status             string              `gorm:"default:PENDING" json:"status"`
}

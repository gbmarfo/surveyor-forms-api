package main

import (
	"time"
)

type Survey struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title        string    `gorm:"size:255;not null;" json:"title"`
	Description  string    `gorm:"type:TEXT;" json:"description"`
	Introduction string    `gorm:"type:TEXT;" json:"introduction"`
	Consent      string    `gorm:"type:TEXT;" json:"consent"`
	CreatedBy    string    `gorm:"size:255;not null;" json:"created_by"`
	UpdatedBy    string    `gorm:"size:255;not null;" json:"updated_by"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Question struct {
	ID             uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Survey         Survey `json:"survey"`
	SurveyID       uint32 `gorm:"not null" json:"survey_id"`
	QuestionNum    uint32 `gorm:"not null" json:"question_num"`
	QuestionType   uint32 `gorm:"not null" json:"question_type"`
	QuestionLabel  string `gorm:"size:100;not null;" json:"question_label"`
	QuestionText   string `gorm:"size:100;not null;" json:"question_text"`
	AnswerRequired *bool  `gorm:"default=false" json:"answer_required"`
}

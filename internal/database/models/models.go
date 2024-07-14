package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID              uuid.UUID  `gorm:"type:uuid"`
	Ip              string     `gorm:"type:varchar(50)"`
	Method          string     `gorm:"type:varchar(10)"`
	Endpoint        string     `gorm:"type:varchar(255)"`
	RequestBody     *string    `gorm:"type:text"`
	RequestHeaders  *string    `gorm:"type:text"`
	RequestQuery    *string    `gorm:"type:text"`
	RequestParams   *string    `gorm:"type:text"`
	ResponseBody    *string    `gorm:"type:text"`
	ResponseHeaders *string    `gorm:"type:text"`
	ResponseTime    string     `gorm:"type:varchar(50)"`
	StatusCode      int        `gorm:"type:int"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

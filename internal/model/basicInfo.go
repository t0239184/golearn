package model

import (
	"time"
)

type BasicInfo struct {
	Creator  string    `gorm:"type: varchar(50) not null;"`
	CreateAt time.Time `gorm:"type: datetime not null; index:idx_create_at"`
	Updater  string    `gorm:"type: varchar(50);"`
	UpdateAt time.Time `gorm:"type: datetime;"`
	Deleter  string    `gorm:"type: varchar(50);"`
	DeleteAt time.Time `gorm:"type: datetime;"`
	IsDelete bool      `gorm:"type: tinyint(1) not null default 0;"`
}

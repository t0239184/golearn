package model

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/t0239184/golearn/internal/router/api/v1/request"
)

const (
	USER_STATUS_NORMAL = "0"
	USER_STATUS_LOCKED = "1"
)

type User struct {
	BasicInfo `gorm:"embedded"`
	ID        int64  `gorm:"type: bigint(20) not null auto_increment; primary_key;"`
	Account   string `gorm:"type: varchar(50) not null; index:idx_account,unique"`
	Password  string `gorm:"type: varchar(50) not null;"`
	Status    string `gorm:"type: varchar(1) not null default '0';"`
}

/* Costem table name */
func (t User) TableName() string {
	return "user"
}

/* GORM Hook */
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	logrus.Info("[BeforeCreate]")

	if (u.CreateAt == time.Time{}) {
		u.CreateAt = time.Now()
	}
	if u.Creator == "" {
		u.Creator = "system"
	}
	if u.Status == "" {
		u.Status = USER_STATUS_NORMAL
	}
	u.IsDelete = false
	return
}

/* GORM Hook */
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	logrus.Info("[BeforeSave]")
	if (u.UpdateAt == time.Time{}) {
		u.UpdateAt = time.Now()
	}
	if u.Updater == "" {
		u.Updater = "system"
	}
	return
}

/* GORM Hook */
func (u *User) AfterSave(tx *gorm.DB) (err error) {
	logrus.Info("[AfterSave] User Updated: ", u.ID)
	return
}

/* GORM Hook */
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	logrus.Info("[AfterCreate] User Created: ", u.ID)
	return
}

func NewUser(request *request.CreateUserRequest) *User {
	user := &User{
		Account:  request.Account,
		Password: request.Password,
		Status:   USER_STATUS_NORMAL,
		BasicInfo: BasicInfo{
			Creator:  "system",
			CreateAt: time.Now(),
			IsDelete: false,
		},
	}
	return user
}

func UpdateUser(request *request.UpdateUserRequest) *User {
	user := &User{
		ID:       request.ID,
		Password: request.Password,
	}
	return user
}

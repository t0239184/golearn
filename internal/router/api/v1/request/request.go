package request

import "time"

type CreateUserRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       int64     `json:"id"`
	Account  string    `json:"account"`
	Password string    `json:"password"`
	Status   string    `json:"status"`
	Creator  string    `json:"creator"`
	CreateAt time.Time `json:"create_at"`
	Updater  string    `json:"updater"`
	UpdateAt time.Time `json:"update_at"`
	Deleter  string    `json:"deleter"`
	DeleteAt time.Time `json:"delete_at"`
	IsDelete bool      `json:"is_delete"`
}

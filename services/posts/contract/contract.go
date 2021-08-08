package contract

import "time"

type PostModel struct {
	Id        string    `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
}

type UpdatePostReq struct {
	UserId  string
	Id      string
	Content string
}

type GetPostsReq struct {
	Id string `json:"id,omitempty"`
}

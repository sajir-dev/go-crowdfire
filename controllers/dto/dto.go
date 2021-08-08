package dto

type UpdatePostReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Id       string `json:"id,omitempty"`
	Content  string `json:"content,omitempty"`
}

type FollowReq struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Following string `json:"following,omitempty"`
}

type ListPostsReq struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

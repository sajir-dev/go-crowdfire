package contract

type UserModel struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateUser struct {
	Name     string
	Email    string
	Password string
}

type LoginReq struct {
	Email    string
	Password string
}

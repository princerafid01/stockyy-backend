package user

type UserCreateRequest struct {
	Email     string  `json:"email"`
	Password  *string `json:"password,omitempty"`
	GoogleID  *string `json:"google_id,omitempty"`
	Name      string  `json:"name"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type UserCreateResponse struct {
	ID        int64   `json:"id"`
	Email     string  `json:"email"`
	GoogleID  *string `json:"google_id,omitempty"`
	Name      string  `json:"name"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken string `json:"access_token"`
}

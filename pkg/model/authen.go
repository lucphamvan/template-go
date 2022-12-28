package model

type AuthenRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthenResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	User         User   `json:"user,omitempty"`
}

type RefreshTokenRequest struct {
	Token string `json:"token,omitempty"`
}

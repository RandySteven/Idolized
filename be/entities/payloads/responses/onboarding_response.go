package responses

import "time"

type (
	RegisterResponse struct {
		RequestID   string    `json:"request_id"`
		UserID      string    `json:"user_id"`
		Status      string    `json:"status"`
		VerifiedURL string    `json:"verified_url"`
		CreatedAt   time.Time `json:"created_at"`
	}

	LoginResponse struct {
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		LoginTime    time.Time `json:"login_time"`
	}

	OnboardUserResponse struct {
	}
)

package requests

type (
	OnboardingRequest struct {
		Type    string `json:"type"`
		Request any    `json:"request"`
	}

	OnboardingUser struct {
		FirstName string
		LastName  string
		Date      int
		Month     int
		Year      int
		Email     string
		Password  string
	}

	OnboardingTalent struct {
		FirstName string
		LastName  string
		Date      int
		Month     int
		Year      int
		Email     string
		Password  string
	}

	OnboardingGroup struct {
		FirstName string
		LastName  string
		Date      int
		Month     int
		Year      int
		Email     string
		Password  string
	}
)

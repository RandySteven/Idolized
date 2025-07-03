package requests

type (
	RegisterRequest struct {
		Type              string `json:"type"`
		OnboardingRequest any    `json:"request"`
	}

	OnboardingUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Date      string `json:"date"`
		Month     string `json:"month"`
		Year      string `json:"year"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	OnboardingTalent struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Date      string `json:"date"`
		Month     string `json:"month"`
		Year      string `json:"year"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	OnboardingGroup struct {
		FirstName     string
		LastName      string
		Date          int
		Month         int
		Year          int
		Email         string
		Password      string
		TalentMembers []OnboardingTalent
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

package validators

type IsValidUser struct {
	Token  string `json:"token" validate:"required"`
	UserId string `json:"userid" validate:"required"`
}

package validators

type IsValidStream struct {
	Secret   string `json:"secret" validate:"required"`
	StreamId string `json:"streamId" validate:"required,min=11,max=11"`
	AdminId  string `json:"adminId" validate:"required"`
}

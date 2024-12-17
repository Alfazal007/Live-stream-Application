package validators

type StartEndStream struct {
	StreamId string `json:"streamId" validate:"required,min=11,max=11"`
}

package managers

type Message struct {
	TypeOfMessage string `json:"typeOfMessage"`
	Message       string `json:"message"`
}

type AdminJoinType struct {
	AdminId string `json:"adminId"`
	RoomId  string `json:"roomId"`
	Token   string `json:"accessToken"`
}

type UserJoinType struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
	Token  string `json:"accessToken"`
}

type MessageType struct {
	RoomId   string `json:"roomId"`
	UserName string `json:"username"`
	Message  string `json:"message"`
}

var (
	JoinAdminMessage string = "JOINADMIN"
	JoinUserMessage  string = "JOINUSER"
	TextMessage      string = "TEXTMESSAGE"
)

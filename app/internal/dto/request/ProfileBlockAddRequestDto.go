package request

type ProfileBlockRequestDto struct {
	SessionId            string `json:"sessionId"`
	BlockedUserSessionId string `json:"blockedUserSessionId"`
}

package request

type BlockAddRequestDto struct {
	SessionId            string `json:"sessionId"`
	BlockedUserSessionId string `json:"blockedUserSessionId"`
}

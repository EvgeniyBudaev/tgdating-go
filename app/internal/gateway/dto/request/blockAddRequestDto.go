package request

type BlockRequestDto struct {
	SessionId            string `json:"sessionId"`
	BlockedUserSessionId string `json:"blockedUserSessionId"`
}

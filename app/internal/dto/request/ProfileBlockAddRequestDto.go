package request

type ProfileBlockRequestDto struct {
	SessionID            string `json:"sessionId"`
	BlockedUserSessionID string `json:"blockedUserSessionId"`
}

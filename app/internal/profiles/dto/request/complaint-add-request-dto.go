package request

type ComplaintAddRequestDto struct {
	SessionId         string `json:"sessionId"`
	CriminalSessionId string `json:"criminalSessionId"`
	Reason            string `json:"reason"`
}

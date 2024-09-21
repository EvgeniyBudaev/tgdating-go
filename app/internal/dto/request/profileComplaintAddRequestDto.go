package request

type ProfileComplaintAddRequestDto struct {
	SessionId         string `json:"sessionId"`
	CriminalSessionId string `json:"criminalSessionId"`
	Reason            string `json:"reason"`
}

package enum

type Complaint string

const (
	ComplaintFraud     Complaint = "fraud"
	ComplaintOther     Complaint = "Other"
	ComplaintSpam      Complaint = "spam"
	ComplaintTerrorism Complaint = "terrorism"
)

func (g Complaint) IsValid() bool {
	return g == ComplaintFraud || g == ComplaintOther || g == ComplaintSpam || g == ComplaintTerrorism
}

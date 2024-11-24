package ytrackercore

const (
	TicketKeyRegexp = `(?i)([A-Z]+-\d+)`
)

type Status string

const (
	PauseStatus    Status = "pause"
	DoneStatus     Status = "done"
	RejectedStatus Status = "rejected"
)

type Ticket struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	MR          string `json:"mr"`
	Customer    string `json:"customer"`
	Status      Status `json:"status"`
}

type TicketPatch struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	MR          *string `json:"mr"`
}

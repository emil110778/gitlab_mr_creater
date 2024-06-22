package ytrackercore

type Ticket struct {
	Self        string `json:"self"`
	Id          string `json:"id"`
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	MR          string `json:"mr"`
	//Priority    string `json:"priority"`
	//Status      string `json:"status"`
}

type TicketPatch struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	MR          *string `json:"mr"`
}

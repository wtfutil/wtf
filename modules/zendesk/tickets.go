package zendesk

import (
	"encoding/json"
	"log"
)

type TicketArray struct {
	Count         int    `json:"count"`
	Created       string `json:"created"`
	Next_page     string `json:"next_page"`
	Previous_page string `json:"previous_page"`
	Tickets       []Ticket
}

type Ticket struct {
	Id                    uint64      `json:"id"`
	URL                   string      `json:"url"`
	ExternalId            string      `json:"external_id"`
	CreatedAt             string      `json:"created_at"`
	UpdatedAt             string      `json:"updated_at"`
	Type                  string      `json:"type"`
	Subject               string      `json:"subject"`
	RawSubject            string      `json:"raw_subject"`
	Description           string      `json:"description"`
	Priority              string      `json:"priority"`
	Status                string      `json:"status"`
	Recipient             string      `json:"recipient"`
	RequesterId           uint64      `json:"requester_id"`
	SubmitterId           uint64      `json:"submitter_id"`
	AssigneeId            uint64      `json:"assignee_id"`
	OrganizationId        uint32      `json:"organization_id"`
	GroupId               uint32      `json:"group_id"`
	CollaboratorIds       []int64     `json:"collaborator_ids"`
	ForumTopicId          uint32      `json:"forum_topic_id"`
	ProblemId             uint32      `json:"problem_id"`
	HasIncidents          bool        `json:"has_incidents"`
	DueAt                 string      `json:"due_at"`
	Tags                  []string    `json:"tags"`
	Satisfaction_rating   string      `json:"satisfaction_rating"`
	Ticket_form_id        uint32      `json:"ticket_form_id"`
	Sharing_agreement_ids interface{} `json:"sharing_agreement_ids"`
	Via                   interface{} `json:"via"`
	Custom_Fields         interface{} `json:"custom_fields"`
	Fields                interface{} `json:"fields"`
}

func (widget *Widget) listTickets(pag ...string) (*TicketArray, error) {
	tickets := &TicketArray{}

	resource, err := widget.api("GET")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resource.Raw), tickets)
	if err != nil {
		return nil, err
	}

	return tickets, err
}

func (widget *Widget) newTickets() (*TicketArray, error) {
	newTicketArray := &TicketArray{}
	tickets, err := widget.listTickets(widget.settings.apiKey)
	if err != nil {
		log.Fatal(err)
	}
	for _, Ticket := range tickets.Tickets {
		if Ticket.Status == widget.settings.status && Ticket.Status != "closed" && Ticket.Status != "solved" {
			newTicketArray.Tickets = append(newTicketArray.Tickets, Ticket)
		}
	}

	return newTicketArray, nil
}

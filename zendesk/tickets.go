package zendesk

import (
	"encoding/json"
	"fmt"
	"log"
)

type TicketArray struct {
	Count         int    `json:"count"`
	Created       string `json:"created"`
	Next_page     string `json:"next_page"`
	Previous_page string `json:"previous_page"`
	Tickets       []Ticket
}

type SingleTicket struct {
	Ticket *Ticket `json:"ticket"`
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
	RequesterId           uint32      `json:"requester_id"`
	SubmitterId           uint32      `json:"submitter_id"`
	AssigneeId            uint32      `json:"assignee_id"`
	OrganizationId        uint32      `json:"organization_id"`
	GroupId               uint32      `json:"group_id"`
	CollaboratorIds       []int32     `json:"collaborator_ids"`
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

//var a = os.Getenv("ZENDESK_API")

func listTickets(pag ...string) (*TicketArray, error) {

	TicketStruct := &TicketArray{}

	var path string
	if len(pag) < 1 {
		path = "/tickets.json"
	} else {
		path = pag[0]
	}
	resource, err := api(a, "GET", path, "")
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(resource.Raw), TicketStruct)

	return TicketStruct, err

}

func newTickets() ([]string, error) {
	var newTickets []string
	tickets, err := listTickets()
	if err != nil {
		log.Fatal(err)
	}
	for i := range tickets.Tickets {
		if tickets.Tickets[i].Status == "new" {
			requester := tickets.Tickets[i].Via
			req, _ := requester.(map[string]interface{})
			source := req["source"]
			fromMap, _ := source.(map[string]interface{})
			from := fromMap["from"]
			fromValue, _ := from.(map[string]interface{})
			name := fromValue["name"]

			newTicket := fmt.Sprintf("%v - %v - %v - %v", tickets.Tickets[i].Id, tickets.Tickets[i].Status, tickets.Tickets[i].Subject, name)
			newTickets = append(newTickets, newTicket)
		}
	}
	if len(newTickets) < 1 {
		fmt.Println("No unassigned tickets in queue - woop!!")
	} else {
		for i := range newTickets {
			fmt.Println(newTickets[i])
		}
	}
	return newTickets, nil
}

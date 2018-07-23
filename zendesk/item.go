package zendesk

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

type Item struct {
	Requester string
	TicketID  string
	Subject   string
	Status    string
}

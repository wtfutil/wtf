package mail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/rivo/tview"
	log "github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/view"
)

type IMAPClient interface {
	List(ref, name string, mailboxes chan *imap.MailboxInfo) error
	Select(name string, readOnly bool) (*imap.MailboxStatus, error)
	Fetch(set *imap.SeqSet, items []imap.FetchItem, messages chan *imap.Message) error
	Logout() error
	Login(username, password string) error
	Close() error
}

// Widget is the container for your module's data
type Widget struct {
	view.ScrollableWidget

	settings       *Settings
	config         *Config
	client         IMAPClient
	currentMailbox *imap.MailboxStatus
	loggedIn       bool
	clientError    error
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	c, err := client.DialTLS(settings.imapAddress, nil)

	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.common),
		settings:         settings,
		config: &Config{
			page:     0,
			pageSize: uint32(settings.defaultPageSize),
		},
		client:      c,
		clientError: err,
		loggedIn:    err == nil,
	}

	if err == nil {
		widget.login()
	}

	if err == nil {
		widget.selectMailbox("INBOX")
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {

	// The last call should always be to the display function
	widget.display()
}

func (widget *Widget) Stop() {
	err := widget.client.Logout()
	err = widget.client.Close()
	if err != nil {
		log.Log(fmt.Sprintf("Error logging out: %s", err.Error()))
	}
}

/* -------------------- Unexported Functions -------------------- */
func (widget *Widget) login() {
	if widget.loggedIn {
		return
	}

	if err := widget.client.Login(widget.settings.username, widget.settings.password); err != nil {
		widget.clientError = err
		return
	}

	widget.loggedIn = true
}

func (widget *Widget) selectMailbox(mailboxName string) {
	mbox, err := widget.client.Select(mailboxName, false)
	if err != nil {
		widget.clientError = err
	}
	widget.currentMailbox = mbox
}

func (widget *Widget) listMailboxes() ([]*imap.MailboxInfo, error) {
	return listMailboxes(widget.client.List)
}

func (widget *Widget) listMessages() ([]*imap.Message, error) {
	return listMessages(widget.client.Fetch, widget.currentMailbox, widget.config)
}

func (widget *Widget) content() string {
	messages, err := widget.listMessages()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}

	content := ""
	for i := len(messages) - 1; i >= 0; i-- {
		content += fmt.Sprintf(
			"%s - %s <%s>\n",
			messages[i].Envelope.Subject,
			messages[i].Envelope.From[0].PersonalName,
			messages[i].Envelope.From[0].Address(),
		)
	}

	return content
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

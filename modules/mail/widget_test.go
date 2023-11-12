package mail

import (
	"errors"
	"fmt"
	"github.com/emersion/go-imap"
	"testing"
)

type FakeIMAPClient struct {
	users     map[string]string
	mailboxes map[string]*imap.MailboxStatus
	messages  []*imap.Message
	loggedIn  bool
	closed    bool
}

func (client *FakeIMAPClient) Close() error {
	client.closed = true
	return nil
}

func (client *FakeIMAPClient) Select(name string, readOnly bool) (*imap.MailboxStatus, error) {
	value, exists := client.mailboxes[name]
	if !exists {
		return nil, fmt.Errorf("mailbox %q does not exist", name)
	}
	return value, nil
}

func (client *FakeIMAPClient) Fetch(set *imap.SeqSet, items []imap.FetchItem, messages chan *imap.Message) error {
	defer close(messages)
	if client.messages == nil {
		return fmt.Errorf("no messages")
	}
	for _, message := range client.messages {
		messages <- message
	}
	return nil
}

func (client *FakeIMAPClient) Logout() error {
	client.loggedIn = false
	return nil
}

func (client *FakeIMAPClient) Login(username, password string) error {
	password, exists := client.users[username]
	if !exists || password != password {
		return errors.New("invalid username or password")
	}
	return nil
}

func (client *FakeIMAPClient) List(ref, name string, mailboxes chan *imap.MailboxInfo) error {
	defer close(mailboxes)
	for _, mailbox := range client.mailboxes {
		mailboxes <- &imap.MailboxInfo{
			Name: mailbox.Name,
		}
	}
	return nil
}

func TestSelectMailbox(t *testing.T) {
	t.Run("Existing mailbox", func(t *testing.T) {
		fakeClient := FakeIMAPClient{
			mailboxes: map[string]*imap.MailboxStatus{
				"INBOX": &imap.MailboxStatus{
					Name: "INBOX",
				},
			},
		}
		widget := Widget{
			client: &fakeClient,
		}

		widget.selectMailbox("INBOX")
		if widget.currentMailbox != fakeClient.mailboxes["INBOX"] {
			t.Errorf("Current mailbox is not the selected mailbox")
		}
	})
	t.Run("Non-existing mailbox", func(t *testing.T) {
		fakeClient := FakeIMAPClient{}
		widget := Widget{
			client: &fakeClient,
		}
		widget.selectMailbox("FOOBAR")
		if widget.currentMailbox != nil {
			t.Errorf("Current mailbox is not nil")
		}
		if widget.clientError.Error() != "mailbox \"FOOBAR\" does not exist" {
			t.Errorf("Client error %q is not %q", widget.clientError.Error(), "mailbox \"FOOBAR\" does not exist")
		}
	})
	t.Run("Log-in success", func(t *testing.T) {
		fakeClient := FakeIMAPClient{
			users: map[string]string{
				"test": "test",
			},
		}
		widget := Widget{
			client: &fakeClient,
			settings: &Settings{
				username: "test",
				password: "test",
			},
		}
		widget.login()
		if widget.loggedIn != true {
			t.Errorf("Not logged in")
		}
	})
	t.Run("Log-in failure", func(t *testing.T) {
		fakeClient := FakeIMAPClient{}
		widget := Widget{
			client: &fakeClient,
			settings: &Settings{
				username: "test",
				password: "test",
			},
		}
		widget.login()
		if widget.loggedIn != false {
			t.Errorf("Logged in when it shouldn't")
		}
		if widget.clientError.Error() != "invalid username or password" {
			t.Errorf("Client error is not %q", "invalid username or password")
		}
	})
	t.Run("No messages", func(t *testing.T) {
		widget := Widget{
			client:         &FakeIMAPClient{},
			config:         &Config{},
			currentMailbox: &imap.MailboxStatus{},
		}

		if _, err := widget.listMessages(); err == nil {
			t.Errorf("No error returned")
		}
	})
	t.Run("Returns messages", func(t *testing.T) {
		message := &imap.Message{
			SeqNum: 1,
		}

		widget := Widget{
			client: &FakeIMAPClient{
				messages: []*imap.Message{
					message,
				},
			},
			config: &Config{
				pageSize: 1,
			},
			currentMailbox: &imap.MailboxStatus{},
		}

		messages, err := widget.listMessages()

		if err != nil {
			t.Errorf("Error returned")
		}

		if len(messages) != 1 {
			t.Errorf("Incorrect number of messages returned")
		}

		if messages[0] != message {
			t.Errorf("Incorrect message returned")
		}
	})
	t.Run("Returns mailboxes", func(t *testing.T) {
		widget := Widget{
			client: &FakeIMAPClient{
				mailboxes: map[string]*imap.MailboxStatus{
					"INBOX": &imap.MailboxStatus{
						Name: "INBOX",
					},
				},
			},
		}

		mailboxes, err := widget.listMailboxes()

		if err != nil {
			t.Errorf("Error returned")
		}

		if len(mailboxes) != 1 {
			t.Errorf("Incorrect number of mailboxes returned")
		}

		if mailboxes[0].Name != "INBOX" {
			t.Errorf("Incorrect mailbox returned")
		}
	})
	t.Run("Stop logs-out and closes", func(t *testing.T) {
		fakeClient := FakeIMAPClient{
			loggedIn: true,
			closed:   false,
		}
		widget := Widget{
			client: &fakeClient,
		}
		widget.Stop()
		if fakeClient.loggedIn != false {
			t.Errorf("Not logged out")
		}
		if fakeClient.closed != true {
			t.Errorf("Not closed")
		}
	})
}

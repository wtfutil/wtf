package mail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"testing"
)

type getSequenceSetTestCase struct {
	config   *Config
	mailbox  *imap.MailboxStatus
	expected *imap.SeqSet
}

var getSequenceSetTestCases = []getSequenceSetTestCase{
	getSequenceSetTestCase{
		&Config{page: 0, pageSize: 10},
		&imap.MailboxStatus{Messages: 9},
		&imap.SeqSet{Set: []imap.Seq{imap.Seq{Start: uint32(1), Stop: uint32(9)}}},
	},
	getSequenceSetTestCase{
		&Config{page: 0, pageSize: 10},
		&imap.MailboxStatus{Messages: 11},
		&imap.SeqSet{Set: []imap.Seq{imap.Seq{Start: uint32(2), Stop: uint32(11)}}},
	},
	getSequenceSetTestCase{
		&Config{page: 0, pageSize: 10},
		&imap.MailboxStatus{Messages: 15},
		&imap.SeqSet{Set: []imap.Seq{imap.Seq{Start: uint32(6), Stop: uint32(15)}}},
	},
	getSequenceSetTestCase{
		&Config{page: 1, pageSize: 10},
		&imap.MailboxStatus{Messages: 15},
		&imap.SeqSet{Set: []imap.Seq{imap.Seq{Start: uint32(1), Stop: uint32(5)}}},
	},
	getSequenceSetTestCase{
		&Config{page: 1, pageSize: 10},
		&imap.MailboxStatus{Messages: 5},
		&imap.SeqSet{Set: []imap.Seq{}},
	},
}

func TestGetSequenceSet(t *testing.T) {
	for _, test := range getSequenceSetTestCases {
		if output := getSequenceSet(test.mailbox, test.config); output.String() != test.expected.String() {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func FakeFetchFunc(set *imap.SeqSet, items []imap.FetchItem, messages chan *imap.Message) error {
	defer close(messages)
	for _, s := range set.Set {
		for i := s.Start; i <= s.Stop; i++ {
			messages <- &imap.Message{
				SeqNum:   i,
				Envelope: &imap.Envelope{Subject: fmt.Sprintf("Subject %d", i)},
			}
		}
	}

	return nil
}

func TestListMessages(t *testing.T) {
	messages, err := listMessages(FakeFetchFunc, &imap.MailboxStatus{Messages: 15}, &Config{page: 0, pageSize: 10})

	if err != nil {
		t.Errorf("Error %q", err)
	}

	if len(messages) != 10 {
		t.Errorf("Expected 10 messages, got %d", len(messages))
	}

	if messages[0].SeqNum != 6 {
		t.Errorf("Expected first message to have SeqNum 6, got %d", messages[0].SeqNum)
	}

	if messages[9].SeqNum != 15 {
		t.Errorf("Expected last message to have SeqNum 15, got %d", messages[9].SeqNum)
	}
}

func FakeListFunc(ref, name string, mailboxes chan *imap.MailboxInfo) error {
	defer close(mailboxes)
	for i := 1; i <= 5; i++ {
		mailboxes <- &imap.MailboxInfo{
			Name: fmt.Sprintf("Mailbox %d", i),
		}
	}

	return nil
}

func TestListMailboxes(t *testing.T) {
	mailboxes, err := listMailboxes(FakeListFunc)

	if err != nil {
		t.Errorf("Error %q", err)
	}

	if len(mailboxes) != 5 {
		t.Errorf("Expected 5 messages, got %d", len(mailboxes))
	}

	if mailboxes[0].Name != "Mailbox 1" {
		t.Errorf("Expected first mailbox to have name 'Mailbox 1', got %q", mailboxes[0].Name)
	}

	if mailboxes[4].Name != "Mailbox 5" {
		t.Errorf("Expected first mailbox to have name 'Mailbox 5', got %q", mailboxes[0].Name)
	}
}

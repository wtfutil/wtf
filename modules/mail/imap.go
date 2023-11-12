package mail

import (
	"github.com/emersion/go-imap"
)

type Config struct {
	page     uint32
	pageSize uint32
}

type FetchFunc func(set *imap.SeqSet, items []imap.FetchItem, messages chan *imap.Message) error
type ListFunc func(ref, name string, mailboxes chan *imap.MailboxInfo) error

func getSequenceSet(mailbox *imap.MailboxStatus, config *Config) *imap.SeqSet {
	seqSet := new(imap.SeqSet)

	offset := config.page * config.pageSize

	if offset > mailbox.Messages {
		return seqSet
	}

	to := mailbox.Messages - offset
	from := uint32(1)

	if to > config.pageSize {
		from = to - config.pageSize + 1
	}

	seqSet.AddRange(from, to)

	return seqSet
}

func listMailboxes(listFunc ListFunc, numMailboxes uint32) ([]*imap.MailboxInfo, error) {
	mailboxes := make(chan *imap.MailboxInfo, numMailboxes)
	done := make(chan error, 1)
	defer close(done)

	go func() {
		done <- listFunc("", "*", mailboxes)
	}()

	if err := <-done; err != nil {
		return nil, err
	}

	mailboxesArray := make([]*imap.MailboxInfo, 0, len(mailboxes))

	for mailbox := range mailboxes {
		mailboxesArray = append(mailboxesArray, mailbox)
	}

	return mailboxesArray, nil
}

func listMessages(fetchFunc FetchFunc, mailbox *imap.MailboxStatus, config *Config) ([]*imap.Message, error) {
	seqSet := getSequenceSet(mailbox, config)

	messages := make(chan *imap.Message, config.pageSize)
	done := make(chan error, 1)
	defer close(done)

	go func() {
		done <- fetchFunc(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	if err := <-done; err != nil {
		return nil, err
	}

	messageArray := make([]*imap.Message, 0, len(messages))
	for message := range messages {
		messageArray = append(messageArray, message)
	}

	return messageArray, nil
}

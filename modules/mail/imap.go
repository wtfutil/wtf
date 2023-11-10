package mail

import (
	"github.com/emersion/go-imap"
)

type Config struct {
	page     uint32
	pageSize uint32
}

type FetchFunc func(set *imap.SeqSet, items []imap.FetchItem, messages chan *imap.Message) error

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

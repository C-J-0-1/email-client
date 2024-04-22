package main

import (
	"errors"
	"io"
	"log"
	"mime"
	"os"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
)

type Client struct {
	conn *imapclient.Client
}

func NewConnection() (Client, error) {
	options := &imapclient.Options{
		WordDecoder: &mime.WordDecoder{CharsetReader: charset.Reader},
	}

	conn, err := imapclient.DialTLS("imap.gmail.com:993", options)
	if err != nil {
		return Client{}, err
	}

	return Client{conn: conn}, nil
}

func (c *Client) Login() {
	if c.conn == nil {
		log.Fatal("no established connection")
	}

	loginCmd := c.conn.Login("", "")
	if err := loginCmd.Wait(); err != nil {
		log.Fatal("Failed to login: ", err)
	}
}

func (c *Client) Select() error {
	if c.conn == nil {
		return errors.New("no established connection")
	}

	selectCmd := c.conn.Select("INBOX", nil)
	if _, err := selectCmd.Wait(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Fetch() error {
	if c.conn == nil {
		return errors.New("no established connection")
	}

	seqset := imap.SeqSetNum(1201)
	fetchOptions := &imap.FetchOptions{
		Flags:       true,
		Envelope:    true,
		UID:         true,
		BodySection: []*imap.FetchItemBodySection{{}},
	}
	fetchCmd := c.conn.Fetch(seqset, fetchOptions)

	parseMail(fetchCmd.Next())

	return nil
}

func (c *Client) Logout() {
	if c.conn == nil {
		log.Fatal("no established connection")
	}

	logoutCmd := c.conn.Logout()
	if err := logoutCmd.Wait(); err != nil {
		log.Fatal("Failed to Logout: ", err)
	}
}

func parseMail(msg *imapclient.FetchMessageData) {
	if msg == nil {
		log.Println("Fetch no return")
		return
	}

	// Find body section
	var bodySection imapclient.FetchItemDataBodySection
	for {
		item := msg.Next()
		if item == nil {
			log.Println("Empty Mail")
			return
		}
		ok := false
		bodySection, ok = item.(imapclient.FetchItemDataBodySection)
		if ok {
			break
		}
	}

	// Read msg
	mr, err := mail.CreateReader(bodySection.Literal)
	if err != nil {
		log.Println("Failed to create")
		return
	}

	// headers
	h := mr.Header
	if date, err := h.Date(); err != nil {
		log.Println("Failed to parse date: ", err)
	} else {
		log.Println("Date: ", date)
	}
	if from, err := h.AddressList("From"); err != nil {
		log.Println("Failed to parse 'From': ", err)
	} else {
		log.Println("From: ", from)
	}
	if subject, err := h.Subject(); err != nil {
		log.Println("Failed to parse Subject: ", err)
	} else {
		log.Println("Subject: ", subject)
	}

	// Process the message
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			log.Println("EOF")
			break
		} else if err != nil {
			log.Println("Failed to read message part: ", err)
			break
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			b, _ := io.ReadAll(p.Body)
			log.Println("Body: ", string(b))
			f, err := os.OpenFile("mail.txt", os.O_APPEND, 0644)
			if err != nil {
				log.Fatalln("Could not open file: ", err)
			}
			if _, err := f.Write(b); err != nil {
				log.Fatalln("Failed to write ", err)
			}
			if err := f.Close(); err != nil {
				log.Fatalln("Failed to close file ", err)
			}
		case *mail.AttachmentHeader:
			filename, _ := h.Filename()
			log.Println("Attachment: ", filename)
		}
	}
}

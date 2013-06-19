package main

import (
	"fmt"
	"time"
)

type Message struct {
	Sender User
	Time   time.Time
	Text   string
}

func (m Message) String() string {
	return fmt.Sprintf("<%s> %s: %s", m.Time, m.Sender, m.Text)
}

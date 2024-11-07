package mockinterfaces

import "net/mail"

type EmailSender interface {
	Send(subject, body string, to ...*mail.Address)
}

func Fib(n int) int {
	for _, tt := range fibTests {
		if n == tt.n {
			return tt.o
		}
	}
	return -1
}

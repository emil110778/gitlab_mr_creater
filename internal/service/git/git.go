package git

import "regexp"

type Service struct {
	regExpTaskKeyWithDelim *regexp.Regexp
	regExpTaskKey          *regexp.Regexp
}

const (
	regExpTaskKey          = "[A-Z]+-[0-9]+"
	regExpTaskKeyWithDelim = "([-/_]|^)" + regExpTaskKey + "([-/_]|$)"
)

func New() (*Service, error) {
	return &Service{
		regExpTaskKeyWithDelim: regexp.MustCompile(regExpTaskKeyWithDelim),
		regExpTaskKey:          regexp.MustCompile(regExpTaskKey),
	}, nil
}

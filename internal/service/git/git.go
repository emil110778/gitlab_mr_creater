package git

import "regexp"

type Service struct {
	regExpTaskKeyWithDelim *regexp.Regexp
	regExpTaskKey          *regexp.Regexp
}

func New() (*Service, error) {
	regExpTaskKeyWithDelim, err := regexp.Compile("([-/_]|^)[A-Z]+-[0-9]+([-/_]|$)")
	if err != nil {
		return nil, err
	}
	regExpTaskKey, err := regexp.Compile("[A-Z]+-[0-9]+")
	if err != nil {
		return nil, err
	}
	return &Service{
		regExpTaskKeyWithDelim: regExpTaskKeyWithDelim,
		regExpTaskKey:          regExpTaskKey,
	}, nil
}

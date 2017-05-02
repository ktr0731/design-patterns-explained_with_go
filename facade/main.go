package main

import (
	"errors"
	"fmt"
	"os"
)

type member struct {
	id   int
	name string
}

type model []*member

func (m *model) getByLabMemNo(i int) (*member, error) {
	for _, member := range *m {
		if member.id == i {
			return member, nil
		}
	}
	return nil, errors.New("no such lab member number")
}

type db map[string]*model

func (s *db) getLabMems(key string) (*model, error) {
	if val, ok := (*s)[key]; ok {
		return val, nil
	}
	return nil, errors.New("no such lab")
}

var sampleDB = &db{
	"Future Gadget": &model{
		&member{1, "Okabe"},
		&member{2, "Shiina"},
		&member{3, "Hashida"},
		&member{4, "Makise"},
	},
}

func getMember(lab string, num int) (*member, error) {
	labMems, err := sampleDB.getLabMems(lab)
	if err != nil {
		return nil, err
	}
	comrade, err := labMems.getByLabMemNo(num)
	if err != nil {
		return nil, err
	}
	return comrade, nil
}

func main() {
	// Facade pattern: Compose some instructions to one for simplicity
	comrade, err := getMember("Future Gadget", 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Lab member ID: %d, %s\n", comrade.id, comrade.name)
}

package mend

import (
	"fmt"
	"testing"
)

type Person struct {
	Names      []Name
	Age        int
	Attributes map[string]bool
	AKAs       map[string][]string
}

type Name struct {
	FirstName string
	LastName  string
}

func TestMend(t *testing.T) {
	one := Person{
		Names:      []Name{{FirstName: "Bob", LastName: "Dole"}, Name{FirstName: "Robert", LastName: "Dole"}},
		Attributes: map[string]bool{"politician": false},
		AKAs:       map[string][]string{"aka": {"bdole", "blue pill"}},
	}
	two := Person{
		Names:      []Name{{FirstName: "Bob", LastName: "Dole"}, {FirstName: "Roberta", LastName: "DoleWhip"}},
		Age:        46,
		Attributes: map[string]bool{"politician": true, "alive": true},
		AKAs:       map[string][]string{"aka": {"bdoley"}},
	}

	Mend(&one, two)
	fmt.Println(one)
}

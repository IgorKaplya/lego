package walker

import (
	"fmt"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
	Nick    string
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name  string
		Input any
		Calls []string
	}{
		{
			Name:  "one string field",
			Input: struct{ Name string }{"Waka"},
			Calls: []string{"Waka"}},
		{
			Name: "two string fields",
			Input: struct {
				A string
				B string
			}{"Waka", "Waka"},
			Calls: []string{"Waka", "Waka"}},
		{
			Name: "nested fields",
			Input: Person{
				Name:    "Raka",
				Profile: Profile{69, "Maka"},
				Nick:    "Fo",
			},
			Calls: []string{"Raka", "Maka", "Fo"},
		},
		{
			Name: "pointer to things",
			Input: &Person{
				Name: "Biba",
				Profile: Profile{
					Age:  10,
					City: "Bobo",
				},
				Nick: "Boba",
			},
			Calls: []string{"Biba", "Bobo", "Boba"},
		},
		{
			Name:  "slices",
			Input: []Profile{{10, "Bob"}, {1, "Mob"}},
			Calls: []string{"Bob", "Mob"},
		},
		{
			Name:  "arrays",
			Input: [2]Profile{{10, "Jan"}, {1, "Jack"}},
			Calls: []string{"Jan", "Jack"},
		},
		{
			Name:  "maps",
			Input: map[string]string{"mini": "many", "moly": "moby"},
			Calls: []string{"many", "moby"},
		},
		{
			Name: "channels",
			Input: func() chan string {
				ch := make(chan string, 2)
				ch <- "zozo"
				ch <- "zaza"
				close(ch)
				return ch
			}(),
			Calls: []string{"zozo", "zaza"},
		},
		{
			Name: "func",
			Input: func() (string, int, string) {
				return "Fufu", 42, "Fifa"
			},
			Calls: []string{"Fufu", "Fifa"},
		},
	}

	for _, test := range cases {
		var got []string
		t.Run(test.Name, func(t *testing.T) {
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if err, equal := equalWithoutOrder(got, test.Calls); !equal {
				t.Error(err.Error())
			}
		})
	}
}

func equalWithoutOrder(s1, s2 []string) (error, bool) {
	if len(s1) != len(s2) {
		return fmt.Errorf("Length %d %v doesn't match %d for %v", len(s1), s1, len(s2), s2), false
	}

	slices.Sort(s1)

	for _, item2 := range s2 {
		if _, found := slices.BinarySearch(s1, item2); !found {
			return fmt.Errorf("Element %q of %v is not found in %v", item2, s2, s1), false
		}
	}

	return nil, true
}

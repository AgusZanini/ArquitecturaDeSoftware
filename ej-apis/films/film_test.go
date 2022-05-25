package films

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Id                       string
	Title                    string
	Original_title           string
	Original_title_romanised string
	Description              string
	Director                 string
	Producer                 string
	Release_date             string
	Running_time             string
	Rt_score                 string
	People                   []string
	Species                  []string
	Locations                []string
	Vehicles                 []string
	Url                      string
}

func TestParsefilms(t *testing.T) {
	films, err := Getfilms("prueba")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	film := films[2]
	ExpectedDescription := "In the latter part of World War II, a boy and his sister, orphaned when their mother is killed in the firebombing of Tokyo, are left to survive on their own in what remains of civilian life in Japan. The plot follows this boy and his sister as they do their best to survive in the Japanese countryside, battling hunger, prejudice, and pride in their own quiet, personal battle."

	if ExpectedDescription != film.Description {
		t.Errorf("film.Description == %q, want %q", film.Description, ExpectedDescription)
	}
}

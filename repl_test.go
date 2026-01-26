package main

import (
	"bytes"
	"encoding/json"
	"pokedex/pokeapi"
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "    awkward, string1 and the rest ",
			expected: []string{
				"awkward,",
				"string1",
				"and",
				"the",
				"rest",
			},
		},
		{
			input: "beedrill 52",
			expected: []string{
				"beedrill",
				"52",
			},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) == 0 && len(actual) == 0 {
			continue
		}

		if len(actual) != len(c.expected) {
			t.Errorf("expected word count was %d and function returned %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("expected word was %s and function returned %s", expectedWord, word)
			}
		}
	}

}

func TestCommandHelp(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "Diaplays a help Message",
		},
	}

	for range cases {
		error := commandHelp(&config{})
		if error != nil {
			t.Errorf("Failed commmand %v", error)
		}
	}
}

func TestGetLocationAreas(t *testing.T) {
	locationAreaAPIResponse, err := pokeapi.GetURl("https://pokeapi.co/api/v2/location-area/")
	locationAreaData, err := pokeapi.DecodeData[LocationAreaResponse](bytes.NewBuffer(locationAreaAPIResponse))
	if err != nil {
		t.Errorf("Failed to get location areas: %v", err)
	}
	controlDataJSON := []byte(`{"count":1089,"next":"https://pokeapi.co/api/v2/location-area/?offset=20&limit=20","previous":null,
	"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},
	{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},
	{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},
	{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},
	{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},
	{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},
	{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},
	{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},
	{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},
	{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},
	{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},
	{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},
	{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},
	{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},
	{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},
	{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},
	{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},
	{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},
	{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},
	{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`)
	var controlData LocationAreaResponse
	err = json.Unmarshal(controlDataJSON, &controlData)
	if err != nil {
		t.Errorf("Failed to unmarshal control data: %v", err)
	}

	if !reflect.DeepEqual(controlData, locationAreaData) {
		t.Errorf("API response does not match control data")
	}
}

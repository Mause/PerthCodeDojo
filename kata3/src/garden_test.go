package main

import (
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "testing"
)

func ReadJson(t *testing.T, filename string) Garden {
	raw_json_result, err := ioutil.ReadFile("japanese_garden.json.out.json")
	assert.Nil(t, err)
	var json_result Garden
	err = json.Unmarshal(raw_json_result, &json_result)
	assert.Nil(t, err)

	return json_result
}

func TestJapaneseGarden(t *testing.T) {
	GardenFunction("japanese_garden.json")

	end_result := ReadJson(t, "japanese_garden.json.out.json")
	original := ReadJson(t, "japanese_garden.json")

	assert.Equal(
		t,
		end_result,
		original,
	)
}

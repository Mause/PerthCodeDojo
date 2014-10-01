package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func readJson(t *testing.T, filename string) Garden {
	raw_json_result, err := ioutil.ReadFile("japanese_garden.json.out.json")
	assert.Nil(t, err)
	var json_result Garden
	err = json.Unmarshal(raw_json_result, &json_result)
	assert.Nil(t, err)

	return json_result
}

func testGarden(t *testing.T, filename string) {
	GardenFunction(filename)

	end_result := readJson(t, filename+".out.json")
	original := readJson(t, filename)

	assert.Equal(
		t,
		end_result,
		original,
	)
}

func TestJapaneseGarden(t *testing.T) {
    testGarden(t, "japanese_garden.json")
}

func TestNormalGarden(t *testing.T) {
    testGarden(t, "garden.json")
}

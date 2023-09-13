package main

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setup() {
	inputDescription = ""
	inputFooter = ""
	templatePath = "./README.md.tmpl"
}

func TestOneTil(t *testing.T) {
	setup()
	repoPath = "./test_data/1_til"
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

func TestZeroTil(t *testing.T) {
	setup()
	repoPath = "./test_data/zero_til"
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

func TestManyTil(t *testing.T) {
	setup()
	repoPath = "./test_data/many_til"
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

func TestManyTilToLowercase(t *testing.T) {
	setup()
	repoPath = "./test_data/many_til_to_lowercase"
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

func TestOneTilInputs(t *testing.T) {
	setup()
	repoPath = "./test_data/many_with_inputs"
	inputDescription = "This is a placeholder description used for testing."
	inputFooter = "here is where the markdown footer links would go"
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

func TestOneTilInputsAndMostRecent(t *testing.T) {
	setup()
	repoPath = "./test_data/many_with_inputs_and_most_recent"
	inputDescription = "This is a placeholder description used for testing."
	inputFooter = "here is where the markdown footer links would go"
	inputListMostRecent = "3"
	inputDateFormat = time.RFC822
	main()
	expected, err := ioutil.ReadFile(repoPath + "/README.md.expected")
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadFile(repoPath + "/README.md")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(expected), string(actual))
}

package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneTil(t *testing.T) {
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

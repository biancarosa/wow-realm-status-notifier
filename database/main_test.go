package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServerFromTextWithBadInput(t *testing.T) {
	path := "test.db"
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	Init("test.db")

	f, err := os.Stat("test.db")
	assert.Nil(t, err)
	assert.NotNil(t, f)
}

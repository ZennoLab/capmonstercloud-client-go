package client

import (
	"os"
	"testing"
)

func TestGetBalance(t *testing.T) {
	client := New(os.Getenv(testingKeyEnvVarName))
	balance, err := client.GetBalance()

	switch {
	case err != nil:
		t.Error(err)
	case balance < 0:
		t.Error("negative balance")
	}
}

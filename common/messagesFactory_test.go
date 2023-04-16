package common_test

import (
	"davidPardoC/rest/common"
	"testing"
)

func TestSuccesMessage(t *testing.T) {
	message := common.CreateSuccesCreatedMessage(common.UserResource)

	if message.Message != "User created successfully" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", message.Message, "User created successfully")
	}
}

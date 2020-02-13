package helpers

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	guid := GenerateUUID()
	if len(guid) != 36 {
		t.Errorf("expected: %d\n got: %d\n", len(guid), 36)
	}
}

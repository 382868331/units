package units

import "testing"

func TestTaskUnits016LeadingIntegerRejectsOverflow(t *testing.T) {
	if _, _, err := leadingInt("99999999999999999999"); err == nil {
		t.Fatal("oversized integer prefix was accepted")
	}
}

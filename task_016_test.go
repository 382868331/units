package units

import "testing"

func TestTaskUnits016LeadingIntegerRejectsOverflow(t *testing.T) {
	if _, _, err := leadingInt("99999999999999999999"); err == nil {
		t.Fatal("oversized integer prefix was accepted")
	}
}

func TestTaskUnits016LeadingIntegerAcceptsMaximum(t *testing.T) {
	got, rem, err := leadingInt("922337203685477579B")
	if err != nil || rem != "B" || got != int64(922337203685477579) {
		t.Fatalf("got=%d rem=%q err=%v", got, rem, err)
	}
}

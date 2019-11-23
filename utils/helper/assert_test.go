package helper

import "testing"

func TestIsContain(t *testing.T) {
	week := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}

	today := "Tuesday"

	got := IsContain(today, week, false)
	want := true
	if got != want {
		t.Errorf("got is %t, want is %t", got, want)
	}
}

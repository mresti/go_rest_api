package models

import (
	"app/testdata"
	"testing"
)

func TestVisitAPI_marshall(t *testing.T) {
	d := &VisitAPI{
		Visits: 1,
	}
	want := `{"visits":1}`
	testdata.TestJSONMarshal(t, d, want)
}

func TestMessageAPI_marshall(t *testing.T) {
	d := &MessageAPI{
		Message: "Hello",
	}
	want := `{"message":"Hello"}`
	testdata.TestJSONMarshal(t, d, want)
}

func TestErrorAPI_marshall(t *testing.T) {
	d := &ErrorAPI{
		Error: "My error",
	}
	want := `{"error":"My error"}`
	testdata.TestJSONMarshal(t, d, want)
}

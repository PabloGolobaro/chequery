package entity

import (
	"encoding/json"
	"testing"
)

func TestOrderDetails_Details(t *testing.T) {
	type fields struct {
		PointID int
		Order   string
	}

	js := struct {
		Foo int    `json:"foo,omitempty"`
		Bar bool   `json:"bar,omitempty"`
		Baz string `json:"baz,omitempty"`
	}{
		Foo: 1,
		Bar: true,
		Baz: "golo",
	}

	bytes, err := json.Marshal(&js)
	if err != nil {
		t.Fatal(err)
	}

	f := fields{
		PointID: 1,
		Order:   string(bytes),
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "simple", fields: f, want: "foo - 1\nbar - true\nbaz - golo\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderDetails{
				PointID: tt.fields.PointID,
				Order:   tt.fields.Order,
			}
			if got := o.Details(); got != tt.want {
				t.Errorf("Details() = %v, want %v", got, tt.want)
			}
		})
	}
}

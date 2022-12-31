package main

import "testing"

var tests = []struct {
	name     string
	enemy    Enemy
	expected float64
	isErr    bool
}{
	{"valid-data", validEnemy, 17, false},
	{"invalid-data", invalidEnemy, 0, true},
}

func TestCalculateDamage(t *testing.T) {
	for _, tt := range tests {
		got, err := tt.enemy.calculateDamage()
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

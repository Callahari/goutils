package goutils

import (
	"testing"
	"time"
)

func TestRandomDateBetween(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		isWrong bool
	}{
		{name: "OK", args: args{startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), endDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}, want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), isWrong: false},
		{name: "Wrong", args: args{startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), endDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}, want: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), isWrong: true},
		{name: "SameDay", args: args{startDate: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC), endDate: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)}, want: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC), isWrong: false},
		{name: "OneDayApart", args: args{startDate: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC), endDate: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)}, want: time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC), isWrong: true},
		{name: "CheckLeapYear", args: args{startDate: time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC), endDate: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)}, want: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), isWrong: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomDateBetween(tt.args.startDate, tt.args.endDate)
			if tt.isWrong {
				if got == tt.want {
					t.Errorf("RandomDateBetween() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			} else {
				if got != tt.want {
					t.Errorf("RandomDateBetween() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			}
		})
	}
}
func TestRandomBasedOnProbability(t *testing.T) {
	tests := []struct {
		name        string
		expected    bool
		inputString string
	}{
		{name: "OK", expected: true, inputString: "1:1"},
		{name: "False", expected: false, inputString: "1:20000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomBasedOnProbability(tt.inputString)
			if err != nil {
				t.Errorf("RandomBasedOnProbability() error = %v", err)
				t.FailNow()
			}
			if got != tt.expected {
				t.Errorf("RandomBasedOnProbability() = %v, want %v", got, tt.expected)
				t.FailNow()
			}
		})
	}
}
func TestRandomFloat32Between(t *testing.T) {
	type args struct {
		min float32
		max float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		isWrong bool
	}{
		{name: "OK", args: args{min: 3.666, max: 3.666}, want: 3.666, isWrong: false},
		{name: "Wrong", args: args{min: 7.666, max: 7.666}, want: 6.666, isWrong: true},
	}
	for _, tt := range tests {
		got := RandomFloat32Between(tt.args.min, tt.args.max)
		if tt.isWrong {
			if got == tt.want {
				t.Errorf("RandomFloat32Between() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		} else {
			if got != tt.want {
				t.Errorf("RandomFloat32Between() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		}
	}
}

func TestRandomFloat64Between(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		isWrong bool
	}{
		{name: "OK", args: args{min: 3.666, max: 3.666}, want: 3.666, isWrong: false},
		{name: "Wrong", args: args{min: 7.666, max: 7.666}, want: 6.666, isWrong: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomFloat64Between(tt.args.min, tt.args.max)
			if tt.isWrong {
				if got == tt.want {
					t.Errorf("RandomFloat64Between() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			} else {
				if got != tt.want {
					t.Errorf("RandomFloat64Between() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			}
		})
	}
}

func TestRandomIntBetween(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		isWrong bool
	}{
		{name: "OK", args: args{min: 3, max: 3}, want: 3, isWrong: false},
		{name: "Wrong", args: args{min: 7, max: 7}, want: 6, isWrong: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomIntBetween(tt.args.min, tt.args.max)
			if tt.isWrong {
				if got == tt.want {
					t.Errorf("RandomIntBetween() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			} else {
				if got != tt.want {
					t.Errorf("RandomIntBetween() = %v, want %v", got, tt.want)
					t.FailNow()
				}
			}
		})
	}
}

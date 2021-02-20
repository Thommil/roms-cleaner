package core

import "testing"

func TestGetRegion(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    Region
		wantErr bool
	}{
		{"simple key", args{"USA"}, USA, false},
		{"mixed case key", args{"eURope"}, EUROPE, false},
		{"rom ID", args{"slus"}, USA, false},
		{"rom meta", args{"[E]"}, EUROPE, false},
		{"invalid region & fallback", args{"japon"}, EUROPE, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegion(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

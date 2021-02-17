package core

import (
	"reflect"
	"testing"
)

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

func TestGetSystem(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"simple key", args{"n64"}, System{"n64"}, false},
		{"mixed case key", args{"aMIga"}, System{"amiga"}, false},
		{"bad key", args{"foo"}, System{}, true},

		{"path OK", args{"./test/roms/arcade"}, System{"arcade"}, false},
		{"path KO", args{"./test/roms"}, System{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSystem(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

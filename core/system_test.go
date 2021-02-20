package core

import (
	"reflect"
	"testing"
)

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
		{"simple key", args{"n64"}, systems["n64"], false},
		{"mixed case key", args{"aMIga"}, systems["amiga"], false},
		{"bad key", args{"foo"}, System{}, true},

		{"path OK", args{"./test/roms/arcade"}, systems["arcade"], false},
		{"custom OK", args{"./test/roms/fba"}, systems["arcade"], false},
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

package scanner

import (
	"testing"

	"github.com/thommil/roms-cleaner/core"
)

func TestScan(t *testing.T) {
	type args struct {
		options core.Options
		games   []core.Game
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"get arcade scanner", args{core.Options{System: core.System{ID: "arcade"}}, make([]core.Game, 0)}, false},
		{"get unknown scanner", args{core.Options{System: core.System{ID: "foo"}}, make([]core.Game, 0)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Scan(tt.args.options, tt.args.games); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

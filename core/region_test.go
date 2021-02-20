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
		want    []Region
		wantErr bool
	}{
		{"simple key", args{"USA"}, []Region{USA}, false},
		{"mixed case key", args{"eURope"}, []Region{EUROPE}, false},
		{"rom meta", args{"[E]"}, []Region{EUROPE}, false},
		{"invalid region & fallback", args{"japon"}, nil, true},
		{"cyber-speed-u-slus-00116-", args{"cyber-speed-u-slus-00116-"}, []Region{USA}, false},
		{"Akumajou Dracula (T-Eng).fds", args{"Akumajou Dracula (T-Eng).fds"}, []Region{EUROPE}, false},
		{"Akuu Senki Raijin (Japan).fds", args{"Akuu Senki Raijin (Japan).fds"}, []Region{JAPAN}, false},
		{"Dandy - Zeuon no Fukkatsu (T-Eng).fds", args{"Dandy - Zeuon no Fukkatsu (T-Eng).fds"}, []Region{EUROPE}, false},
		{"ax battler - a legend of golden axe (usa, europe) (v2.4).bin", args{"ax battler - a legend of golden axe (usa, europe) (v2.4).bin"}, []Region{USA, EUROPE}, false},
		{"batman forever (world).bin", args{"batman forever (world).bin"}, []Region{EUROPE}, false},
		{"Battletoads (Euro, Jpn)", args{"Battletoads (Euro, Jpn)"}, []Region{EUROPE, JAPAN}, false},
		{"Advanced Daisenryaku - Deutsch Dengeki Sakusen (Jpn, Rev. A)", args{"Advanced Daisenryaku - Deutsch Dengeki Sakusen (Jpn, Rev. A)"}, []Region{JAPAN}, false},
		{"Arkanoid (Euro?)", args{"Arkanoid (Euro?)"}, nil, true},
		{"123 Sesame Street - Ready, Set, Grover! - With Elmo - The Videogame (USA) (En,Es)", args{"123 Sesame Street - Ready, Set, Grover! - With Elmo - The Videogame (USA) (En,Es)"}, []Region{USA}, false},
		{"SNK Neo-Geo Pocket BIOS (1998)(SNK)(en-ja).bin", args{"SNK Neo-Geo Pocket BIOS (1998)(SNK)(en-ja).bin"}, []Region{JAPAN}, false},
		{"city hunter (english translation).pce", args{"city hunter (english translation).pce"}, []Region{EUROPE}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegion(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

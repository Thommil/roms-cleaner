package core

import (
	"os"
	"testing"
)

func TestDAT_FromMemory(t *testing.T) {
	var dat DAT
	var err error

	t.Run("load bad system", func(t *testing.T) {
		err = dat.FromMemory("foo")

		if err == nil {
			t.Errorf("FromMemory() error should not be nil")
			return
		}
	})

	t.Run("load system", func(t *testing.T) {
		err = dat.FromMemory("arcade")

		if err != nil {
			t.Errorf("FromMemory() error = %v", err)
			return
		}

		if len(dat.Games) == 0 {
			t.Errorf("FromMemory() empty games list")
			return
		}

	})
}

func TestDAT_FromXML(t *testing.T) {
	var dat DAT
	var data []byte
	var err error

	t.Run("load bad XML", func(t *testing.T) {
		data = []byte("corrupted data")

		err = dat.FromXML(data)

		if err == nil {
			t.Errorf("FromXML() error should not be nil")
			return
		}
	})

	t.Run("load XML", func(t *testing.T) {
		data, err = os.ReadFile("../data/dats/arcade.dat")

		if err != nil {
			t.Errorf("FromXML() read error = %v", err)
			return
		}

		err = dat.FromXML(data)

		if err != nil {
			t.Errorf("FromXML() error = %v", err)
			return
		}

		if len(dat.Games) == 0 {
			t.Errorf("FromXML() empty games list")
			return
		}

	})
}

func TestDAT_Serialize(t *testing.T) {
	var dat DAT
	var data []byte
	var err error

	t.Run("init", func(t *testing.T) {
		data, err = os.ReadFile("../data/dats/arcade.dat")

		if err != nil {
			t.Errorf("FromXML() error = %v", err)
			return
		}

		err = dat.FromXML(data)

		if err != nil {
			t.Errorf("FromXML() error = %v", err)
			return
		}

		if len(dat.Games) == 0 {
			t.Errorf("FromXML() empty games list")
			return
		}

	})

	t.Run("check content", func(t *testing.T) {
		data, err = dat.Serialize()

		if err != nil {
			t.Errorf("Serialize() error = %v", err)
			return
		}

		if len(data) == 0 {
			t.Errorf("Serialize() out data is empty")
			return
		}

	})
}

func TestDAT_Deserialize(t *testing.T) {
	var dat DAT
	var data []byte
	var err error

	t.Run("nil content", func(t *testing.T) {
		err = dat.Deserialize(nil)

		if err == nil {
			t.Errorf("Deserialize() error should not be nil")
			return
		}
	})

	t.Run("corrupted content", func(t *testing.T) {
		err = dat.Deserialize([]byte("corrupted content"))

		if err == nil {
			t.Errorf("Deserialize() error should not be nil")
			return
		}
	})

	t.Run("init from XML", func(t *testing.T) {
		data, err = os.ReadFile("../data/dats/arcade.dat")

		if err != nil {
			t.Errorf("FromXML() init error = %v", err)
			return
		}

		err = dat.FromXML(data)

		if err != nil {
			t.Errorf("FromXML() error = %v", err)
			return
		}

		if len(dat.Games) == 0 {
			t.Errorf("FromXML() empty games list")
			return
		}

	})

	t.Run("init Serialize", func(t *testing.T) {
		data, err = dat.Serialize()

		if err != nil {
			t.Errorf("Writ() error = %v", err)
			return
		}

		if err != nil {
			t.Errorf("Serialize() error = %v", err)
			return
		}
	})

	t.Run("check content", func(t *testing.T) {
		if err != nil {
			t.Errorf("READ() error = %v", err)
			return
		}
		err = dat.Deserialize(data)

		if err != nil {
			t.Errorf("Deserialize() error = %v", err)
			return
		}

		if len(dat.Games) == 0 {

			t.Errorf("Deserialize() empty games list")
			return
		}

	})
}

// func BenchmarkGetGame(b *testing.B) {
// 	data, err := os.ReadFile("../test/dats/arcade.dat")

// 	if err != nil {
// 		b.Errorf("error = %v", err)
// 		return
// 	}

// 	dat, err := ParseDAT(data)

// 	if err != nil {
// 		b.Errorf("ParseDAT() error = %v", err)
// 		return
// 	}
// 	for i := 0; i < b.N; i++ {
// 		search := "mrdobl"

// 		for _, game := range dat.Games {
// 			if game.Name == search {
// 				break
// 			}
// 		}
// 	}
// }

package core

import (
	"testing"
)

var testData = []struct {
	Name       string
	FieldStr   string
	FieldInt   int
	FieldArray []string
	FieldOmit  string
	FieldDoc   struct {
		SubName string
		SubOmit string
	}
	DocOmit struct {
		SubName string
	}
}{
	{"one", "1", 1, []string{"first", "premier"}, "omitOne",
		struct {
			SubName string
			SubOmit string
		}{"subOne", "subOmitOne"},
		struct {
			SubName string
		}{"docOmitOne"},
	},
	{"two", "2", 2, []string{"second", "deuxième"}, "omitTwo",
		struct {
			SubName string
			SubOmit string
		}{"subTwo", "subOmitTwo"},
		struct {
			SubName string
		}{"docOmitTwo"},
	},
	{"three", "-", 3, []string{"-", "-"}, "-",
		struct {
			SubName string
			SubOmit string
		}{"-", "-"},
		struct {
			SubName string
		}{"-"},
	},
}

func TestCreateIndex(t *testing.T) {
	var index Index
	var err error

	t.Run("index create, nil options", func(t *testing.T) {
		index, err = CreateIndex(nil)

		if err != nil {
			t.Errorf("CreateIndex() error should not be nil")
			return
		}

		index.Close()
	})

	t.Run("index create", func(t *testing.T) {
		index, err = CreateIndex([]string{"FieldOmit", "DocOmit", "FieldDoc.SubOmit"})

		if err != nil {
			t.Errorf("CreateIndex() error = %v", err)
			return
		}

		if index == nil {
			t.Errorf("CreateIndex() index should not be nil")
			return
		}

		index.Close()
	})

}

func Test_bleveIndex_Add(t *testing.T) {
	var index Index
	var err error

	t.Run("init", func(t *testing.T) {
		index, err = CreateIndex([]string{"FieldOmit", "DocOmit", "FieldDoc.SubOmit"})

		if err != nil {
			t.Errorf("CreateIndex() error = %v", err)
			return
		}

		if index == nil {
			t.Fatalf("CreateIndex() index should not be nil")
			return
		}
	})

	t.Run("add simple", func(t *testing.T) {
		for _, data := range testData {
			err = index.Add(data.Name, data)

			if err != nil {
				t.Errorf("Add() error = %v", err)
				return
			}
		}
	})

	t.Run("close index", func(t *testing.T) {
		err = index.Close()

		if err != nil {
			t.Errorf("Close() error = %v", err)
			return
		}
	})
}

func Test_bleveIndex_Search(t *testing.T) {
	var index Index
	var err error

	t.Run("init", func(t *testing.T) {
		index, err = CreateIndex([]string{"FieldOmit", "DocOmit", "FieldDoc.SubOmit"})

		if err != nil {
			t.Errorf("CreateIndex() error = %v", err)
			return
		}

		if index == nil {
			t.Fatalf("CreateIndex() index should not be nil")
			return
		}

		for _, data := range testData {
			err = index.Add(data.Name, data)

			if err != nil {
				t.Errorf("Add() error = %v", err)
				return
			}
		}
	})

	t.Run("search by Name", func(t *testing.T) {
		result, err := index.Search("two")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 1 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 1, len(result))
		}
	})

	t.Run("search by FieldStr", func(t *testing.T) {
		result, err := index.Search("2")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 1 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 1, len(result))
		}
	})

	t.Run("search by FieldInt", func(t *testing.T) {
		result, err := index.Search("3")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 1 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 1, len(result))
		}
	})

	t.Run("search by FieldArray", func(t *testing.T) {
		result, err := index.Search("deuxième")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 1 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 1, len(result))
		}
	})

	t.Run("search by FieldDoc", func(t *testing.T) {
		result, err := index.Search("subOne")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 1 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 1, len(result))
		}
	})

	t.Run("search by DocOmit", func(t *testing.T) {
		result, err := index.Search("docOmitOne")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 0 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 0, len(result))
		}
	})

	t.Run("search by FieldOmit", func(t *testing.T) {
		result, err := index.Search("omitTwo")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 0 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 0, len(result))
		}
	})

	t.Run("search by SubOmit", func(t *testing.T) {
		result, err := index.Search("subOmitOne")

		if err != nil {
			t.Errorf("Search() error = %v", err)
			return
		}

		if len(result) != 0 {
			t.Log(result)
			t.Errorf("Search() wanted %d, got %d", 0, len(result))
		}
	})

	index.Close()
}

/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polylang_test

import (
	"os"
	"testing"

	"github.com/v1def/go-polybase/polylang"

	"github.com/alecthomas/participle/v2"
)

func TestCollection(t *testing.T) {
	parser := participle.MustBuild[polylang.Collection]()

	f, err := os.Open("./fixtures/collection.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	if got.Name != "Polybase" {
		t.Fatal("error: collection name does not match")
	}
}

func TestField(t *testing.T) {
	parser := participle.MustBuild[polylang.Field]()

	f, err := os.Open("./fixtures/field.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Name != "polybase":
		t.Fatal("error: field name does not match")
	case got.Optional != true:
		t.Fatal("error: field optional does not match")
	case got.Type != "string":
		t.Fatal("error: field type does not match")
	}
}

func TestFunction(t *testing.T) {
	parser := participle.MustBuild[polylang.Function]()

	f, err := os.Open("./fixtures/function.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	if got.Name != "constructor" {
		t.Fatal("error: function name does not match")
	}
}

func TestParameter(t *testing.T) {
	parser := participle.MustBuild[polylang.Parameter]()

	f, err := os.Open("./fixtures/parameter.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Name != "test":
		t.Fatal("error: parameter name does not match")
	case got.Optional != true:
		t.Fatal("error: parameter required does not match")
	case got.Type != "string":
		t.Fatal("error: parameter type does not match")
	}
}

func TestIndex(t *testing.T) {
	parser := participle.MustBuild[polylang.Index]()

	f, err := os.Open("./fixtures/index.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	if got.Unique != false {
		t.Fatal("error: unique index does not match")
	}
}

func TestUnique(t *testing.T) {
	parser := participle.MustBuild[polylang.Index]()

	f, err := os.Open("./fixtures/unique.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	if got.Unique != true {
		t.Fatal("error: unique index does not match")
	}
}

func TestIndexField(t *testing.T) {
	parser := participle.MustBuild[polylang.IndexField]()

	f, err := os.Open("./fixtures/index_field.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Name != "name":
		t.Fatal("error: index field name does not match")
	case got.Order != "desc":
		t.Fatal("error: index field order does not match")
	}
}

func TestFull(t *testing.T) {
	parser := participle.MustBuild[polylang.Collection]()

	f, err := os.Open("./fixtures/full.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Name != "Full":
		t.Fatal("error: collection name does not match")
	case got.Items[1].Field.Name != "name":
		t.Fatal("error: field name does not match")
	case got.Items[2].Index.Unique != false:
		t.Fatal("error: index unique does not match")
	case got.Items[4].Function.Name != "update":
		t.Fatal("error: function name does not match")
	}
}

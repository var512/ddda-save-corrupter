package parser

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/beevik/etree"

	"github.com/var512/ddda-save-corrupter/internal/nodemap"
)

var (
	Doc *etree.Document
)

func TestMain(m *testing.M) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile("../testdata/base.sav.xml")
	if err != nil {
		panic(err)
	}
	Doc = doc

	os.Exit(m.Run())
}

func TestParseAttributes(t *testing.T) {
	var o map[string]string
	var err error

	mproot := nodemap.MainPawnParent() + "/" + nodemap.Nodemap()["mEdit"] + "/"
	fproot := nodemap.FirstPawnParent() + "/" + nodemap.Nodemap()["mEdit"] + "/"
	sproot := nodemap.SecondPawnParent() + "/" + nodemap.Nodemap()["mEdit"] + "/"

	// Main pawn mGender.
	o, err = parse(&Doc.Element, mproot+`u8[@name='mGender']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if o["value"] != "1" {
		t.Errorf("mGender: want 1, got: %#v", o)
	}

	// First pawn mFaceBase.
	o, err = parse(&Doc.Element, fproot+`u8[@name='mFaceBase']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if o["value"] != "73" {
		t.Errorf("mFaceBase: want 73, got: %#v", o)
	}

	// Second pawn mColorMakeup.
	o, err = parse(&Doc.Element, sproot+`u8[@name='mColorMakeup']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if o["value"] != "13" {
		t.Errorf("mColorMakeup: want 13, got: %#v", o)
	}

	// First pawn u8 array mNameStr, u8 fields, u8 value to string.
	var s strings.Builder
	nameStr, err := GetU8Array(&Doc.Element, fproot+`array[@name='(u8*)mNameStr']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	for _, child := range nameStr.Children {
		if child.Value < 1 {
			break
		}
		s.WriteString(fmt.Sprintf("%c", child.Value))
	}

	if s.String() != "Wilson" {
		t.Errorf("mNameStr: want Wilson, got: %#v", s.String())
	}
}

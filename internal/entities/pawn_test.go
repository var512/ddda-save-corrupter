package entities

import (
	"os"
	"testing"

	"github.com/beevik/etree"

	"github.com/var512/ddda-save-corrupter/internal/nodemap"
	"github.com/var512/ddda-save-corrupter/internal/parser"
)

var (
	Doc    *etree.Document
	mproot = nodemap.MainPawnParent() + "/" + nodemap.Nodemap()["mEdit"] + "/"
	sproot = nodemap.SecondPawnParent() + "/" + nodemap.Nodemap()["mEdit"] + "/"
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

// Nodemap.MainPawnAppearanceOverrideParent() + mEditPawn.
func TestAppearanceOverrideFromAppearanceOverrideParent(t *testing.T) {
	aoroot := Doc.Element.FindElement(nodemap.MainPawnAppearanceOverrideParent())
	appOver, err := NewAppearanceOverrideFromRootNode(aoroot, nodemap.Nodemap()["mEditPawn"])

	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	mHairNo, err := parser.GetU8(appOver.Element, `u8[@name='mHairNo']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if int(mHairNo.Value) != 99 {
		t.Errorf("mHairNo: want 99, got: %v", mHairNo.Value)
	}

	nameStr, err := parser.GetU8Array(appOver.Element, `array[@name='(u8*)mNameStr']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if len(nameStr.Children) != 25 {
		t.Fatalf("nameStr.Children size: want 25, got: %#v", len(nameStr.Children))
	}
	if nameStr.Children[0].Value != 11 {
		t.Errorf("nameStr.Children: want 11, got: %v", nameStr.Children[0].Value)
	}
	if nameStr.Children[4].Value != 55 {
		t.Errorf("nameStr.Children: want 55, got: %v", nameStr.Children[4].Value)
	}
	if nameStr.Children[7].Value != 88 {
		t.Errorf("nameStr.Children: want 88, got: %v", nameStr.Children[7].Value)
	}
	if nameStr.Children[9].Value != 0 {
		t.Errorf("nameStr.Children: want 0, got: %v", nameStr.Children[9].Value)
	}
}

// Nodemap.MainPawnParent() + mEdit.
func TestAppearanceOverrideFromParent(t *testing.T) {
	aoroot := Doc.Element.FindElement(nodemap.MainPawnParent())
	appOver, err := NewAppearanceOverrideFromRootNode(aoroot, nodemap.Nodemap()["mEdit"])

	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	mHairNo, err := parser.GetU8(appOver.Element, `u8[@name='mHairNo']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if int(mHairNo.Value) != 0 {
		t.Errorf("want 0, got: %#v", int(mHairNo.Value))
	}

	nameStr, err := parser.GetU8Array(appOver.Element, `array[@name='(u8*)mNameStr']`)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if len(nameStr.Children) != 25 {
		t.Fatalf("nameStr.Children size: want 25, got: %#v", len(nameStr.Children))
	}
	if nameStr.Children[0].Value != 77 {
		t.Errorf("nameStr.Children: want 77, got: %v", nameStr.Children[0].Value)
	}
	if nameStr.Children[4].Value != 112 {
		t.Errorf("nameStr.Children: want 112, got: %v", nameStr.Children[4].Value)
	}
	if nameStr.Children[7].Value != 110 {
		t.Errorf("nameStr.Children: want 110, got: %v", nameStr.Children[7].Value)
	}
	if nameStr.Children[8].Value != 0 {
		t.Errorf("nameStr.Children: want 0, got: %v", nameStr.Children[8].Value)
	}
}

func TestNewPawnFromDataXML(t *testing.T) {
	// Second pawn.
	se := Doc.FindElement(sproot)
	spdoc := etree.NewDocumentWithRoot(se)
	dataXML, err := spdoc.WriteToString()
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	sp, err := NewPawnFromDataXML(dataXML, "second")
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	if sp.Attributes.Name != "Leora" {
		t.Errorf("sp.Attributes.Name: want Leora, got: %v", sp.Attributes.Name)
	}
	if sp.Attributes.Nickname.Name != "mNickname" {
		t.Errorf("sp.Attributes.Nickname.Value: want mNickname, got: %v", sp.Attributes.Nickname.Name)
	}
	if sp.Attributes.Nickname.Value != 525 {
		t.Errorf("sp.Attributes.Nickname.Value: want 525, got: %v", sp.Attributes.Nickname.Value)
	}

	// Main pawn.
	me := Doc.FindElement(mproot)
	mpdoc := etree.NewDocumentWithRoot(me)
	dataXML, err = mpdoc.WriteToString()
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	mp, err := NewPawnFromDataXML(dataXML, "main")
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	// Main pawn appearance override.
	aoroot := Doc.FindElement(nodemap.MainPawnAppearanceOverrideParent())
	if aoroot == nil {
		t.Fatal("invalid main pawn appearance override root")
	}
	appOver, err := NewAppearanceOverrideFromRootNode(aoroot, nodemap.Nodemap()["mEditPawn"])
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}
	mp.AppearanceOverride = *appOver

	if mp.Attributes.Name != "Mainpawn" {
		t.Errorf("mp.Attributes.Name: want Mainpawn, got: %v", mp.Attributes.Name)
	}
	if mp.Attributes.Nickname.Name != "mNickname" {
		t.Errorf("mp.Attributes.Nickname.Value: want mNickname, got: %v", mp.Attributes.Nickname.Name)
	}
	if mp.Attributes.Nickname.Value != 0 {
		t.Errorf("mp.Attributes.Nickname.Value: want 0, got: %v", mp.Attributes.Nickname.Value)
	}
	if mp.AppearanceOverride.Attributes.Name != "\v\x16!,7BMXc" {
		t.Errorf("mp.AppearanceOverride.Attributes.Name: want \v\u0016!,7BMXc, got: %#v", mp.AppearanceOverride.Attributes.Name)
	}
	if mp.AppearanceOverride.Attributes.Nickname.Name != "mNickname" {
		t.Errorf("mp.AppearanceOverride.Attributes.Nickname.Value: want mNickname, got: %v", mp.AppearanceOverride.Attributes.Nickname.Name)
	}
	if mp.AppearanceOverride.Attributes.Nickname.Value != 99 {
		t.Errorf("mp.AppearanceOverride.Attributes.Nickname.Value: want 99, got: %v", mp.AppearanceOverride.Attributes.Nickname.Value)
	}
}

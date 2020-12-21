package entities

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/beevik/etree"

	"github.com/var512/ddda-save-corrupter/internal/nodemap"
	"github.com/var512/ddda-save-corrupter/internal/parser"
	"github.com/var512/ddda-save-corrupter/internal/types"
	"github.com/var512/ddda-save-corrupter/internal/util"
)

// PawnCategory is an enum of pawn categories.
type PawnCategory string

const (
	MainPawn   PawnCategory = "MAIN_PAWN"
	FirstPawn  PawnCategory = "FIRST_PAWN"
	SecondPawn PawnCategory = "SECOND_PAWN"
)

type Attributes struct {
	Name     string
	NameStr  types.U8Array `json:"-"`
	Nickname types.U32
	Gender   types.U8
}

// mEditPawn (main pawn appearance override).
// TODO remove DataXML and use Element.
type AppearanceOverride struct {
	Element    *etree.Element `json:"-"`
	DataXML    string
	Attributes Attributes
}

type Pawn struct {
	Category           PawnCategory
	Element            *etree.Element `json:"-"`
	Attributes         Attributes
	AppearanceOverride AppearanceOverride
}

// NewPawnFromRootNode creates a new pawn from their root node.
func NewPawnFromRootNode(root *etree.Element, c PawnCategory) (*Pawn, error) {
	p := &Pawn{
		Category: c,
	}

	// Element here has everything from {Category}PawnParent
	// class name="mEdit" [...] class name="mParam" [...]
	// because it simplifies import/export and has all the
	// relevant pawn data.
	e := root.Copy()
	p.Element = e

	err := p.parseAttributesFromRootNode(e)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// parse parses attributes from a root element.
func (a *Attributes) parse(root *etree.Element) error {
	nameStr, err := parser.GetU8Array(root, `array[@name='(u8*)mNameStr']`)
	if err != nil {
		return err
	}
	a.NameStr = *nameStr
	a.Name = NameStrToString(a.NameStr)

	nickname, err := parser.GetU32(root, "u32[@name='mNickname']")
	if err != nil {
		return err
	}
	a.Nickname = *nickname

	gender, err := parser.GetU8(root, "u8[@name='mGender']")
	if err != nil {
		return err
	}
	a.Gender = *gender

	return nil
}

// parseAttributesFromRootNode parses pawn attributes.
func (p *Pawn) parseAttributesFromRootNode(root *etree.Element) error {
	subroot := root.FindElement(nodemap.Nodemap()["mEdit"])
	if subroot == nil {
		return errors.New("invalid main pawn subroot node")
	}

	err := p.Attributes.parse(subroot)
	if err != nil {
		return err
	}

	return nil
}

// NewAppearanceOverrideFromRootNode creates a new appearance override from their root node.
// Parent: mEdit, mEditPawn, so it can be reused without reparsing AOs from different sources.
func NewAppearanceOverrideFromRootNode(root *etree.Element, parent string) (*AppearanceOverride, error) {
	subroot := root.FindElement(parent)
	if subroot == nil {
		return nil, errors.New("new AO: invalid main pawn appearance override subroot node")
	}

	doc := etree.NewDocument()
	for _, e := range subroot.ChildElements() {
		doc.AddChild(e.Copy())
	}

	ao := &AppearanceOverride{
		Element: &doc.Element,
	}

	err := ao.parseAttributes()
	if err != nil {
		return nil, err
	}

	return ao, nil
}

// parseAttributesFromRootNode parses appearance override attributes.
func (ao *AppearanceOverride) parseAttributes() error {
	subroot := ao.Element
	if subroot == nil {
		return errors.New("parseAttributes: invalid main pawn appearance override subroot node")
	}

	// dataXML here, unlike Pawn that has the entire parent contents,
	// has only the appearance override values from mEditPawn.
	// <array name [...].
	dataXML, err := util.ChildrenToString(subroot, 4)
	if err != nil {
		return err
	}
	ao.DataXML = dataXML

	err = ao.Attributes.parse(subroot)
	if err != nil {
		return err
	}

	return nil
}

// NewPawnFromDataXML creates a new pawn from DataXML (string of the root node of a pawn).
func NewPawnFromDataXML(dataXML string, c string) (*Pawn, error) {
	category, err := ToPawnCategory(c)
	if err != nil {
		return nil, err
	}

	root := etree.NewDocument()
	err = root.ReadFromString(dataXML)
	if err != nil {
		return nil, err
	}

	p, err := NewPawnFromRootNode(&root.Element, *category)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ToDataXML returns the pawn tree elements DataXML.
func (p *Pawn) ToDataXML() (string, error) {
	doc := etree.NewDocument()
	for _, e := range p.Element.ChildElements() {
		doc.AddChild(e.Copy())
	}

	doc.Indent(4)

	data, err := doc.WriteToString()
	if err != nil {
		return "", err
	}

	return data, nil
}

// ToBytes returns the pawn tree elements DataXML as bytes.
func (p *Pawn) ToBytes() ([]byte, error) {
	doc := etree.NewDocument()

	for _, e := range p.Element.ChildElements() {
		doc.AddChild(e.Copy())
	}

	doc.Indent(4)

	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToPawnCategory returns the PawnCategory from a string (the frontend representation).
func ToPawnCategory(c string) (*PawnCategory, error) {
	var pc PawnCategory
	switch c {
	case "main":
		pc = MainPawn
	case "first":
		pc = FirstPawn
	case "second":
		pc = SecondPawn
	default:
		return nil, errors.New("invalid pawn category: " + c)
	}

	return &pc, nil
}

// NameStrToString converts a NameStr U8Array to a string.
func NameStrToString(nameStr types.U8Array) string {
	var s strings.Builder

	var b []byte
	for _, child := range nameStr.Children {
		if child.Value < 1 {
			break
		}
		b = append(b, child.Value)
	}

	l := len(b)
	size := l
	for i := 0; i < l; {
		str, size := utf8.DecodeRune(b[i:size])
		s.WriteString(fmt.Sprintf("%c", str))
		i += size
	}

	return s.String()
}

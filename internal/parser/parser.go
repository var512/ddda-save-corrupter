package parser

import (
	"errors"

	"github.com/beevik/etree"

	"github.com/var512/ddda-save-corrupter/internal/types"
)

var (
	errInvalidNodeAttr = errors.New("parser: invalid node attribute")
)

// TODO:
// Auto generated types: make sure they work properly.
// The encoding/xml package is too limited for un/marshaling and parsing.
// Hacking a type registry with interfaces, reflection and assertions still
// requires a lot of boilerplate at the cost of extra complexity.

// parse parses a node.
func parse(root *etree.Element, expr string) (map[string]string, error) {
	m := map[string]string{}

	q := root.FindElement(expr)
	if q == nil {
		return nil, errors.New("parser: invalid node expression " + expr)
	}

	m["name"] = q.SelectAttrValue("name", "")
	if m["name"] == "" {
		return nil, errInvalidNodeAttr
	}

	m["value"] = q.SelectAttrValue("value", "")
	if m["value"] == "" {
		return nil, errInvalidNodeAttr
	}

	return m, nil
}

// parseArray parses a node array and its child elements.
func parseArray(root *etree.Element, expr string) (map[string]string, error) {
	m := map[string]string{}

	q := root.FindElement(expr)
	if q == nil {
		return nil, errors.New("parser: invalid node expression " + expr)
	}

	m["name"] = q.SelectAttrValue("name", "")
	if m["name"] == "" {
		return nil, errInvalidNodeAttr
	}

	m["t"] = q.SelectAttrValue("type", "")
	if m["t"] == "" {
		return nil, errInvalidNodeAttr
	}

	m["count"] = q.SelectAttrValue("count", "")
	if m["count"] == "" {
		return nil, errInvalidNodeAttr
	}

	return m, nil
}

// GetU8 finds the element matching the xpath expression in the root node.
func GetU8(root *etree.Element, expr string) (*types.U8, error) {
	m, err := parse(root, expr)
	if err != nil {
		return nil, err
	}

	n, err := types.ToU8(m["name"], m["value"])
	if err != nil {
		return nil, err
	}

	return n, nil
}

// GetU8Array finds the element matching the xpath expression in the root node
// including their child elements.
func GetU8Array(root *etree.Element, expr string) (*types.U8Array, error) {
	m, err := parseArray(root, expr)
	if err != nil {
		return nil, err
	}

	a := &types.U8Array{
		Name:  m["name"],
		Type:  m["t"],
		Count: m["count"],
	}

	qc := root.FindElements(expr + "/u8")
	if len(qc) < 1 {
		return nil, errors.New("parser: invalid node expression " + expr)
	}

	var children []types.U8

	for _, v := range qc {
		el, err := types.ToU8("", v.SelectAttrValue("value", ""))
		if err != nil {
			return nil, err
		}

		children = append(children, *el)
	}

	a.Children = children

	return a, nil
}

// GetU32 finds the element matching the xpath expression in the root node.
func GetU32(root *etree.Element, expr string) (*types.U32, error) {
	m, err := parse(root, expr)
	if err != nil {
		return nil, err
	}

	n, err := types.ToU32(m["name"], m["value"])
	if err != nil {
		return nil, err
	}

	return n, nil
}

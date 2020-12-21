package util

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/beevik/etree"
)

// Uints32ToBytes converts []uint32 to []byte.
func Uints32ToBytes(u []uint32) []byte {
	l := len(u) * 4
	buf := make([]byte, l)

	for i, v := range u {
		binary.LittleEndian.PutUint32(buf[i*4:], v)
	}

	return buf
}

// ReaderToString returns a string from an io.Reader.
func ReaderToString(r io.Reader) (string, error) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ChildrenToString formats the root node tree as a DataXML string.
func ChildrenToString(root *etree.Element, indent int) (string, error) {
	var o strings.Builder

	for _, e := range root.ChildElements() {
		doc := etree.NewDocumentWithRoot(e.Copy())
		doc.Indent(indent)
		d, err := doc.WriteToString()
		if err != nil {
			return "", err
		}
		o.WriteString(d)
	}

	return o.String(), nil
}

// ReplaceChildren replaces the children of a node with the children of another.
func ReplaceChildren(root *etree.Element, n *etree.Element) {
	// Remove children from the root node.
	RemoveChildren(root)
	// Add a deep copy of the children of n to root.
	AddChildren(root, n.Copy())
}

// AddChildren adds the children of a node to another.
func AddChildren(root *etree.Element, n *etree.Element) {
	for _, e := range n.ChildElements() {
		root.AddChild(e)
	}
}

// RemoveChildren removes child elements of a node.
func RemoveChildren(root *etree.Element) {
	for _, e := range root.ChildElements() {
		root.RemoveChild(e)
	}
}

// PrintlnChildren prints the child elements of a node.
func PrintlnChildren(root *etree.Element) {
	for k, e := range root.ChildElements() {
		fmt.Println(k, e)
	}
}

// Dump prints stuff as a serialized (json) string.
func Dump(i interface{}) {
	s, _ := json.Marshal(i)
	fmt.Println("----------------------------------------------------")
	fmt.Println(string(s))
	fmt.Println()
}

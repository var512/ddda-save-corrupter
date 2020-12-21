package types

import (
	"strconv"
)

type F32 struct {
	Name  string
	Value float32
}

type F32Array struct {
	Name     string
	Type     string
	Count    string
	Children []F32
}

type S8 struct {
	Name  string
	Value int8
}

type S8Array struct {
	Name     string
	Type     string
	Count    string
	Children []S8
}

type S16 struct {
	Name  string
	Value int16
}

type S16Array struct {
	Name     string
	Type     string
	Count    string
	Children []S16
}

type S32 struct {
	Name  string
	Value int32
}

type S32Array struct {
	Name     string
	Type     string
	Count    string
	Children []S32
}

type U8 struct {
	Name  string
	Value uint8
}

type U8Array struct {
	Name     string
	Type     string
	Count    string
	Children []U8
}

type U16 struct {
	Name  string
	Value uint16
}

type U16Array struct {
	Name     string
	Type     string
	Count    string
	Children []U16
}

type U32 struct {
	Name  string
	Value uint32
}

type U32Array struct {
	Name     string
	Type     string
	Count    string
	Children []U32
}

type U64 struct {
	Name  string
	Value uint64
}

type U64Array struct {
	Name     string
	Type     string
	Count    string
	Children []U64
}

func ToF32(name, value string) (*F32, error) {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil, err
	}

	n := &F32{
		Name:  name,
		Value: float32(v),
	}

	return n, nil
}

func ToS8(name, value string) (*S8, error) {
	v, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return nil, err
	}

	n := &S8{
		Name:  name,
		Value: int8(v),
	}

	return n, nil
}

func ToS16(name, value string) (*S16, error) {
	v, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return nil, err
	}

	n := &S16{
		Name:  name,
		Value: int16(v),
	}

	return n, nil
}

func ToS32(name, value string) (*S32, error) {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}

	n := &S32{
		Name:  name,
		Value: int32(v),
	}

	return n, nil
}

func ToU8(name, value string) (*U8, error) {
	v, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return nil, err
	}

	n := &U8{
		Name:  name,
		Value: uint8(v),
	}

	return n, nil
}

func ToU16(name, value string) (*U16, error) {
	v, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return nil, err
	}

	n := &U16{
		Name:  name,
		Value: uint16(v),
	}

	return n, nil
}

func ToU32(name, value string) (*U32, error) {
	v, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return nil, err
	}

	n := &U32{
		Name:  name,
		Value: uint32(v),
	}

	return n, nil
}

func ToU64(name, value string) (*U64, error) {
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}

	n := &U64{
		Name:  name,
		Value: v,
	}

	return n, nil
}

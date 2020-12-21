package sav

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	Contents []byte
)

func TestMain(m *testing.M) {
	f, err := openReadFile("../testdata/base.sav")
	if err != nil {
		panic(err)
	}
	Contents = *f

	os.Exit(m.Run())
}

func TestHasValidContentAndHeader(t *testing.T) {
	h, err := NewHeaderFromContents(Contents)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	err = h.Validate()
	if err != nil {
		t.Errorf("unexpected error: %#v", err.Error())
	}
}

func TestUncompressZlibAndGetDataXML(t *testing.T) {
	dataXML, err := GetDataXMLFromUserfile(Contents)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	expected := []string{
		`<?xml version="1.0" encoding="utf-8"?>`,
		`<class name="dd_savedata1018" type="sSave::saveDataAllDA">`,
		`<class name="mSystemData" type="sSave::systemData">`,
		`<bool name="mClearSaveData"`,
		`<bool name="mHaveTrialEditData"`,
	}

	for _, s := range expected {
		t.Run("Expects "+s, func(t *testing.T) {
			if !strings.Contains(dataXML, s) {
				t.Errorf("invalid xml, expected: %#v", s)
			}
		})
	}
}

func openReadFile(path string) (*[]byte, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := fd.Close(); err != nil {
		return nil, err
	}

	return &c, nil
}

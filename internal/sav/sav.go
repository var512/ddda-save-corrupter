package sav

import (
	"bytes"
	"compress/flate"
	"compress/zlib"

	"github.com/var512/ddda-save-corrupter/internal/util"
)

// Size of full file is always 524288 bytes (extra data are nulls).
const UserfileExpectedLength = 524288

type Sav struct {
	Header  Header
	DataXML string
}

// GetDataXMLFromUserfile extracts the XML data from a compressed
// userfile (a .sav that already has a valid header and the compressed data).
// It skips the header bytes.
func GetDataXMLFromUserfile(f []byte) (string, error) {
	ct := f[HeaderExpectedLength:]

	b := bytes.NewReader(ct)
	r, err := zlib.NewReader(b)
	if err != nil {
		return "", err
	}
	defer r.Close()

	data, err := util.ReaderToString(r)
	if err != nil {
		return "", err
	}

	if err := r.Close(); err != nil {
		return "", err
	}

	return data, nil
}

// Compress compresses with zlib.
func Compress(ct []byte) ([]byte, error) {
	var buf bytes.Buffer

	w, err := zlib.NewWriterLevel(&buf, flate.DefaultCompression)

	if err != nil {
		return nil, err
	}
	_, err = w.Write(ct)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

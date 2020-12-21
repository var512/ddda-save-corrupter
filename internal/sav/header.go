package sav

import (
	"encoding/binary"
	"fmt"

	"github.com/var512/ddda-save-corrupter/internal/logger"
	"github.com/var512/ddda-save-corrupter/internal/util"
)

// Total size (bytes).
const HeaderExpectedLength = 32

type Header struct {
	Version        uint32 // PC = 21.
	Size           uint32 // Size of the uncompressed data.
	CompressedSize uint32 // Size of the compressed data.
	UH1            uint32 // Always 860693325.
	UH2            uint32 // Always 0.
	UH3            uint32 // Always 860700740.
	Checksum       uint32 // Checksum of compressed save data (Bitwise not CRC32 Jam).
	UH4            uint32 // Always 1079398965.
}

// NewHeader creates a new Header from the contents of a .sav userfile.
func NewHeaderFromContents(ct []byte) (*Header, error) {
	l := len(ct)
	if l < HeaderExpectedLength {
		return nil, fmt.Errorf("header: expected length %v, got %v", HeaderExpectedLength, l)
	}

	h := &Header{}
	h.Version = binary.LittleEndian.Uint32(ct[0:])
	h.Size = binary.LittleEndian.Uint32(ct[4:])
	h.CompressedSize = binary.LittleEndian.Uint32(ct[8:])
	h.UH1 = binary.LittleEndian.Uint32(ct[12:])
	h.UH2 = binary.LittleEndian.Uint32(ct[16:])
	h.UH3 = binary.LittleEndian.Uint32(ct[20:])
	h.Checksum = binary.LittleEndian.Uint32(ct[24:])
	h.UH4 = binary.LittleEndian.Uint32(ct[28:])

	return h, nil
}

// Validate validates if the header has the expected content.
func (h *Header) Validate() error {
	if h.Version != 21 {
		return fmt.Errorf("wrong Version: got %v, expected 21", h.UH1)
	}

	if h.UH1 != 860693325 {
		return fmt.Errorf("wrong UH1: got %v, expected 860693325", h.UH1)
	}

	if h.UH2 != 0 {
		return fmt.Errorf("wrong UH2: got %v, expected 0", h.UH2)
	}

	if h.UH3 != 860700740 {
		return fmt.Errorf("wrong UH3: got %v, expected 860700740", h.UH3)
	}

	if h.UH4 != 1079398965 {
		return fmt.Errorf("wrong UH3: got %v, expected 1079398965", h.UH4)
	}

	return nil
}

// Serialize serializes a Header to []byte.
func (h *Header) Serialize() ([]byte, error) {
	d := make([]uint32, 0)
	d = append(d, h.Version)
	d = append(d, h.Size)
	d = append(d, h.CompressedSize)
	d = append(d, h.UH1)
	d = append(d, h.UH2)
	d = append(d, h.UH3)
	d = append(d, h.Checksum)
	d = append(d, h.UH4)

	return util.Uints32ToBytes(d), nil
}

// Print prints the Header structure for log debugging.
func (h *Header) Print() {
	logger.Default.Println("== Sav file header")
	logger.Default.Printf("==== Version: %v\n", h.Version)
	logger.Default.Printf("==== (Uncompressed) Size: %v\n", h.Size)
	logger.Default.Printf("==== CompressedSize: %v\n", h.CompressedSize)
	logger.Default.Printf("==== UH1: %v\n", h.UH1)
	logger.Default.Printf("==== UH2: %v\n", h.UH2)
	logger.Default.Printf("==== UH3: %v\n", h.UH3)
	logger.Default.Printf("==== Checksum: %v\n", h.Checksum)
	logger.Default.Printf("==== UH4: %v\n", h.UH4)
}

package nodemap

import (
	"strings"
)

// Nodemap maps the xpath of each element to a descriptive identifier.
// TODO make the map identifiers actually useful.
func Nodemap() map[string]string {
	return map[string]string{
		"root":               "class[@name='dd_savedata1018'][@type='sSave::saveDataAllDA']",
		"playerData":         "class[@name='mPlayerDataManual'][@type='sSave::playerData']",
		"mPlCmcEditAndParam": "class[@name='mPlCmcEditAndParam'][@type='sSave::playerEditAndParam']",
		"mCmc":               "array[@name='mCmc'][@type='class']",
		"cSAVE_DATA_CMC":     "class[@type='cSAVE_DATA_CMC']",
		"mEdit":              "class[@name='mEdit'][@type='cSAVE_DATA_EDIT']",
		"mSystemData":        "class[@name='mSystemData'][@type='sSave::systemData']",
		"mEditPawn":          "class[@name='mEditPawn'][@type='cSAVE_DATA_EDIT']",
	}
}

func PawnsParent() string {
	var s strings.Builder

	s.WriteString(Nodemap()["root"])
	s.WriteString("/")
	s.WriteString(Nodemap()["playerData"])
	s.WriteString("/")
	s.WriteString(Nodemap()["mPlCmcEditAndParam"])
	s.WriteString("/")
	s.WriteString(Nodemap()["mCmc"])

	return s.String()
}

func MainPawnParent() string {
	var s strings.Builder
	s.WriteString(PawnsParent())
	s.WriteString("/")
	s.WriteString(Nodemap()["cSAVE_DATA_CMC"])
	s.WriteString("[1]")

	return s.String()
}

func FirstPawnParent() string {
	var s strings.Builder
	s.WriteString(PawnsParent())
	s.WriteString("/")
	s.WriteString(Nodemap()["cSAVE_DATA_CMC"])
	s.WriteString("[2]")

	return s.String()
}

func SecondPawnParent() string {
	var s strings.Builder
	s.WriteString(PawnsParent())
	s.WriteString("/")
	s.WriteString(Nodemap()["cSAVE_DATA_CMC"])
	s.WriteString("[3]")

	return s.String()
}

func MainPawnAppearanceOverrideParent() string {
	var s strings.Builder

	s.WriteString(Nodemap()["root"])
	s.WriteString("/")
	s.WriteString(Nodemap()["mSystemData"])

	return s.String()
}

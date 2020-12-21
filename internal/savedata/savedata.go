package savedata

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/beevik/etree"

	"github.com/var512/ddda-save-corrupter/internal/entities"
	"github.com/var512/ddda-save-corrupter/internal/logger"
	"github.com/var512/ddda-save-corrupter/internal/nodemap"
)

type SaveData struct {
	// Doc stores the userfile XML data parsed as etree.Elements.
	// It shouldn't be modified, use a deep copy on export and operations.
	Doc *etree.Document

	MainPawn   entities.Pawn
	FirstPawn  entities.Pawn
	SecondPawn entities.Pawn
}

// NewSaveDataFromDataXML creates a SaveData from a string containing the DataXML.
func NewSaveDataFromDataXML(s *string) (*SaveData, error) {
	sd := &SaveData{}

	doc := etree.NewDocument()
	_, err := doc.ReadFrom(strings.NewReader(*s))
	if err != nil {
		log.Fatal(err)
	}

	sd.Doc = doc

	err = sd.parseAllPawns()
	if err != nil {
		return nil, err
	}

	return sd, nil
}

// PawnRootCount gets the count attribute from the parent array of the pawn tree.
// Tested with a new save and it's always 3, with zeroed values.
func (sd *SaveData) PawnRootCount() (uint8, error) {
	root := sd.Doc.FindElement(nodemap.PawnsParent())
	if root == nil {
		return 0, errors.New("invalid root pawn node")
	}

	attr := root.SelectAttrValue("count", "")
	if attr == "" {
		return 0, errors.New("invalid root pawn node count")
	}

	c, err := strconv.ParseUint(attr, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(c), nil
}

// parseAllPawns parses data from all pawns.
func (sd *SaveData) parseAllPawns() error {
	count, err := sd.PawnRootCount()
	if err != nil {
		return err
	}

	if count != 3 {
		return fmt.Errorf("pawns array count: got %v, expected %v", count, 3)
	}

	err = sd.parseMainPawn()
	if err != nil {
		return err
	}
	logger.Default.Println("== Parsed MainPawn")

	err = sd.parseFirstPawn()
	if err != nil {
		return err
	}
	logger.Default.Println("== Parsed FirstPawn")

	err = sd.parseSecondPawn()
	if err != nil {
		return err
	}
	logger.Default.Println("== Parsed SecondPawn")

	return nil
}

// parseMainPawn parses data from the main pawn.
func (sd *SaveData) parseMainPawn() error {
	root := sd.Doc.FindElement(nodemap.MainPawnParent())
	if root == nil {
		return errors.New("invalid main pawn root")
	}

	pawn, err := entities.NewPawnFromRootNode(root, entities.MainPawn)
	if err != nil {
		return err
	}

	// Appearance override.
	aoroot := sd.Doc.FindElement(nodemap.MainPawnAppearanceOverrideParent())
	if aoroot == nil {
		return errors.New("invalid main pawn appearance override root")
	}

	appOver, err := entities.NewAppearanceOverrideFromRootNode(aoroot, nodemap.Nodemap()["mEditPawn"])
	if err != nil {
		return err
	}

	pawn.AppearanceOverride = *appOver

	sd.MainPawn = *pawn

	return nil
}

// parseFirstPawn parses data from the first pawn.
func (sd *SaveData) parseFirstPawn() error {
	root := sd.Doc.FindElement(nodemap.FirstPawnParent())
	if root == nil {
		return errors.New("invalid first pawn root")
	}

	pawn, err := entities.NewPawnFromRootNode(root, entities.FirstPawn)
	if err != nil {
		return err
	}
	sd.FirstPawn = *pawn

	return nil
}

// parseSecondPawn parses data from the second pawn.
func (sd *SaveData) parseSecondPawn() error {
	root := sd.Doc.FindElement(nodemap.SecondPawnParent())
	if root == nil {
		return errors.New("invalid second pawn root")
	}

	pawn, err := entities.NewPawnFromRootNode(root, entities.SecondPawn)
	if err != nil {
		return err
	}
	sd.SecondPawn = *pawn

	return nil
}

// stackoverflow.com/a/17306470 and more annoyances.
// func (sd *SaveData) MarshalJSON() ([]byte, error).
func (sd *SaveData) GetPawnMap(category string) (map[string]interface{}, error) {
	var v map[string]interface{}

	switch category {
	case "main":
		dataXML, err := sd.MainPawn.ToDataXML()
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"pawn": map[string]interface{}{
				"data":    sd.MainPawn,
				"dataXML": dataXML,
			},
		}
	case "first":
		dataXML, err := sd.FirstPawn.ToDataXML()
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"pawn": map[string]interface{}{
				"data":    sd.FirstPawn,
				"dataXML": dataXML,
			},
		}
	case "second":
		dataXML, err := sd.SecondPawn.ToDataXML()
		if err != nil {
			return nil, err
		}
		v = map[string]interface{}{
			"pawn": map[string]interface{}{
				"data":    sd.SecondPawn,
				"dataXML": dataXML,
			},
		}
	default:
		return nil, errors.New("invalid pawn category")
	}

	return v, nil
}

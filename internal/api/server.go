package api

import (
	"bytes"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/var512/ddda-save-corrupter/internal/assets"
	"github.com/var512/ddda-save-corrupter/internal/entities"
	"github.com/var512/ddda-save-corrupter/internal/logger"
	"github.com/var512/ddda-save-corrupter/internal/nodemap"
	"github.com/var512/ddda-save-corrupter/internal/parser"
	"github.com/var512/ddda-save-corrupter/internal/sav"
	"github.com/var512/ddda-save-corrupter/internal/savedata"
	"github.com/var512/ddda-save-corrupter/internal/util"
)

const (
	Host = "localhost"
	Port = "3333"
)

var CorsAllowOrigin = "http://" + Host + ":" + Port

// State for forced single session with no client storage.
type State struct {
	// Parsed data from the sav userfile.
	SaveData savedata.SaveData
}

var ServerState State

// Start starts the API server.
func Start() {
	logger.Default.Println("== [Web UI] http://" + Host + ":" + Port)
	logger.Default.Println("============================================================")
	logger.Default.Println("== Warning: backup your .sav files before using this tool")
	logger.Default.Println("============================================================")

	r := registerRoutes()

	log.Fatal(http.ListenAndServe(Host+":"+Port, r))
}

// registerRoutes register routes with middleware and handlers.
func registerRoutes() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(false)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Methods("OPTIONS")
	api.Use(loggingMiddleware)
	api.Use(preflightMiddleware)
	api.Use(errorMiddleware)
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorJSON(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	api.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorJSON(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})
	api.HandleFunc("/files", importFile).Methods("POST")
	api.HandleFunc("/files/{format:sav|xml}/export", exportFile).Methods("GET")
	api.HandleFunc("/pawns", importPawn).Methods("POST")
	api.HandleFunc("/pawns/{category:main|first|second}", pawns).Methods("GET")
	api.HandleFunc("/pawns/{category:main|first|second}/export", exportPawn).Methods("GET")
	api.HandleFunc("/pawns/main/export-with-appearance-override", exportPawnWithAO).Methods("GET")
	api.HandleFunc("/status", status).Methods("GET")

	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(assets.Assets).ServeHTTP(w, r)
	}))

	return r
}

// importFile imports a .sav userfile.
func importFile(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	file, header, err := r.FormFile("userfile")
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}
	defer file.Close()

	logger.Default.Printf("== Importing file: %s\n", header.Filename)

	if _, err := io.Copy(&buf, file); err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	if err := file.Close(); err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	l := buf.Len()
	if l != sav.UserfileExpectedLength {
		err := fmt.Errorf("wrong file size: got %v, expected %v", l, sav.UserfileExpectedLength)
		ErrorJSON(w, err.Error(), 500)
		return
	}

	dataXML, err := sav.GetDataXMLFromUserfile(buf.Bytes())
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	sd, err := savedata.NewSaveDataFromDataXML(&dataXML)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}
	ServerState.SaveData = *sd

	buf.Reset()

	logger.Default.Println("== Success: file imported")
	ResponseJSON(w, "Success: file imported")
}

// exportFile exports the savedata as a .sav or .xml file.
func exportFile(w http.ResponseWriter, r *http.Request) {
	// TODO use validUserfileMiddleware.
	if ServerState.SaveData.Doc == nil {
		ErrorJSON(w, "userfile is empty", 500)
		return
	}

	var err error
	var format string
	var file []byte
	vars := mux.Vars(r)

	switch vars["format"] {
	case "sav":
		format = "sav"
		file, err = fileSAV()
		if err != nil {
			ErrorJSON(w, err.Error(), 500)
			return
		}
	case "xml":
		format = "xml"
		file, err = fileXML()
		if err != nil {
			ErrorJSON(w, err.Error(), 500)
			return
		}
	default:
		ErrorJSON(w, "invalid file format: "+format, 500)
		return
	}

	t := time.Now().Format("2006_01_02_150405")
	filename := strings.Join([]string{"DDDASC_", t, ".", format}, "")

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	_, err = w.Write(file)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	logger.Default.Println("== File exported:", filename)
}

// fileXML returns the contents representing the savedata as a .xml file.
// Modifies a copy of the original Doc with the changes made by the user.
func fileXML() ([]byte, error) {
	var err error
	doc := ServerState.SaveData.Doc.Copy()

	mproot := doc.FindElement(nodemap.MainPawnParent())
	util.ReplaceChildren(mproot, ServerState.SaveData.MainPawn.Element)

	fproot := doc.FindElement(nodemap.FirstPawnParent())
	util.ReplaceChildren(fproot, ServerState.SaveData.FirstPawn.Element)

	sproot := doc.FindElement(nodemap.SecondPawnParent())
	util.ReplaceChildren(sproot, ServerState.SaveData.SecondPawn.Element)

	// Main pawn appearance override: copy from the SaveData.
	aoroot := doc.FindElement(nodemap.MainPawnAppearanceOverrideParent() + "/" + nodemap.Nodemap()["mEditPawn"])
	if aoroot == nil {
		return nil, errors.New("filexml: invalid main pawn appearance override root")
	}
	sdroot := ServerState.SaveData.MainPawn.AppearanceOverride.Element
	util.ReplaceChildren(aoroot, sdroot)

	// Indent = 0 is closer to an original .sav file and the
	// compressed size doesn't go over UserfileExpectedLength.
	doc.Indent(0)
	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// fileSAV returns the contents representing the savedata as a .sav file with a valid header.
func fileSAV() ([]byte, error) {
	uncompressedData, err := fileXML()
	if err != nil {
		return nil, err
	}
	size := len(uncompressedData)

	compressedData, err := sav.Compress(uncompressedData)
	if err != nil {
		return nil, err
	}

	compressedSize := len(compressedData)
	checksum := ^crc32.ChecksumIEEE(compressedData)

	h := sav.Header{
		Version:        21,
		Size:           uint32(size),
		CompressedSize: uint32(compressedSize),
		UH1:            860693325,
		UH2:            0,
		UH3:            860700740,
		Checksum:       checksum,
		UH4:            1079398965,
	}

	err = h.Validate()
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	headerData, err := h.Serialize()
	if err != nil {
		return nil, err
	}

	data = append(data, headerData...)
	data = append(data, compressedData...)

	paddingSize := sav.UserfileExpectedLength - len(headerData) - len(compressedData)
	padding := make([]byte, paddingSize)
	data = append(data, padding...)

	l := len(data)
	if l != sav.UserfileExpectedLength {
		return nil, fmt.Errorf("wrong file size: got %v, expected %v", l, sav.UserfileExpectedLength)
	}

	h.Print()

	return data, nil
}

func importPawn(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")

	var buf bytes.Buffer

	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	file, header, err := r.FormFile("pawnuserfile")
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}
	defer file.Close()

	logger.Default.Printf("== Importing pawn file: %s\n", header.Filename)

	if _, err := io.Copy(&buf, file); err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	if err := file.Close(); err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	p, err := entities.NewPawnFromDataXML(buf.String(), category)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	// Update the main pawn appearance override with the appearance from the imported pawn.
	if category == "main" {
		aoroot := p.Element
		if aoroot == nil {
			ErrorJSON(w, "import pawn: invalid main pawn appearance override root", 500)
			return
		}

		appOver, err := entities.NewAppearanceOverrideFromRootNode(aoroot, nodemap.Nodemap()["mEdit"])
		if err != nil {
			ErrorJSON(w, err.Error(), 500)
			return
		}

		p.AppearanceOverride = *appOver
	}

	switch category {
	case "main":
		ServerState.SaveData.MainPawn = *p
	case "first":
		ServerState.SaveData.FirstPawn = *p
	case "second":
		ServerState.SaveData.SecondPawn = *p
	default:
		ErrorJSON(w, "import pawn: invalid category", 500)
		return
	}

	buf.Reset()

	logger.Default.Println("== Success: pawn file imported")
	ResponseJSON(w, "Success: pawn file imported")
}

func exportPawn(w http.ResponseWriter, r *http.Request) {
	// TODO use validUserfileMiddleware.
	if ServerState.SaveData.Doc == nil {
		ErrorJSON(w, "export pawn: userfile is empty", 500)
		return
	}

	var err error
	var p entities.Pawn
	var file []byte
	vars := mux.Vars(r)

	switch vars["category"] {
	case "main":
		p = ServerState.SaveData.MainPawn
	case "first":
		p = ServerState.SaveData.FirstPawn
	case "second":
		p = ServerState.SaveData.SecondPawn
	default:
		ErrorJSON(w, "export pawn: invalid pawn category", 500)
		return
	}

	file, err = p.ToBytes()
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	t := time.Now().Format("2006_01_02_150405")

	var filename strings.Builder
	filename.WriteString("DDDASC_")
	filename.WriteString(string(p.Category))
	filename.WriteString("_")
	filename.WriteString(p.Attributes.Name)
	filename.WriteString("_")
	filename.WriteString(t)
	filename.WriteString(".xml")

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename.String()))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	_, err = w.Write(file)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	logger.Default.Println("== Pawn file exported:", filename.String())
}

func exportPawnWithAO(w http.ResponseWriter, r *http.Request) {
	// TODO use validUserfileMiddleware.
	if ServerState.SaveData.Doc == nil {
		ErrorJSON(w, "export pawn: userfile is empty", 500)
		return
	}

	var err error
	var file []byte

	sdroot := ServerState.SaveData.MainPawn.Element.Copy()
	subroot := sdroot.FindElement(nodemap.Nodemap()["mEdit"])
	if subroot == nil {
		ErrorJSON(w, "invalid main pawn subroot", 500)
		return
	}

	aoroot := ServerState.SaveData.MainPawn.AppearanceOverride.Element

	util.ReplaceChildren(subroot, aoroot)

	nameStr, err := parser.GetU8Array(aoroot, `array[@name='(u8*)mNameStr']`)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}
	name := entities.NameStrToString(*nameStr)

	// A copy with the necessary information to generate the .xml
	// without modifying the original pawn or making a deep copy.
	pcopy := entities.Pawn{
		Category: entities.MainPawn,
		Element:  sdroot,
		Attributes: entities.Attributes{
			Name: name,
		},
	}

	file, err = pcopy.ToBytes()
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	t := time.Now().Format("2006_01_02_150405")

	var filename strings.Builder
	filename.WriteString("DDDASC_MAIN_PAWN_WITH_AO_")
	filename.WriteString(pcopy.Attributes.Name)
	filename.WriteString("_")
	filename.WriteString(t)
	filename.WriteString(".xml")

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename.String()))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	_, err = w.Write(file)
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	logger.Default.Println("== Pawn with appearance override file exported:", filename.String())
}

func pawns(w http.ResponseWriter, r *http.Request) {
	// TODO use validUserfileMiddleware.
	if ServerState.SaveData.Doc == nil {
		ErrorJSON(w, "pawns: userfile is empty", 500)
		return
	}

	vars := mux.Vars(r)
	data, err := ServerState.SaveData.GetPawnMap(vars["category"])
	if err != nil {
		ErrorJSON(w, err.Error(), 500)
		return
	}

	ResponseJSON(w, data)
}

func status(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, "pong")
}

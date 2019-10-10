package items

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	io2 "gostories/engine/io"
)

// These consts are used in producing Go source files
const (
	generatedGoFileSuffix = ".gen.go"

	// Thing Types
	TypeItem    = "Item"
	TypeFeature = "Feature"
	TypeBeing   = "Being"

	// package Names
	PckgItem         = "items"
	PckgFeatures     = "features"
	packageStatement = "package %v\n\n"
	importThings     = "import \"gostories/things\"\n"
	importItems      = "import \"gostories/gen/items\"\n"
)

type itemToAdd struct {
	Name         string
	LookText     string
	IsVisible    bool
	IsEquippable bool
}

// WriteItemsFile outputs a Go source file containing the in-game Items to generate
func WriteItemsFile() error {
	b := &bytes.Buffer{}
	addAutogeneratedWarning(b)
	addPackagesInfo(b, PckgItem, false)
	items := []itemToAdd{
		itemToAdd{"collar", "A small red cat collar with a bell.", false, true},
		itemToAdd{"shrubbery", "A small but rather well cared for shrubbery.", true, false},
		itemToAdd{"sardines", "A tin of tasty sardines preserved in olive oil.", true, false},
	}
	addItemMap(items, b)
	addItemStructs(items, b)
	return genGoLangFile(PckgItem, b)
}

// WriteFeaturesFile outputs a Go source file containing the in-game Features to generate
func WriteFeaturesFile() error {
	b := &bytes.Buffer{}
	addAutogeneratedWarning(b)
	//TODO get rid of this add items boolean because its shit. Will probably move feature and item auto-generation
	// into seperate files
	addPackagesInfo(b, PckgFeatures, false)
	addFeatureStruct(
		"shelf",
		"The shelf seems to contain a few old magazines and a cat collar.",
		"reveal-item(collar)",
		b,
	)
	addFeatureStruct(
		"fridge",
		"The fridge is empty apart from a tin of sardines.",
		"reveal-item(sardines)",
		b,
	)
	return genGoLangFile(PckgFeatures, b)
}

func genGoLangFile(packageName string, buffer *bytes.Buffer) (err error) {
	filename := "./" + packageName + "/" + generatedFileName(packageName)
	io2.ActiveInputOutputHandler.NewLinef("Filename: %v", filename)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0755)
	return
}

func addAutogeneratedWarning(writer *bytes.Buffer) {
	fileStr := `
// Do not edit! Autogenerated file //

`
	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

// TODO changed package name to items for now, but should decide on a new name
func addPackagesInfo(buffer *bytes.Buffer, pckgName string, addItems bool) {
	imports := importThings
	if addItems {
		imports = imports + importItems
	}
	buffer.WriteString(fmt.Sprintf(packageStatement, pckgName) + imports)
}

func addItemMap(items []itemToAdd, buffer *bytes.Buffer) {
	mapStr := `
// Items contains all the items a player can pick up. It is currently indexed by the item name, however ideally 
// it should be indexed by an enum/custom int and this should be the only place to access items.
var Items = map[string]things.Item {
`
	for _, item := range items {
		upperCaseItemName := strings.ToUpper(item.Name[0:1]) + item.Name[1:]
		mapStr = mapStr + "        \"" + item.Name + "\": Item" + upperCaseItemName + ",\n"
	}
	mapStr = mapStr + `}

`
	buffer.WriteString(mapStr)
}

func addItemStructs(items []itemToAdd, buffer *bytes.Buffer) {
	for _, item := range items {
		addItemStruct(item, buffer)
	}
}

func addItemStruct(item itemToAdd, buffer *bytes.Buffer) {
	upperCaseItemName := strings.ToUpper(item.Name[0:1]) + item.Name[1:]
	addBaseStructAndMethods(item.Name, upperCaseItemName, TypeItem, buffer)
	addMethodsForItems(upperCaseItemName, buffer, item.IsEquippable)
	addItemConstructor(item.Name, upperCaseItemName, item.LookText, TypeItem, item.IsVisible, buffer)
}

func addFeatureStruct(ftrName, lookText, trigger string, buffer *bytes.Buffer) {
	upperCaseFtrName := strings.ToUpper(ftrName[0:1]) + ftrName[1:]
	addBaseStructAndMethods(ftrName, upperCaseFtrName, TypeFeature, buffer)
	addFeatureConstructor(ftrName, upperCaseFtrName, lookText, TypeFeature, trigger, buffer)
}

func addBaseStructAndMethods(name, upperName, thingType string, buffer *bytes.Buffer) {
	addStruct(name, upperName, thingType, buffer)
	addBasicMethodsForAThing(upperName, thingType, buffer)
}

func addStruct(lowerName, upperName, thingType string, writer *bytes.Buffer) {
	fileStr := `
// [TT][N] probably should remove this and only access structs through the map
var [TT][N] = New[N][TT]()
var [lc_N][TT] *[N][TT]

// [N][TT] struct
type [N][TT] struct {
	things.Thing
}
`

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", lowerName, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

func addItemConstructor(name, upperName, lookText, thingType string, visible bool, buffer *bytes.Buffer) {
	fileStr := `
// New[N][TT] creates a new [N][TT]. Probably will unexport this soon.
func New[N][TT]() *[N][TT] {
	if [lc_N][TT] == nil {
		[lc_N][TT] = &[N][TT]{}
		[lc_N][TT].Name = "[lc_N]"
		[lc_N][TT].LookText = "[LT]"
	}
`
	if visible {
		fileStr = fileStr + `
	[lc_N][TT].Show()
`
	}

	fileStr = fileStr +
		`	return [lc_N][TT]
}
`

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", name, -1)
	fileStr = strings.Replace(fileStr, "[LT]", lookText, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)

	fsBytes := []byte(fileStr)
	buffer.Write(fsBytes)
}

// TODO: convert trigger string here into a type which maps a verb string to a trigger string
// TODO: then convert it into a slice of trigger strings so multiple can be added to a feature
func addFeatureConstructor(name, upperName, lookText, thingType, trigger string, buffer *bytes.Buffer) {
	fileStr := `
// New[N][TT] creates a new [N][TT]. Probably will unexport this soon.
func New[N][TT]() *[N][TT] {
	if [lc_N][TT] == nil {
		[lc_N][TT] = &[N][TT]{}
		[lc_N][TT].Name = "[lc_N]"
		[lc_N][TT].LookText = "[LT]"
    	[lc_N][TT].Triggers = map[string]string {
			"look": "[T]",
		}
	}
	
	[lc_N][TT].Show()
	return [lc_N][TT]
}
`

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", name, -1)
	fileStr = strings.Replace(fileStr, "[LT]", lookText, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)
	fileStr = strings.Replace(fileStr, "[T]", trigger, -1)

	fsBytes := []byte(fileStr)
	buffer.Write(fsBytes)
}

func addBasicMethodsForAThing(upperName, thingType string, writer *bytes.Buffer) {
	fileStr := `
// GetName returns the name of the thing
func (c *[N][TT]) GetName() string { return c.Name }

// GetLookText returns the description when the player looks at the thing
func (c *[N][TT]) GetLookText() string { return c.LookText }

// Show makes the thing visible to the player
func (c *[N][TT]) Show() { c.Thing.Visible = true }

// Hide makes the thing visible to the player
func (c *[N][TT]) Hide() { c.Thing.Visible = false }

// GetThing returns the underlying Thing struct (need to review if this is used)
func (c *[N][TT]) GetThing() *things.Thing { return &c.Thing }

`
	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

func addMethodsForItems(upperName string, writer *bytes.Buffer, isEquippable bool) {
	fileStr := `
// Take will be used for the player to take the item into the inventory (currently not needed)
func (c *[N]Item) Take() {}
`
	if isEquippable {
		fileStr = fileStr + `
// Toggle is used to equip an equippable item, or unequip it is already equipped
func (c *[N]Item) Toggle() {}
`
	}

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

func generatedFileName(name string) string {
	return name + generatedGoFileSuffix
}

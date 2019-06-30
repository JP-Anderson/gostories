package items

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	io2 "gostories/engine/io"
)

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

func WriteItemsFile() error {
	b := &bytes.Buffer{}
	addAutogeneratedWarning(b)
	addPackagesInfo(b, PckgItem, false)
	addItemStruct("collar", "A small red cat collar with a bell.", false, b)
	addItemStruct("shrubbery", "A small but rather well cared for shrubbery.", true, b)
	addItemStruct("sardines", "A tin of tasty sardines preserved in olive oil.", true, b)
	return GenGoLangFile(PckgItem, b)
}

func WriteFeaturesFile() error {
	b := &bytes.Buffer{}
	addAutogeneratedWarning(b)
	addPackagesInfo(b, PckgFeatures, true)
	addFeatureStruct(
		"shelf",
		"The shelf seems to contain a few old magazines and a cat collar.",
		"Item_Collar",
		b,
	)
	addFeatureStruct(
		"fridge",
		"The fridge is empty apart from a tin of sardines.",
		"Item_Sardines",
		b,
	)
	return GenGoLangFile(PckgFeatures, b)
}

func GenGoLangFile(packageName string, buffer *bytes.Buffer) (err error) {
	filename := "./" + packageName + "/" + generatedFileName(packageName)
	io2.NewLinef("Filename: %v", filename)
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

func addItemStruct(itemName, lookText string, visible bool, buffer *bytes.Buffer) {
	upperCaseItemName := strings.ToUpper(itemName[0:1]) + itemName[1:]
	addBaseStructAndMethods(itemName, upperCaseItemName, TypeItem, buffer)
	addMethodsForItems(upperCaseItemName, buffer)
	addItemConstructor(itemName, upperCaseItemName, lookText, TypeItem, visible, buffer)
}

func addFeatureStruct(ftrName, lookText, containedItem string, buffer *bytes.Buffer) {
	upperCaseFtrName := strings.ToUpper(ftrName[0:1]) + ftrName[1:]
	addBaseStructAndMethods(ftrName, upperCaseFtrName, TypeFeature, buffer)
	addFeatureConstructor(ftrName, upperCaseFtrName, lookText, TypeFeature, containedItem, buffer)
}

func addBaseStructAndMethods(name, upperName, thingType string, buffer *bytes.Buffer) {
	addStruct(name, upperName, thingType, buffer)
	addBasicMethodsForAThing(upperName, thingType, buffer)
}

func addStruct(lowerName, upperName, thingType string, writer *bytes.Buffer) {
	fileStr := `
var [TT]_[N] = New[N][TT]()
var [lc_N]_[TT] *[N][TT]

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
func New[N][TT]() *[N][TT] {
	if [lc_N]_[TT] == nil {
		[lc_N]_[TT] = &[N][TT]{}
		[lc_N]_[TT].Name = "[lc_N]"
		[lc_N]_[TT].LookText = "[LT]"
	}
`
	if visible {
		fileStr = fileStr + `
	[lc_N]_[TT].Show()
`
	}

	fileStr = fileStr +
		`	return [lc_N]_[TT]
}
`

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", name, -1)
	fileStr = strings.Replace(fileStr, "[LT]", lookText, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)

	fsBytes := []byte(fileStr)
	buffer.Write(fsBytes)
}

func addFeatureConstructor(name, upperName, lookText, thingType, itemName string, buffer *bytes.Buffer) {
	fileStr := `
func New[N][TT]() *[N][TT] {
	if [lc_N]_[TT] == nil {
		[lc_N]_[TT] = &[N][TT]{}
		[lc_N]_[TT].Name = "[lc_N]"
		[lc_N]_[TT].LookText = "[LT]"
    	[lc_N]_[TT].Triggers = map[string]things.Trigger{
			"look": things.RevealItemTrigger{
				ItemToReveal: &items.[IN].Thing,
			},
		}
	}
	
	[lc_N]_[TT].Show()
	return [lc_N]_[TT]
}
`

	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", name, -1)
	fileStr = strings.Replace(fileStr, "[LT]", lookText, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)
	fileStr = strings.Replace(fileStr, "[IN]", itemName, -1)

	fsBytes := []byte(fileStr)
	buffer.Write(fsBytes)
}

func addBasicMethodsForAThing(upperName, thingType string, writer *bytes.Buffer) {
	fileStr := `
func (c [N][TT]) GetName() string { return c.Name }

func (c [N][TT]) GetLookText() string { return c.LookText }

func (c [N][TT]) Toggle() {}

func (c *[N][TT]) Show() { c.Thing.Visible = true }

func (c *[N][TT]) Hide() { c.Thing.Visible = false }

func (c [N][TT]) GetThing() things.Thing { return c.Thing }

`
	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)
	fileStr = strings.Replace(fileStr, "[TT]", thingType, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

func addMethodsForItems(upperName string, writer *bytes.Buffer) {
	fileStr := `
func (c [N]Item) Take() {}
`
	fileStr = strings.Replace(fileStr, "[N]", upperName, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

func generatedFileName(name string) string {
	return name + generatedGoFileSuffix
}
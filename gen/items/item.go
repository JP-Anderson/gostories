package items

import (
	"bytes"
	io2 "gostories/engine/io"
	"io/ioutil"

	"strings"
)

func WriteItemsFile() {
	b := &bytes.Buffer{}
	b.WriteString("package items\n\nimport \"gostories/things\"\n")
	AddItem("collar", "A small red cat collar with a bell.", b)
	AddItem("shrubbery", "A small but rather well cared for shrubbery.", b)

	err := ioutil.WriteFile("items.go", b.Bytes(), 0755)
	if err != nil {
		io2.NewLinef("Unable to write file: %v", err)
	}
}

func AddItem(itemName, lookText string, writer *bytes.Buffer) {
	upperCaseItemName := strings.ToUpper(itemName[0:1])+ itemName[1:]
	fileStr := `
var Item_[N] = New[N]Item()

type [N]Item struct {
	things.Thing
}

func New[N]Item() [N]Item {
	c := [N]Item{}
	c.Name = "[lc_N]"
	c.LookText = "[LT]"
	return c
}

func (c [N]Item) GetName() string { return c.Name }

func (c [N]Item) GetLookText() string { return c.LookText }

func (c [N]Item) Take() {}

func (c [N]Item) GetThing() things.Thing { return c.Thing }

func (c [N]Item) Toggle() {}

func (c [N]Item) Show() {
	c.Visible = true
}

func (c [N]Item) Hide() {
	c.Visible = false
}
`
	fileStr = strings.Replace(fileStr, "[N]", upperCaseItemName, -1)
	fileStr = strings.Replace(fileStr, "[lc_N]", itemName, -1)
	fileStr = strings.Replace(fileStr, "[LT]", lookText, -1)

	fsBytes := []byte(fileStr)
	writer.Write(fsBytes)
}

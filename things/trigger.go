package things

import "gostories/engine/io"

type Trigger interface {
	TriggerContextItem()
}

type Triggers map[string]Trigger

type RevealItemTrigger struct {
	ItemToReveal Thing
}

func (r RevealItemTrigger) TriggerContextItem() {
	if !r.ItemToReveal.Visible {
		r.ItemToReveal.Show()
		io.NewLinef("Revealed %v", r.ItemToReveal.Name)
	}
}

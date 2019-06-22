package things

type Trigger interface {
	TriggerContextItem()
}

type Triggers map[string]Trigger

type RevealItemTrigger struct {
	ItemToReveal Thing
}

func (r RevealItemTrigger) TriggerContextItem() {
	r.ItemToReveal.Show()
}

package things

type Trigger interface {
	TriggerContextItem()
}

type Triggers map[string]Trigger

type RevealItemTrigger struct {
	itemToReveal Thing
}
func (r RevealItemTrigger) TriggerContextItem() {
	r.itemToReveal.Show()
}

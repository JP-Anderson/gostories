
// Do not edit! Autogenerated file //

package features

import "gostories/things"

// FeatureShelf probably should remove this and only access structs through the map
var FeatureShelf = NewShelfFeature()
var shelfFeature *ShelfFeature

// ShelfFeature struct
type ShelfFeature struct {
	things.Thing
}

// GetName returns the name of the thing
func (c *ShelfFeature) GetName() string { return c.Name }

// GetLookText returns the description when the player looks at the thing
func (c *ShelfFeature) GetLookText() string { return c.LookText }

// Show makes the thing visible to the player
func (c *ShelfFeature) Show() { c.Thing.Visible = true }

// Hide makes the thing visible to the player
func (c *ShelfFeature) Hide() { c.Thing.Visible = false }

// GetThing returns the underlying Thing struct (need to review if this is used)
func (c *ShelfFeature) GetThing() *things.Thing { return &c.Thing }


// NewShelfFeature creates a new ShelfFeature. Probably will unexport this soon.
func NewShelfFeature() *ShelfFeature {
	if shelfFeature == nil {
		shelfFeature = &ShelfFeature{}
		shelfFeature.Name = "shelf"
		shelfFeature.LookText = "The shelf seems to contain a few old magazines and a cat collar."
    	shelfFeature.Triggers = map[string]string {
			"look": "reveal-item(collar)",
		}
	}
	
	shelfFeature.Show()
	return shelfFeature
}

// FeatureFridge probably should remove this and only access structs through the map
var FeatureFridge = NewFridgeFeature()
var fridgeFeature *FridgeFeature

// FridgeFeature struct
type FridgeFeature struct {
	things.Thing
}

// GetName returns the name of the thing
func (c *FridgeFeature) GetName() string { return c.Name }

// GetLookText returns the description when the player looks at the thing
func (c *FridgeFeature) GetLookText() string { return c.LookText }

// Show makes the thing visible to the player
func (c *FridgeFeature) Show() { c.Thing.Visible = true }

// Hide makes the thing visible to the player
func (c *FridgeFeature) Hide() { c.Thing.Visible = false }

// GetThing returns the underlying Thing struct (need to review if this is used)
func (c *FridgeFeature) GetThing() *things.Thing { return &c.Thing }


// NewFridgeFeature creates a new FridgeFeature. Probably will unexport this soon.
func NewFridgeFeature() *FridgeFeature {
	if fridgeFeature == nil {
		fridgeFeature = &FridgeFeature{}
		fridgeFeature.Name = "fridge"
		fridgeFeature.LookText = "The fridge is empty apart from a tin of sardines."
    	fridgeFeature.Triggers = map[string]string {
			"look": "reveal-item(sardines)",
		}
	}
	
	fridgeFeature.Show()
	return fridgeFeature
}

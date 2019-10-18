package items

import "testing"

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Lowercase item test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// err := WriteItemsFile()
			// if err != nil {
			// 	t.Log(err)
			// }
			err := WriteFeaturesFile()
			if err != nil {
				t.Log(err)
			}
		})
	}
}

package someRandomPackage

import "testing"

func TestTotalCostOk(t *testing.T) {
	shopList := []Shopping{
		{"Kruh", 8, "Prehrana", 2},
		{"Pecivo", 4, "Prehrana", 1},
		{"Mlijeko", 15, "Prehrana", 1},
	}

	total, err := TotalCost(shopList)

	if total != 35 || err != nil {
		t.Errorf("Ne valja")
	}
}

func TestTotalCostTooExpensive(t *testing.T) {
	shopList := []Shopping{
		{"Kruh", 8, "Prehrana", 2},
		{"Pecivo", 4, "Prehrana", 1},
		{"Mlijeko", 15, "Prehrana", 3},
	}

	_, err := TotalCost(shopList)

	if err == nil {
		t.Errorf("Ne valja")
	}
}

func TestTotalCostWrongItem(t *testing.T) {
	shopList := []Shopping{
		{"Kruh", 8, "Prehrana", 2},
		{"Gitara", 4, "Instrument", 1},
		{"Mlijeko", 15, "Prehrana", 1},
	}

	_, err := TotalCost(shopList)

	if err == nil {
		t.Errorf("Ne valja")
	}
}

package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "egor",
		Price: 1.00,
		SKU:   "abs-abs-abss",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

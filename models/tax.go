package models

type Tax struct {
	TaxCode    int     `json:"tax_code"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Type       string  `gorm:"-" json:"type,omitempty"`
	Refundable string  `gorm:"-" json:"refundable,omitempty"`
	Tax        float64 `gorm:"-" json:"tax,omitempty"`
	Amount     float64 `gorm:"-" json:"amount,omitempty"`
}

func (Tax) TableName() string {
	return "tax"
}

func (tax *Tax) GenerateBill() {
	switch tax.TaxCode {
	case 1:
		tax.Type = "Food & Beverage"
		tax.Refundable = "Yes"
		tax.Tax = tax.Price * 0.1
		break
	case 2:
		tax.Type = "Tobacco"
		tax.Refundable = "No"
		tax.Tax = (tax.Price * 0.02) + 10
		break
	case 3:
		tax.Type = "Entertainment"
		tax.Refundable = "No"
		if tax.Price > 0 && tax.Price < 100 {
			tax.Tax = 0
		} else {
			tax.Tax = 0.01 * (tax.Price - 100)
		}
		break
	}
	tax.Amount = tax.Price + tax.Tax
}

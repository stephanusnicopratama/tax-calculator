package services

import (
	"tax-calculator/core"
	"tax-calculator/models"
)

type TaxServices struct {}

var taxes []*models.Tax

func (taxServices *TaxServices) FindAll() ([]models.ITax) {
	core.Globals.DB.Find(&taxes)
	taxResp := make([]models.ITax, 0)
	for _, tax := range taxes {
		tax.GenerateBill()
		taxResp = append(taxResp, tax)
	}
	return taxResp
}

func (taxServices *TaxServices) Load(id string) models.ITax {
	core.Globals.DB.Where("tax_code = ?", id).First(&taxes)
	return taxes[0]
}

func (taxServices *TaxServices) Create(payload models.ITax) bool {
	transformPayload := payload.(*models.Tax)
	if transformPayload.TaxCode == 0 || transformPayload.Name == "" || transformPayload.Price == 0 {
		return false
	}
	core.Globals.DB.NewRecord(transformPayload)
	res := core.Globals.DB.Create(&transformPayload)
	if res.Error == nil {
		return true
	}
	return false
}
package bmecat

import "encoding/xml"

type BMEcat2005 struct {
	XMLName xml.Name `xml:"BMECAT"`
	Version string   `xml:"version,attr"`

	Header         Header           `xml:"HEADER"`
	NewCatalog     TNewCatalog      `xml:"T_NEW_CATALOG"`
	UpdateProducts *TUpdateProducts `xml:"T_UPDATE_PRODUCTS"`
	UpdatePrice    *TUpdatePrice    `xml:"T_UPDATE_PRICE"`
}

func NewBMEcat2005(header Header, catalog TNewCatalog, updateProduct *TUpdateProducts, updatePrice *TUpdatePrice) *BMEcat2005 {
	return &BMEcat2005{
		Header:         header,
		NewCatalog:     catalog,
		UpdateProducts: updateProduct,
		UpdatePrice:    updatePrice,
	}
}

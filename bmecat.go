package bmecat

import "encoding/xml"

const (
	XMLNS    string = "http://www.bmecat.org/bmecat/2005/bmecat_new_catalog"
	XMLNSXSI        = "http://www.w3.org/2001/XMLSchema-instance"
	XSI             = "http://www.bmecat.org/bmecat/2005 bmecat_2005.xsd"
)

type BMEcat2005 struct {
	XMLName  xml.Name `xml:"BMECAT"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLNSXSI string   `xml:"xmlns:xsi,attr"`
	XSI      string   `xml:"xsi,attr"`

	Version string `xml:"version,attr"`

	Header         Header           `xml:"HEADER"`
	NewCatalog     TNewCatalog      `xml:"T_NEW_CATALOG"`
	UpdateProducts *TUpdateProducts `xml:"T_UPDATE_PRODUCTS"`
	UpdatePrice    *TUpdatePrice    `xml:"T_UPDATE_PRICE"`
}

func NewBMEcat2005(header Header, catalog TNewCatalog, updateProduct *TUpdateProducts, updatePrice *TUpdatePrice, xmlns string, xmlnsxsi string, xsi string, version string) *BMEcat2005 {
	return &BMEcat2005{
		Header:         header,
		XMLNS:          xmlns,
		XMLNSXSI:       xmlnsxsi,
		XSI:            xsi,
		Version:        version,
		NewCatalog:     catalog,
		UpdateProducts: updateProduct,
		UpdatePrice:    updatePrice,
	}
}

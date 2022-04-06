package bmecat

type TUpdatePrice struct {
	Formulas Formulas  `xml:"FORMULAS"`
	Product  []Product `xml:"PRODUCT"`
}

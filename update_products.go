package bmecat

type TUpdateProducts struct {
	Formulas Formulas  `xml:"FORMULAS"`
	Product  []Product `xml:"PRODUCT"`
}

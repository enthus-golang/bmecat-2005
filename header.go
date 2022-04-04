package johannesbadbmecat

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"HEADER"`

	GeneratorInformation string  `xml:"GENERATOR_INFO"`
	SupplierID           string  `xml:"SUPPLIER_IDREF"`
	DocumentCreatorID    string  `xml:"DOCUMENT_CREATOR_IDREF"`
	Parties              []Party `xml:"PARTIES"`
	CatalogInformation   Catalog `xml:"CATALOG"`
}

type Catalog struct {
	XMLName xml.Name `xml:"CATALOG"`

	Language string `xml:"LANGUAGE" validate:"omitempty,len=2"`
	ID       string `xml:"CATALOG_ID" validate:"max=20"`
	Version  string `xml:"CATALOG_VERSION" validate:"max=7"`
	Name     string `xml:"CATALOG_NAME,omitempty" validate:"max=100"`
}

type Party struct {
	XMLName xml.Name `xml:"PARTY"`

	ID        int    `xml:"PARTY_ID"`
	PartyRole string `xml:"PARTY_ROLE"`
}

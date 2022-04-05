package bmecat

import "encoding/xml"

const (
	XMLNS    string = "http://www.bmecat.org/bmecat/2005/bmecat_new_catalog"
	XMLNSXSI        = "http://www.w3.org/2001/XMLSchema-instance"
	XSI             = "http://www.bmecat.org/bmecat/2005 bmecat_2005.xsd"
)

type Header struct {
	XMLName  xml.Name `xml:"HEADER"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLNSXSI string   `xml:"xmlns:xsi,attr"`
	XSI      string   `xml:"xsi,attr"`

	Version              string  `xml:"version,attr"`
	GeneratorInformation string  `xml:"GENERATOR_INFO"`
	CatalogInformation   Catalog `xml:"CATALOG"`

	BuyerID TypeID `xml:"BUYER_IDREF"`

	SkeletonAgreements    []Agreement   `xml:"AGREEMENT"`
	LegalInfo             []LegalInfo   `xml:"LEGAL_INFO"`
	SupplierID            SupplierIDRef `xml:"SUPPLIER_IDREF"`
	DocumentCreatorID     TypeID        `xml:"DOCUMENT_CREATOR_IDREF"`
	Parties               Parties       `xml:"PARTIES"`
	Areas                 Areas         `xml:"AREAS"`
	UserDefinedExtensions string        `xml:"USER_DEFINED_EXTENSIONS"`
}

type Catalog struct {
	XMLName xml.Name `xml:"CATALOG"`

	Language        string           `xml:"LANGUAGE" validate:"omitempty,len=2"`
	Locale          []string         `xml:"LOCALE"`
	ID              string           `xml:"CATALOG_ID" validate:"max=20"`
	Version         string           `xml:"CATALOG_VERSION" validate:"max=7"`
	Name            string           `xml:"CATALOG_NAME,omitempty" validate:"max=100"`
	GenerationDate  *Datetime        `xml:"GENERATION_DATE"`
	Territory       []string         `xml:"TERRITORY"`
	AreaRefs        *AreaRefs        `xml:"AREA_REFS"`
	Currency        *Currency        `xml:"CURRENCY"`
	MimeRoot        *string          `xml:"MIME_	ROOT"`
	PriceFlag       []TypeID         `xml:"PRICE_FLAG"`
	PriceFactor     *float64         `xml:"PRICE_FACTOR"`
	StartDate       *Datetime        `xml:"VALID_START_DATE"`
	EndDate         *Datetime        `xml:"VALID_END_DATE"`
	ProductType     *string          `xml:"PRODUCT_TYPE"`
	CountryOfOrigin *string          `xml:"COUNTRY_OF_ORIGIN"`
	DeliveryTimes   *[]DeliveryTimes `xml:"DELIVERY_TIMES"`
	Transport       *Transport       `xml:"TRANSPORT"`
	SupplierIDRef   *SupplierIDRef   `xml:"SUPPLIER_IDREF"`
}
type Buyer struct {
	XMLName xml.Name `xml:"BUYER"`

	ID      *TypeID `xml:"BUYER_ID,omitempty"`
	Name    string  `xml:"BUYER_NAME" validate:"required,max=50"`
	Address Address `xml:"ADDRESS"`
}

type Address struct {
	XMLName     xml.Name `xml:"ADDRESS"`
	AddressType string   `xml:"type,attr"`

	Name           string           `xml:"NAME"`
	Name2          string           `xml:"NAME2"`
	Name3          string           `xml:"NAME3"`
	Department     string           `xml:"DEPARTMENT"`
	ContactDetails []ContactDetails `xml:"CONTACT_DETAILS"`
	Street         string           `xml:"STREET"`
	Zip            string           `xml:"ZIP"`
	BoxNo          string           `xml:"BOXNO"`
	ZipBox         string           `xml:"ZIPBOX"`
	City           string           `xml:"CITY"`
	State          string           `xml:"STATE"`
	Country        string           `xml:"COUNTRY"`
	CountryCoded   string           `xml:"COUNTRY_CODED"`
	Phone          TypeID           `xml:"PHONE"`
	Fax            TypeID           `xml:"FAX"`
	EMail          string           `xml:"EMAIL"`
	PublicKey      TypeID           `xml:"PUBLICKEY"`
	URL            string           `xml:"URL"`
	AddressRemarks string           `xml:"ADDRESS_REMARKS"`
}

type ContactDetails struct {
	ID                 int           `xml:"CONTACT_ID"`
	ContactName        string        `xml:"CONTACT_NAME" validate:"min=1,max=50"`
	FirstName          string        `xml:"FIRST_NAME" validate:"min=1,max=50"`
	Title              string        `xml:"TITLE" validate:"min=1,max=20"`
	AcademicTitle      string        `xml:"ACADEMIC_TITLE" validate:"min=1,max=50"`
	ContactRole        []ContactRole `xml:"CONTACT_ROLE"`
	ContactDescription string        `xml:"CONTACT_DESCR" validate:"min=1,max=250"`
	Phone              Phone         `xml:"PHONE" validate:"min=1,max=50"`
	Fax                Fax           `xml:"FAX" validate:"min=1,max=50"`
	URL                string        `xml:"URL" validate:"min=1,max=255"`
	Emails             []string      `xml:"EMAILS"`
}

type ContactRole struct {
	XMLName xml.Name `xml:"CONTACT_ROLE"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type Phone struct {
	XMLName xml.Name `xml:"PHONE"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}
type Fax struct {
	XMLName xml.Name `xml:"FAX"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type TypeID struct {
	Type  *string `xml:"type,attr" validate:"required"`
	Value string  `xml:",innerxml" validate:"required"`
}

type Agreement struct {
	XMLName          xml.Name `xml:"AGREEMENT"`
	AgreementType    string   `xml:"type,attr"`
	AgreementDefault string   `xml:"default,attr"`

	ID                   string     `xml:"AGREEMENT_ID" validate:"required"`
	LineID               string     `xml:"AGREEMENT_LINE_ID"`
	StartDate            Datetime   `xml:"AGREEMENT_START_DATE"`
	EndDate              Datetime   `xml:"AGREEMENT_END_DATE"`
	AgreementDescription string     `xml:"AGREEMENT_DESC"`
	MIMEInfo             []MIMEInfo `xml:"MIME"`
}

type MIMEInfo struct {
	XMLName xml.Name `xml:"MIME_INFO"`

	MIME []MIME `xml:"MIME"`
}

type LegalInfo struct {
	AreaLegalInfo []AreaLegalInfo `xml:"AREA_LEGAL_INFO"`
}

type AreaLegalInfo struct {
	Territory []string   `xml:"TERRITORY"`
	AreaRefs  AreaRefs   `xml:"AREA_REFS"`
	LegalText string     `xml:"LEGAL_TEXT"`
	MIMEInfo  []MIMEInfo `xml:"MIME_INFO"`
}

type MIME struct {
	XMLName xml.Name `xml:"MIME"`

	Type            string `xml:"MIME_TYPE,omitempty" validate:"max=20"`
	Source          string `xml:"MIME_SOURCE" validate:"max=250"`
	Description     string `xml:"MIME_DESCR,omitempty" validate:"max=250"`
	AlternativeText string `xml:"MIME_ALT,omitempty" validate:"max=50"`
	Purpose         string `xml:"MIME_PURPOSE,omitempty" validate:"max=20"`
	Order           int    `xml:"MIME_ORDER,omitempty"`
}

type SupplierIDRef struct {
	XMLName xml.Name `xml:"SUPPLIER_IDREF"`
	PartyID
}

type Supplier struct {
	XMLName xml.Name `xml:"SUPPLIER"`

	ID      *TypeID `xml:"SUPPLIER_ID,omitempty"`
	Name    string  `xml:"SUPPLIER_NAME" validate:"required,max=50"`
	Address *TypeID `xml:"ADDRESS"`
	MIMEInfo
}

type Parties struct {
	XMLName xml.Name `xml:"PARTIES"`
	Party   []Party  `xml:"PARTY"`
}

type Party struct {
	XMLName xml.Name `xml:"PARTY"`

	ID        TypeID    `xml:"PARTY_ID"`
	PartyRole []string  `xml:"PARTY_ROLE"`
	Address   *Address  `xml:"ADDRESS"`
	MimeInfo  *MIMEInfo `xml:"MIME_INFO"`
}

type Areas struct {
	Area []Area `xml:"AREA"`
}

type Area struct {
	ID          string      `xml:"AREA_ID"`
	Name        string      `xml:"AREA_NAME"`
	Descr       string      `xml:"AREA_DESCR"`
	Territories Territories `xml:"TERRITORIES"`
}

type Territories struct {
	Territory []string `xml:"TERRITORY"`
}

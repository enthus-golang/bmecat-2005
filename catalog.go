package bmecat

import (
	"encoding/xml"
	"time"
)

type TNewCatalog struct {
	XMLName xml.Name `xml:"T_NEW_CATALOG"`

	ClassificationSystem ClassificationSystem `xml:"CLASSIFICATION_SYSTEM"`
	Formulas             *Formulas            `xml:"FORMULAS"`
	IPPDefinitions       *IPPDefinitions      `xml:"IPP_DEFINITIONS"`
	Product              []Product            `xml:"PRODUCT"`
}

type ClassificationSystem struct {
	XMLName xml.Name `xml:"CLASSIFICATION_SYSTEM"`

	Name           string          `xml:"CLASSIFICATION_SYSTEM_NAME"`
	Groups         Groups          `xml:"CLASSIFICATION_GROUPS"`
	FullName       *string         `xml:"CLASSIFICATION_SYSTEM_FULLNAME"`
	VersionDetails *VersionDetails `xml:"CLASSIFICATION_SYSTEM_VERSION_DETAILS"`
	Description    *string         `xml:"CLASSIFICATION_SYSTEM_DESCR"`
	PartyID        *TypeID         `xml:"CLASSIFICATION_SYSTEM_PARTY_IDREF"`
	Level          *int            `xml:"CLASSIFICATION_SYSTEM_LEVELS"`
	LevelNames     *LevelNames     `xml:"CLASSIFICATION_SYSTEM_LEVEL_NAMES"`
	SystemType     *SystemType     `xml:"CLASSIFICATION_SYSTEM_TYPE"`
	AllowedValues  *AllowedValues  `xml:"ALLOWED_VALUES"`
}
type Groups struct {
	XMLName xml.Name `xml:"CLASSIFICATION_GROUPS"`
	Group   []Group  `xml:"CLASSIFICATION_GROUP"`
}
type Group struct {
	XMLName xml.Name `xml:"CLASSIFICATION_GROUP"`

	ID       int    `xml:"CLASSIFICATION_GROUP_ID"`
	Name     string `xml:"CLASSIFICATION_GROUP_NAME"`
	ParentID int    `xml:"CLASSIFICATION_GROUP_PARENT_ID"`
}

type VersionDetails struct {
	Version      string   `xml:"VERSION"`
	VersionDate  Datetime `xml:"VERSION_DATE"`
	Revision     string   `xml:"REVISION"`
	RevisionDate Datetime `xml:"REVISION_DATE"`
	OriginalDate Datetime `xml:"ORIGINAL_DATE"`
}

type LevelNames struct {
	XMLName   xml.Name  `xml:"CLASSIFICATION_SYSTEM_LEVEL_NAMES"`
	LevelName LevelName `xml:"CLASSIFICATION_SYSTEM_LEVEL_NAME"`
}

type LevelName struct {
	XMLName xml.Name `xml:"CLASSIFICATION_SYSTEM_LEVEL_NAME"`
	Name    string   `xml:"CLASSIFICATION_SYSTEM_LEVEL_NAME"`
	LevelID string   `xml:"attr,level"`
}

type SystemType struct {
	GroupIDHierarchy bool   `xml:"GROUPID_HIERARCHY"`
	MappingType      string `xml:"MAPPING_TYPE"`
	MappingLevel     string `xml:"MAPPING_LEVEL"`
	BalancedTree     bool   `xml:"BALANCEDTREE"`
	Inheritance      bool   `xml:"INHERITANCE"`
}

type AllowedValues struct {
	XMLName      xml.Name       `xml:"ALLOWED_VALUES"`
	AllowedValue []AllowedValue `xml:"AllowedValue"`
}

type AllowedValue struct {
	XMLName        xml.Name        `xml:"ALLOWED_VALUE"`
	ID             string          `xml:"ALLOWED_VALUE_ID"`
	Name           string          `xml:"ALLOWED_VALUE_NAME"`
	VersionDetails VersionDetails  `xml:"ALLOWED_VALUE_VERSION"`
	Shortname      string          `xml:"ALLOWED_VALUE_SHORTNAME"`
	Description    string          `xml:"ALLOWED_VALUE_DESCR"`
	Synonyms       AllowedSynonyms `xml:"ALLOWED_VALUE_SYNONYMS"`
	Source         Source          `xml:"ALLOWED_VALUE_SOURCE"`
}

type AllowedSynonyms struct {
	XMLName xml.Name `xml:"ALLOWED_VALUE_SYNONYMS"`

	Synonym []string `xml:"SYNONYM"`
}

type Source struct {
	XMLName xml.Name `xml:"ALLOWED_VALUE_SOURCE"`
	Name    string   `xml:"SOURCE_NAME"`
	URI     string   `xml:"SOURCE_URI"`
	PartyID TypeID   `xml:"PARTY_IDREF"`
}

type FormulaSource struct {
	XMLName xml.Name `xml:"FORMULA_SOURCE"`
	Name    string   `xml:"SOURCE_NAME"`
	URI     string   `xml:"SOURCE_URI"`
	PartyID TypeID   `xml:"PARTY_IDREF"`
}

type Formulas struct {
	XMLName              xml.Name             `xml:"FORMULAS"`
	FormulaID            string               `xml:"FORMULA_ID"`
	Version              VersionDetails       `xml:"FORMULA_VERSION"`
	Name                 string               `xml:"FORMULA_NAME"`
	Description          string               `xml:"FORMULA_DESCR"`
	Source               FormulaSource        `xml:"FORMULA_SOURCE"`
	MIMEInfo             MIMEInfo             `xml:"MIME_INFO"`
	FormulaFunction      FormulaFunction      `xml:"FORMULA_FUNCTION"`
	ParameterDefinitions ParameterDefinitions `xml:"PARAMETER_DEFINITIONS"`
}

type FormulaFunction struct {
	XMLName xml.Name `xml:"FORMULA_FUNCTION"`
	Term    []Term   `xml:"TERM"`
	Type    string   `xml:"attr,type"`
}

type Term struct {
	XMLName xml.Name `xml:"TERM"`

	ID         string `xml:"TERM_ID"`
	Condition  string `xml:"TERM_CONDITION"`
	Expression string `xml:"TERM_EXPRESSION"`
}

type ParameterDefinitions struct {
	XMLName             xml.Name              `xml:"PARAMETER_DEFINITIONS"`
	ParameterDefinition []ParameterDefinition `xml:"PARAMETER_DEFINITION"`
}
type ParameterDefinition struct {
	XMLName xml.Name `xml:"PARAMETER_DEFINITION"`
	Symbol  string   `xml:"PARAMETER_SYMBOL"`
	Basics  Basics   `xml:"PARAMETER_BASICS"`
	FREF    FREF     `xml:"FREF"`
	Value   string   `xml:"PARAMETER_DEFAULT_VALUE"`
	Meaning string   `xml:"PARAMETER MEANING"`
	Order   int      `xml:"PARAMETER_ORDER"`
}

type Basics struct {
	Name        string `xml:"PARAMETER_NAME"`
	Description string `xml:"PARAMETER_DECR"`
	Unit        string `xml:"PARAMETER_UNIT"`
}

type FREF struct {
	FeatureSystemName string `xml:"REFERENCE_FEATURE_SYSTEM_NAME"`
	IDRef             string `xml:"FT_IDREF"`
}

type IPPDefinitions struct {
	ID           string         `xml:"IPP_ID"`
	Type         string         `xml:"IPP_TYPE"`
	IDRef        TypeID         `xml:"IPP_OPERATOR_IDREF"`
	Description  string         `xml:"IPP_DESCR"`
	IPPOperation []IPPOperation `xml:"IPP_OPERATION"`
}

type IPPOperation struct {
	ID          string     `xml:"IPP_OPERATION_ID"`
	Type        string     `xml:"IPP_OPERATION_TYPE"`
	Description string     `xml:"IPP_OPERATION_DESCR"`
	Outbound    []Outbound `xml:"IPP_OUTBOUND"`
	Inbound     []Inbound  `xml:"IPP_INBOUND"`
}

type Outbound struct {
	Format string         `xml:"IPP_OUTBOUND_FORMAT"`
	Params OutboundParams `xml:"IPP_OUTBOUND_PARAMS"`
	URI    string         `xml:"IPP_URI"`
}
type OutboundParams struct {
	Languages          Occurence   `xml:"IPP_LANGUAGES"`
	Territories        Occurence   `xml:"IPP_TERRITORIES"`
	PriceCurrencies    Occurence   `xml:"IPP_PRICE_CURRENCY"`
	PriceTypes         Occurence   `xml:"IPP_PRICE_TYPES"`
	PID                Occurence   `xml:"IPP_SUPPLIER_PID"`
	ProductConfigIDRef Occurence   `xml:"IPP PRODUCTCONFIG_IDREF"`
	Info               Occurence   `xml:"IPP_USER_INFO"`
	Definition         []Occurence `xml:"IPP AUTHENTIFICATION_INFO"`
}

type Occurence struct {
	Occurence string `xml:"occurence,attr" validate:"required"`
	Value     string `xml:",innerxml" validate:"required"`
}

type Inbound struct {
	Format        string        `xml:"IPP_INBOUND_FORMAT"`
	InboundParams InboundParams `xml:"IPP_INBOUND_PARAMS"`
	Time          time.Duration `xml:"IPP_RESPONSE_TIME"`
}

type InboundParams struct {
	Definition []Occurence `xml:"IPP_PARAM_DEFINITION"`
}

type Product struct {
	XMLName xml.Name `xml:"PRODUCT"`
	Mode    string   `xml:"mode,attr"`

	SupplierPID            TypeID                  `xml:"SUPPLIER_PID" validate:"required,max=32"`
	SupplierIDRef          *SupplierIDRef          `xml:"SUPPLIER_IDREF"`
	ProductDetails         ProductDetails          `xml:"PRODUCT_DETAILS"`
	ProductFeatures        []ProductFeature        `xml:"PRODUCT_FEATURES"`
	ProductOrderDetails    ProductOrderDetails     `xml:"PRODUCT_ORDER_DETAILS"`
	ProductPriceDetails    ProductPriceDetails     `xml:"PRODUCT_PRICE_DETAILS"`
	MimeInfo               *MIMEInfo               `xml:"MIME_INFO"`
	UserDefinedExtensions  *string                 `xml:"USER_DEFINED_EXTENSIONS"`
	ProductReferences      []ProductReference      `xml:"PRODUCT_REFERENCE"`
	ProductContacts        *ProductContacts        `xml:"PRODUCT_CONTACTS"`
	ProductIPPDetails      *ProductIPPDetails      `xml:"PRODUCT_IPP_DETAILS"`
	ProductLogisticDetails *ProductLogisticDetails `xml:"PRODUCT_LOGISTIC_DETAILS"`
	ProductConfigDetails   *ProductConfigDetails   `xml:"PRODUCT_CONFIG_DETAILS"`
}

type ProductDetails struct {
	XMLName xml.Name `xml:"PRODUCT_DETAILS"`

	DescriptionShort             string               `xml:"DESCRIPTION_SHORT" validate:"required,max=80"`
	DescriptionLong              string               `xml:"DESCRIPTION_LONG" validate:"max=64000"`
	InternationalPID             *[]TypeID            `xml:"INTERNATIONAL_PID" validate:"max=100"`
	SupplierAlternativeArticleID string               `xml:"SUPPLIER_ALT_AID" validate:"max=50"`
	BuyerProductID               []BuyerProductID     `xml:"BUYER_PID"`
	ManufacturerProductID        *string              `xml:"MANUFACTURER_PID" validate:"max=50"`
	ManufacturerID               *TypeID              `xml:"MANUFACTURER_IDREF" validate:"max=250"`
	ManufacturerTypeDescription  *string              `xml:"MANUFACTURER_TYPE_DESCR" validate:"max=50"`
	ERPGroupBuyer                *string              `xml:"ERP_GROUP_BUYER" validate:"max=10"`
	ERPGroupSupplier             *string              `xml:"ERP_GROUP_SUPPLIER" validate:"max=10"`
	SpecialTreatmentClass        *string              `xml:"SPECIAL_TREATMENT_CLASS" validate:"max=20"`
	Keyword                      *string              `xml:"KEYWORD,omitempty" validate:"max=50"`
	Remarks                      *TypeID              `xml:"REMARKS,omitempty" validate:"max=64000"`
	Segment                      *string              `xml:"SEGMENT,omitempty" validate:"max=100"`
	ArticleOrder                 *int                 `xml:"ARTICLE_ORDER"`
	ProductStatus                []TypeID             `xml:"PRODUCT_STATUS" validate:"max=250"`
	InternationalRestrictions    []TypeID             `xml:"INTERNATIONAL_RESTRICTIONS"`
	AccountingInfo               *AccountingInfo      `xml:"ACCOUNTING_INFO"`
	AgreementReference           []AgreementReference `xml:"AGREEMENT_REF"`
	ProductType                  []string             `xml:"PRODUCT_TYPE" validate:"max=50"`
	ProductCategory              []string             `xml:"PRODUCT_CATEGORY" validate:"max=20"`
}

type AccountingInfo struct {
	CostCategoryID TypeID `xml:"COST_CATEGORY_ID"`
	CostType       string `xml:"COST_TYPE" validate:"max=64"`
	CostAccount    string `xml:"COST_ACCOUNT" validate:"max=64"`
}

type AgreementReference struct {
	IDReference     string `xml:"AGREEMENT_IDREF" validate:"max=50"`
	LineIDReference string `xml:"AGREEMENT_LINE_IDREF" validate:"max=50"`
}

type ProductFeature struct {
	XMLName xml.Name `xml:"PRODUCT_FEATURES"`

	ReferenceFeatureSystemName string         `xml:"REFERENCE_FEATURE_SYSTEM_NAME,omitempty" validate:"max=50"`
	ReferenceFeatureGroupID    *TypeID        `xml:"REFERENCE_FEATURE_GROUP_ID,omitempty" validate:"max=60"`
	ReferenceFeatureGroupName  string         `xml:"REFERENCE_FEATURE_GROUP_NAME,omitempty" validate:"max=60"`
	ReferenceFeatureGroupID2   *TypeID        `xml:"REFERENCE_FEATURE_GROUP_ID2,omitempty" validate:"max=60"`
	GroupProductOrder          *int           `xml:"GROUP_PRODUCT_ORDER"`
	Features                   []Feature      `xml:"FEATURE"`
	FeatureGroups              []FeatureGroup `xml:"FEATURE_GROUP"`
}

type Feature struct {
	XMLName      xml.Name         `xml:"FEATURE"`
	Name         *string          `xml:"FNAME"`
	FeatureID    *string          `xml:"FT_IDREF" validate:"max=60"`
	Template     *FeatureTemplate `xml:"FTEMPLATE"`
	Value        *string          `xml:"FVALUE"`
	ValueID      *string          `xml:"VALUE_IDREF"`
	Variants     *Variants        `xml:"VARIANTS"`
	Unit         *string          `xml:"FUNIT"`
	Order        *int             `xml:"FORDER"`
	Descriptions *string          `xml:"FDESCR"`
	Details      *string          `xml:"FVALUE_DETAILS"`
	ValueType    *string          `xml:"FVALUE_TYPE"`
	FID          *string          `xml:"FID"`
	ParentID     *string          `xml:"FPARENT_ID"`
}

type FeatureTemplate struct {
	ID             string         `xml:"FT_ID"`
	Name           string         `xml:"FT_NAME"`
	Shortname      string         `xml:"FT_SHORTNAME"`
	Description    string         `xml:"FT_DESCR"`
	Version        VersionDetails `xml:"FT_VERSION"`
	GroupID        string         `xml:"FT_GROUP_IDREF"`
	GroupName      string         `xml:"FT_GROUP_NAME"`
	Dependencies   Dependencies   `xml:"FT_DEPENDENCIES"`
	FeatureContent FeatureContent `xml:"FEATURE_CONTENT"`
}
type Dependencies struct {
	ID string `xml:"FT_IDREF"`
}

type FeatureContent struct {
	DataType  string   `xml:"FT_DATATYPE"`
	Facets    Facets   `xml:"FT_FACETS"`
	Values    Values   `xml:"FT_VALUES"`
	Valency   string   `xml:"FT_VALENCY"`
	UnitID    string   `xml:"FT_UNIT_IDREF"`
	Unit      string   `xml:"FT_UNIT"`
	Mandatory bool     `xml:"FT_MANDATORY"`
	Order     int      `xml:"FT_ORDER"`
	Symbol    string   `xml:"FT_SYMBOL"`
	Synonyms  Synonyms `xml:"FT_SYNONYMS"`
	MIMEInfo  MIMEInfo `xml:"MIME_INFO"`
	Source    FTSource `xml:"FT_SOURCE"`
	Note      string   `xml:"FT_NOTE"`
	Remark    string   `xml:"FT_REMARK"`
}

type FTSource struct {
	XMLName xml.Name `xml:"FT_SOURCE"`
	Name    string   `xml:"SOURCE_NAME"`
	URI     string   `xml:"SOURCE_URI"`
	PartyID TypeID   `xml:"PARTY_IDREF"`
}

type Facets struct {
	Facet []TypeID `xml:"FACET"`
}

type Values struct {
	Value []Value `xml:"VALUE"`
}

type Value struct {
	ID          string     `xml:"VALUE_IDREF"`
	ValueSimple string     `xml:"VALUE_SIMPLE"`
	Text        string     `xml:"VALUE_TEXT"`
	ValueRange  ValueRange `xml:"VALUE_RANGE"`
	MIMEInfo    MIMEInfo   `xml:"MIME_INFO"`
	ConfigInfo  ConfigInfo `xml:"CONFIG_INFO"`
	Order       int        `xml:"VALUE_ORDER"`
	Flag        bool       `xml:"DEFAULT_FLAG"`
}
type ValueRange struct {
	StartValue    Interval `xml:"STARTVALUE"`
	EndValue      Interval `xml:"ENDVALUE"`
	IntervalValue float64  `xml:"INTERVALVALUE"`
}

type Interval struct {
	Type  string `xml:"intervaltype,attr"`
	Value float64
}
type ConfigInfo struct {
	StartDate    Datetime  `xml:"VALID_START_DATE"`
	EndDate      Datetime  `xml:"VALID_END_DATE"`
	DailyPrice   bool      `xml:"DAILY_PRICE"`
	ProductPrice PriceType `xml:"PRODUCT_PRICE"`
}
type PriceType struct {
	ProductPrice float64 `xml:"PRODUCT_PRICE"`
	PriceType    string  `xml:"price_type,attr"`
}
type Synonyms struct {
	Name string `xml:"SOURCE_NAME"`
	URI  string `xml:"SOURCE_URI"`
	ID   TypeID `xml:"PARTY_IDREF"`
}

type Variants struct {
	Variant Variant `xml:"VARIANT"`
	Order   int     `xml:"VORDER"`
}

type Variant struct {
	Value      []string `xml:"FVALUE"`
	ValueID    []string `xml:"VALUE_IDREF"`
	Supplement string   `xml:"SUPPLIER_AID_SUPPLEMENT"`
}

type FeatureGroup struct {
	GroupType               string                    `xml:"featureGroupType,attr"`
	Name                    string                    `xml:"FEATURE_GROUP_NAME"`
	FeatureGroupDescription []FeatureGroupDescription `xml:"FEATURE_GROUP_DESCRIPTION"`
	ReferenceFeatureGroupID TypeID                    `xml:"REFERENCE_FEATURE_GROUP_ID"`
	Feature                 []Feature                 `xml:"FEATURE"`
}

type FeatureGroupDescription struct {
	Value    string
	Language string `xml:"lang,attr"`
	Locale   string `xml:"locale,attr"`
}

type BuyerProductID struct {
	XMLName xml.Name `xml:"BUYER_PID"`
	Type    string   `xml:"type,attr" validate:"required"`
	Value   string   `xml:",innerxml" validate:"required,max=50"`
}

type ProductOrderDetails struct {
	XMLName xml.Name `xml:"PRODUCT_ORDER_DETAILS"`

	OrderUnit                     string         `xml:"ORDER_UNIT" validate:"max=3"`
	ContentUnit                   string         `xml:"CONTENT_UNIT,omitempty" validate:"max=3"`
	NumberContentUnitPerOrderUnit float64        `xml:"NO_CU_PER_OU,omitempty"`
	SupplierPID                   *SupplierPID   `xml:"SUPPLIER_PIDREF"`
	SupplierIDRef                 *SupplierIDRef `xml:"SUPPLIER_IDREF"`
	PriceQuantity                 *float64       `xml:"PRICE_QUANTITY"`
	QuantityMinimum               float64        `xml:"QUANTITY_MIN,omitempty"`
	QuantityInterval              float64        `xml:"QUANTITY_INTERVAL,omitempty"`
	QuantityMaximum               *float64       `xml:"QUANTITY_MAX"`
	PackingUnits                  *PackingUnits  `xml:"PACKING_UNITS"`
}
type PackingUnits struct {
	PackingUnit PackingUnit `xml:"PACKING_UNIT"`
}
type PackingUnit struct {
	QuantityMinimum float64       `xml:"QUANTITY_MIN,omitempty"`
	QuantityMaximum float64       `xml:"QUANTITY_MAX"`
	UnitCode        string        `xml:"PACKING_UNIT_CODE"`
	UnitDescription string        `xml:"PACKING_UNIT_DESCR"`
	SupplierPID     SupplierPID   `xml:"SUPPLIER_PID"`
	SupplierPIDRef  string        `xml:"SUPPLIER_PIDREF"`
	SupplierIDRef   SupplierIDRef `xml:"SUPPLIER_IDREF"`
}

type ProductPriceDetails struct {
	XMLName xml.Name `xml:"PRODUCT_PRICE_DETAILS"`

	StartDate    Datetime     `xml:"VALID_START_DATE"`
	EndDate      Datetime     `xml:"VALID_END_DATE"`
	DailyPrice   bool         `xml:"DAILY_PRICE"`
	ProductPrice ProductPrice `xml:"PRODUCT_PRICE"`
}

type ProductPrice struct {
	XMLName xml.Name `xml:"PRODUCT_PRICE"`
	Type    string   `xml:"price_type,attr"`

	PriceAmount  *float64      `xml:"PRICE_AMOUNT"`
	PriceFormula *PriceFormula `xml:"PRICE_FORMULA"`
	Currency     *string       `xml:"PRICE_CURRENCY"`
	TaxDetails   []TaxDetails  `xml:"TAX_DETAILS"`
	Tax          *float64      `xml:"TAX"`
	Factor       *float64      `xml:"PRICE_FACTOR"`
	LowerBound   *float64      `xml:"LOWER_BOUND"`
	Territory    []string      `xml:"TERRITORY"`
	AreaRefs     *AreaRefs     `xml:"AREA_REFS"`
	PriceBase    *PriceBase    `xml:"PRICE_BASE"`
	PriceFlag    []PriceFlag   `xml:"PRICE_FLAG"`
}

type PriceFormula struct {
	ID         string     `xml:"FORMULA_IDREF"`
	Parameters Parameters `xml:"PARAMETERS"`
}

type TaxDetails struct {
	CalculationSequence int     `xml:"CALCULATION_SEQUENCE"`
	Category            string  `xml:"TAX_CATEGORY"`
	Type                string  `xml:"TAX_TYPE"`
	Tax                 float64 `xml:"TAX"`
	ExemptionReason     string  `xml:"EXEMPTION_REASON"`
	Jurisdiction        string  `xml:"JURISDICTION"`
}

type PriceBase struct {
	PriceUnit       float64 `xml:"PRICE_UNIT"`
	PriceUnitFactor float64 `xml:"PRICE_UNIT_FACTOR"`
	Synonym         string  `xml:"SYNONYM"`
}

type PriceFlag struct {
	Value bool
	Type  string `xml:"type,attr"`
}

type ProductReference struct {
	XMLName  xml.Name `xml:"PRODUCT_REFERENCE"`
	Type     string   `xml:"type,attr"`
	Quantity float64  `xml:"quantity,attr"`

	Value string `xml:"PROD_ID_TO"`
}

type ProductContacts struct {
	PartyID   TypeID `xml:"PARTY_IDREF"`
	ContactID string `xml:"CONTACT_IDREFF"`
}
type ProductIPPDetails struct {
	IPP []IPP `xml:"IPP"`
}

type IPP struct {
	ID           string        `xml:"IPP_IDREF"`
	OperationID  []string      `xml:"IPP_OPERATION_IDREF"`
	ResponseTime time.Duration `xml:"IPP_RESPONSE_TIME"`
	URI          []string      `xml:"IPP_URI"`
	Param        []IPPParam    `xml:"IPP_PARAM"`
}

type IPPParam struct {
	NameRef string `xml:"IPP_PARAM_NAMEREF"`
	Value   string `xml:"IPP_PARAM_VALUE"`
}

type ProductLogisticDetails struct {
	CustomsTariffNo  CustomsTariffNo  `xml:"CUSTOMS_TARIFF_NUMBER"`
	StatisticFactor  float64          `xml:"STATISTICS_FACTOR"`
	CountryOfOrigin  string           `xml:"COUNTRY_OF_ORIGIN"`
	ProductDimension ProductDimension `xml:"PRODUCT_DIMENSIONS"`
	DeliveryTimes    []DeliveryTimes  `xml:"DELIVERY_TIMES"`
	Transport        []Transport      `xml:"TRANSPORT"`
	MeansOfTransport []TypeID         `xml:"MEANS_OF_TRANSPORT"`
}

type CustomsTariffNo struct {
	CustomsNo string    `xml:"CUSTOMS_NUMBER"`
	Territory []Country `xml:"TERRITORY"`
	AreaRefs  AreaRefs  `xml:"AREA_REFS"`
}
type AreaRefs struct {
	AreaID []string `xml:"AREA_IDREF"`
}

type ProductDimension struct {
	Volume float64 `xml:"VOLUME"`
	Weight float64 `xml:"WEIGHT"`
	Length float64 `xml:"LENGTH"`
	Width  float64 `xml:"WIDTH"`
	Depth  float64 `xml:"DEPTH"`
}

type DeliveryTimes struct {
	Territory []Country `xml:"TERRITORY"`
	AreaRefs  AreaRefs  `xml:"AREA_REFS"`
	TimeSpan  TimeSpan  `xml:"Time_SPAN"`
	LeadTime  float64   `xml:"LEADTIME"`
}

type TimeSpan struct {
	TimeBase          string     `xml:"TIME_BASE"`
	TimeValueDuration string     `xml:"TIME_VALUE_DURATION"`
	TimeValueInterval string     `xml:"TIME_VALUE_INTERVAL"`
	TimeValueStart    string     `xml:"TIME_VALUE_START"`
	TimeValueEnd      string     `xml:"TIME_VALUE_END"`
	SubTimeSpans      []TimeSpan `xml:"SUB_TIME_SPANS"`
}

type Transport struct {
	Incoterm        string `xml:"INCOTERM"`
	Location        string `xml:"LOCATION"`
	TransportRemark string `xml:"TRANSPORT_REMARK"`
}

type ProductConfigDetails struct {
	ConfigStep        []ConfigStep      `xml:"CONFIG_STEP"`
	PredefinedConfigs PredefinedConfigs `xml:"PREDEFINED_CONFIGS"`
	ConfigRules       ConfigRules       `xml:"CONFIG_RULES"`
	ConfigFormulas    ConfigFormulas    `xml:"CONFIG_FORMULAS"`
}

type ConfigStep struct {
	ID                  string              `xml:"STEP_ID"`
	Header              string              `xml:"STEP_HEADER"`
	DescrShort          string              `xml:"STEP_DESCR_SHORT"`
	DescrLong           string              `xml:"STEP_DESCR_LONG"`
	Order               int                 `xml:"STEP_ORDER"`
	InteractionType     string              `xml:"STEP_INTERACTION_TYPE"`
	Code                string              `xml:"CONFIG_CODE"`
	ProductPriceDetails ProductPriceDetails `xml:"PRODUCT_PRICE_DETAILS"`
	ConfigFeature       ConfigFeature       `xml:"CONFIG_FEATURE"`
	ConfigParts         ConfigParts         `xml:"CONFIG_PARTS"`
	OccuranceMin        int                 `xml:"MIN_OCCURANCE"`
	OccurenceMax        int                 `xml:"MAX_OCCURANCE"`
}

type ConfigFeature struct {
	FREF      FREF            `xml:"FREF"`
	FTemplate FeatureTemplate `xml:"FTEMPLATE"`
	MIMEInfo  MIMEInfo        `xml:"MIME_INFO"`
}

type ConfigParts struct {
	Alternative   Alternative `xml:"PART_ALTERNATIVE"`
	SelectionType string      `xml:"PART_SELECTION_TYPE"`
}
type Alternative struct {
	SupplierPID         string              `xml:"SUPPLIER_PIDREF"`
	SupplierIDRef       TypeID              `xml:"SUPPLIER_IDREF"`
	ProductOrderDetails ProductOrderDetails `xml:"PRODUCT_ORDER"`
	DefaultFlag         bool                `xml:"DEFAULT_FLAG"`
	ConfigCode          int                 `xml:"CONFIG_CODE"`
	ProductPriceDetails ProductPriceDetails `xml:"PRODUCT_PRICE_DETAILS"`
}

type PredefinedConfigs struct {
	PredefinedConfig []PredefinedConfig `xml:"PREDEFINED_CONFIG"`
	ConfigCoverage   string             `xml:"PREDEFINED_CONFIG_COVERAGE"`
}

type PredefinedConfig struct {
	ConfigCode          string              `xml:"PREDEFINED_CONFIG_CODE"`
	ConfigName          string              `xml:"PREDEFINED_CONFIG_NAME"`
	ConfigDescr         string              `xml:"PREDEFINED_CONFIG_DESCR"`
	ConfigOrder         string              `xml:"PREDEFINED_CONFIG_ORDER"`
	ProductPriceDetails ProductPriceDetails `xml:"PRODUCT_PRICE_DETAILS"`
	SupplierPID         TypeID              `xml:"SUPPLIER_PID"`
	InternationalPID    TypeID              `xml:"INTERNATIONAL_PID"`
}

type ConfigRules struct {
	Term []Term `xml:"TERM"`
}
type ConfigFormulas struct {
	ID         string     `xml:"FORMULA_IDREF"`
	Parameters Parameters `xml:"PARAMETERS"`
}

type Parameters struct {
	Parameter []Parameter `xml:"PARAMETER"`
}
type Parameter struct {
	SymbolRef string `xml:"PARAMETER_SYMBOLREF"`
	Value     string `xml:"PARAMETER_VALUE"`
}

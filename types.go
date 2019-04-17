package bmecat

import (
	"encoding/xml"
)

type Language struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 LANGUAGE"`
	Default bool     `xml:"default,attr,omitempty"`
	Value   string   `xml:",chardata" validate:"oneof=aar abk ace ach ada afa afh afr aka akk alb ale alg amh ang apa ara arc arm arn arp art arw asm ath aus ava ave awa aym aze bad bai bak bal bam ban baq bas bat bej bel bem ben ber bho bih bik bin bis bla bnt bod bos bra bre btk bua bug bul bur cad cai car cat cau ceb cel ces cha chb che chg chi chk chm chn cho chp chr chu chv chy cmc cop cor cos cpe cpf cpp cre crp cus cym cze dak dan day del den deu dgr din div doi dra dua dum dut dyu dzo efi egy eka ell elx eng enm epo est eus ewe ewo fan fao fas fat fij fin fiu fon fra fre frm fro fry ful fur gaa gay gba gem geo ger gez gil gla gle glg glv gmh goh gon gor got grb grc gre grn guj gwi hai hau haw heb her hil him hin hit hmn hmo hrv hun hup hye iba ibo ice ijo iku ile ilo ina inc ind ine ipk ira iro isl ita jav jpn jpr jrb kaa kab kac kal kam kan kar kas kat kau kaw kaz kha khi khm kho kik kin kir kmb kok kom kon kor kos kpe kro kru kua kum kur kut lad lah lam lao lat lav lez lin lit lol loz ltz lua lub lug lui lun luo lus mac mad mag mah mai mak mal man mao map mar mas may mdr men mga mic min mis mkd mkh mlg mlt mnc mni mno moh mol mon mos mri msa mul mun mus mwr mya myn nah nai nau nav nbl nde ndo nds nep new nia nic niu nld nno nob non nor nso nub nya nym nyn nyo nzi oci oji ori orm osa oss ota oto paa pag pal pam pan pap pau peo per phi phn pli pol pon por pra pro pus qaa que raj rap rar roa roh rom ron rum run rus sad sag sah sai sal sam san sas sat scc sco scr sel sem sga sgn shn sid sin sio sit sla slk slo slv sme smi smo sna snd snk sog som son sot spa sqi srd srp srr ssa ssw suk sun sus sux swa swe syr tah tai tam tat tel tem ter tet tgk tgl tha tib tig tir tiv tkl tli tmh tog ton tpi tsi tsn tso tuk tum tur tut tvl twi tyv uga uig ukr umb und urd uzb vai ven vie vol vot wak wal war was wel wen wol xho yao yap yid yor ypk zap zen zha zho znd zul"`
}

type ContactID string

type PartyType string

const (
	PartyTypeBuyerSpecific    PartyType = "buyer_specific"
	PartyTypeCustomerSpecific           = "customer_specific"
	PartyTypeDUNS                       = "duns"
	PartyTypeILN                        = "iln"
	PartyTypeGLN                        = "gln"
	PartyTypeSupplierSpecific           = "supplier_specific"
)

type PartyID struct {
	XMLName xml.Name  `xml:"http://www.bmecat.org/bmecat/2005 PARTY_ID"`
	Type    PartyType `xml:"type,attr" validate:"min=1,max=250"`
	Value   string    `xml:",chardata" validate:"min=1,max=250"`
}

type Name struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Name2 struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 NAME2"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Name3 struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 NAME3"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Department struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 DEPARTMENT"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Street struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 STREET"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Zip struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 ZIP"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type BoxNo struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 BOXNO"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type ZipBox struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 ZIPBOX"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type City struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 CITY"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type State struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 STATE"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Country struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 COUNTRY"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type CountryCoded string

type VatID string

type PhoneType string

const (
	PhoneMobile  PhoneType = "mobile"
	PhoneOffice            = "office"
	PhonePrivate           = "private"
)

type Phone struct {
	MultiLocale
	XMLName xml.Name  `xml:"http://www.bmecat.org/bmecat/2005 PHONE"`
	Type    PhoneType `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string    `xml:",chardata" validate:"min=1,max=50"`
}

type FaxType string

const (
	FaxOffice  FaxType = "office"
	FaxPrivate         = "private"
)

type Fax struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 FAX"`
	Type    FaxType  `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type ContactName struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 CONTACT_NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type FirstName struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 FIRST_NAME"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type Title struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 TITLE"`
	Value   string   `xml:",chardata" validate:"min=1,max=20"`
}

type AcademicTitle struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 ACADEMIC_TITLE"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type ContactDescription struct {
	MultiLocale
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 CONTACT_DESCR"`
	Value   string   `xml:",chardata" validate:"min=1,max=250"`
}

// 1-255
type URL string

type EMails struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 EMAILS"`
	EMail   []EMail  `xml:"http://www.bmecat.org/bmecat/2005 EMAIL" validate:"required,dive,required,max=255"`
}

type EMail string

type Authentification struct {
	XMLName  xml.Name `xml:"http://www.bmecat.org/bmecat/2005 AUTHENTIFICATION"`
	Login    string   `xml:"http://www.bmecat.org/bmecat/2005 LOGIN" validate:"required,max=60"`
	Password string   `xml:"http://www.bmecat.org/bmecat/2005 PASSWORD,omitempty" validate:"max=20"`
}

type BuyerIDRef struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 BUYER_IDREF"`
	PartyID
}

type SupplierIDRef struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_IDREF"`
	PartyID
}

type PIDType string

const (
	PIDBuyerSpecific    PIDType = "buyer_specific"
	PIDEAN                      = "ean"
	PIDGTIN                     = "gtin"
	PIDSupplierSpecific         = "supplier_specific"
	PIDUPC                      = "upc"
)

type SupplierPID struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 SUPPLIER_PID"`
	Type    PIDType  `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=32"`
}

type BuyerPID struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 BUYER_PID"`
	Type    PIDType  `xml:"type,attr,omitempty" validate:"min=1,max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=50"`
}

type InternationalAID struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 INTERNATIONAL_AID"`
	Type    PIDType  `xml:"type,attr,omitempty" validate:"max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=100"`
}

type InternationalPID struct {
	XMLName xml.Name `xml:"http://www.bmecat.org/bmecat/2005 INTERNATIONAL_PID"`
	Type    PIDType  `xml:"type,attr,omitempty" validate:"max=50"`
	Value   string   `xml:",chardata" validate:"min=1,max=100"`
}

type OrderUnit string

const (
	UnitPiece OrderUnit = "C62"
)

type PriceAmount float64

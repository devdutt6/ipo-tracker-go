package static

type CMP int
type STT int

const (
	CAMEO CMP = iota
	BIGSHARE
	MAASHITLA
	LINKINTIME
)

const (
	NOT_APPLIED STT = iota
	NOT_ALLOTED
	ALLOTED
)

var REGISTRAR = map[CMP]string{
	CAMEO:      "CAMEO",
	BIGSHARE:   "BIGSHARE",
	MAASHITLA:  "MAASHITLA",
	LINKINTIME: "LINKINTIME",
}

var SCRAP_URL = map[CMP]string{
	CAMEO:      "https://ipostatus3.cameoindia.com:3000/api/1.0/ipostatus",
	MAASHITLA:  "https://maashitla.com/PublicIssues/Search",
	BIGSHARE:   "https://ipo.bigshareonline.com/Data.aspx/FetchIpodetails",
	LINKINTIME: "https://in.mpms.mufg.com/Initial_Offer/IPO.aspx/SearchOnPan",
}

var ALLOTMENT_STATUS = map[STT]string{
	NOT_APPLIED: "NOT APPLIED",
	NOT_ALLOTED: "NOT ALLOTED",
	ALLOTED:     " ALLOTED SHARES",
}

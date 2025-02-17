package api

import "encoding/xml"

type CameoRequest struct {
	Code  string `json:"code"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewCameoRequest(code string, panNumber string) *CameoRequest {
	return &CameoRequest{Code: code, Type: "pan", Value: panNumber}
}

type BigshareRequest struct {
	CompanyCode   string `json:"Company"`
	Pan           string `json:"PanNo"`
	Type          string `json:"SelectionType"`
	DdlType       string `json:"ddlType"`
	Applicationno string `json:"Applicationno"`
	Txtcsdl       string `json:"txtcsdl"`
	TxtDPID       string `json:"txtDPID"`
	TxtClId       string `json:"txtClId"`
}

func NewBigshareRequest(code string, panNumber string) *BigshareRequest {
	return &BigshareRequest{
		CompanyCode:   code,
		Pan:           panNumber,
		Type:          "PN",
		DdlType:       "0",
		Applicationno: "",
		Txtcsdl:       "",
		TxtDPID:       "",
		TxtClId:       "",
	}
}

type LinkintimeRequest struct {
	Code   string `json:"clientid"`
	Pan    string `json:"PAN"`
	CHKVAL string `json:"CHKVAL"`
	IFSC   string `json:"IFSC"`
	Token  string `json:"token"`
}

func NewLinkintimeRequest(code string, panNumber string) *LinkintimeRequest {
	return &LinkintimeRequest{
		Code:   code,
		Pan:    panNumber,
		CHKVAL: "1",
		IFSC:   "",
		Token:  "",
	}
}

type Table struct {
	XMLName xml.Name `xml:"Table"`
	// Id          int      `xml:"id"`
	// OfferPrice  int      `xml:"offer_price"`
	// Pull        string   `xml:"pull"`
	// Speed       string   `xml:"speed"`
	// Match       string   `xml:"match"`
	// DPCLITID    float64  `xml:"DPCLITID"`
	// RFNDNO      int      `xml:"RFNDNO"`
	// RFNDAMT     int      `xml:"RFNDAMT"`
	// NAME1       string   `xml:"NAME1"`
	// Companyname string   `xml:"companyname"`
	ALLOT  int `xml:"ALLOT"`
	SHARES int `xml:"SHARES"`
	// AMTADJ      int      `xml:"AMTADJ"`
	// PEMNDG      string   `xml:"PEMNDG"`
	// INVCODE     int      `xml:"INVCODE"`
	// BNKCODE     int      `xml:"BNKCODE"`
}

type DataSet struct {
	XMLName xml.Name `xml:"NewDataSet"`
	Table   Table    `xml:"Table"`
}

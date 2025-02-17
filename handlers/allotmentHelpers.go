package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/devdutt6/ipo-tracker-go/api"
	"github.com/devdutt6/ipo-tracker-go/static"
)

func CheckWithCameo(company *api.CompanyDocument, pans *[]api.PanDocument) map[string]string {
	var response = map[string]string{}

	for _, pan := range *pans {
		requestBody := api.NewCameoRequest(company.CompanyCode, pan.PanNumber)
		jsonBody, err := json.Marshal(requestBody)

		if err != nil {
			fmt.Println("failed to parse to json")
			response[pan.PanNumber] = "Failed"
			continue
		}
		resp, err := http.Post(static.SCRAP_URL[static.CAMEO], "application/json", strings.NewReader(string(jsonBody)))
		if err != nil {
			fmt.Println("failed request to cameo")
			response[pan.PanNumber] = "Failed"
			continue
		}
		defer resp.Body.Close()
		stringData, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("failed load response from cameo")
			response[pan.PanNumber] = "Failed"
			continue
		}

		var adata []any

		if err := json.Unmarshal(stringData, &adata); err != nil {
			fmt.Println("failed to parse to json reponse body")
			response[pan.PanNumber] = "Failed"
			continue
		}

		if len(adata) == 0 {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
		} else {
			data := adata[0].(map[string]any)
			refundAmount := data["refundAmount"].(float64)
			allotedShares := data["allotedShares"].(float64)

			if refundAmount > 0 && allotedShares == 0 {
				response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_ALLOTED]
			} else if allotedShares > 0 {
				response[pan.PanNumber] = fmt.Sprintf("%v", allotedShares) + " " + static.ALLOTMENT_STATUS[static.ALLOTED]
			}
		}
	}

	return response
}

func CheckWithBigShare(company *api.CompanyDocument, pans *[]api.PanDocument) map[string]string {
	var response = map[string]string{}

	for _, pan := range *pans {
		requestBody := api.NewBigshareRequest(company.CompanyCode, pan.PanNumber)
		jsonBody, err := json.Marshal(requestBody)

		if err != nil {
			fmt.Println("failed to parse to json")
			response[pan.PanNumber] = "Failed"
			continue
		}
		resp, err := http.Post(static.SCRAP_URL[static.BIGSHARE], "application/json", strings.NewReader(string(jsonBody)))
		if err != nil {
			fmt.Println("failed request to maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}
		defer resp.Body.Close()
		stringData, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("failed load response from maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}

		var adata map[string]any

		if err := json.Unmarshal(stringData, &adata); err != nil {
			fmt.Println("failed to parse to json reponse body")
			response[pan.PanNumber] = "Failed"
			continue
		}

		adata = adata["d"].(map[string]any)
		if adata == nil {
			fmt.Println("reponse body nil")
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
			continue
		}

		var alloted = adata["ALLOTED"].(string)
		var applied = adata["APPLIED"].(string)

		if alloted == "NON-ALLOTTE" {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_ALLOTED]
		} else if applied == "" && alloted == "" {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
		} else {
			response[pan.PanNumber] = alloted + static.ALLOTMENT_STATUS[static.ALLOTED]
		}
	}

	return response
}

func CheckWithMaashitla(company *api.CompanyDocument, pans *[]api.PanDocument) map[string]string {
	var response = map[string]string{}

	for _, pan := range *pans {
		url := fmt.Sprintf("%v?company=%v&search=%v", static.SCRAP_URL[static.MAASHITLA], company.CompanyCode, pan.PanNumber)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("failed request to maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}
		defer resp.Body.Close()
		stringData, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("failed load response from maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}

		var adata map[string]any

		if err := json.Unmarshal(stringData, &adata); err != nil {
			fmt.Println("failed to parse to json reponse body")
			response[pan.PanNumber] = "Failed"
			continue
		}

		if adata == nil {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
			continue
		}
		var alloted = adata["share_Alloted"].(float64)
		var applied = adata["share_Applied"].(float64)
		if alloted == 0 && applied == 0 {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
		} else if applied > 0 {
			if alloted == 0 {
				response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_ALLOTED]
			} else {
				response[pan.PanNumber] = fmt.Sprintf("%v", alloted) + " " + static.ALLOTMENT_STATUS[static.ALLOTED]
			}
		}
	}

	return response
}

func CheckWithLinkintime(company *api.CompanyDocument, pans *[]api.PanDocument) map[string]string {
	var response = map[string]string{}

	for _, pan := range *pans {
		requestBody := api.NewLinkintimeRequest(company.CompanyCode, pan.PanNumber)
		jsonBody, err := json.Marshal(requestBody)

		if err != nil {
			fmt.Println("failed to parse to json")
			response[pan.PanNumber] = "Failed"
			continue
		}
		resp, err := http.Post(static.SCRAP_URL[static.LINKINTIME], "application/json", strings.NewReader(string(jsonBody)))
		if err != nil {
			fmt.Println("failed request to maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}
		defer resp.Body.Close()
		stringData, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("failed load response from maashitla")
			response[pan.PanNumber] = "Failed"
			continue
		}

		var adata map[string]any

		if err := json.Unmarshal(stringData, &adata); err != nil {
			fmt.Println("failed to parse to json reponse body")
			response[pan.PanNumber] = "Failed"
			continue
		}
		data := adata["d"].(string)
		if data == "" {
			fmt.Println("xml reponse body nil")
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
			continue
		}

		var xmlData api.DataSet
		if err := xml.Unmarshal([]byte(data), &xmlData); err != nil {
			fmt.Println("failed to parse to xml reponse body", err)
			response[pan.PanNumber] = "Failed"
			continue
		}

		var alloted = xmlData.Table.ALLOT
		var applied = xmlData.Table.SHARES

		if alloted == 0 && applied == 0 {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_APPLIED]
		} else if applied > 0 && alloted == 0 {
			response[pan.PanNumber] = static.ALLOTMENT_STATUS[static.NOT_ALLOTED]
		} else {
			response[pan.PanNumber] = fmt.Sprintf("%d", alloted) + static.ALLOTMENT_STATUS[static.ALLOTED]
		}
	}

	return response
}

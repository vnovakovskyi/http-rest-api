package apiserver

type Marketing struct {
	HTTPStatus        int               `json:"httpStatus"`
	MarketingResponse MarketingResponse `json:"response"`
}

type MarketingResponse struct {
	MarketingItems []MarketingItem `json:"items"`
}

type MarketingItem struct {
	Type              string  `json:"type"`
	HarvesterID       string  `json:"harvesterId"`
	CommercialPartner string  `json:"commercialPartner"`
	LogoURL           string  `json:"logoURL"`
	CerebroScore      float64 `json:"cerebro-score"`
	URL               string  `json:"url"`
	Title             string  `json:"title"`
	CleanImage        string  `json:"cleanImage"`
}

package apiserver

type Article struct {
	HTTPStatus      int             `json:"httpStatus"`
	ArticleResponse ArticleResponse `json:"response"`
}

type ArticleResponse struct {
	ArticleItems []ArticleItem `json:"items"`
}

type ArticleItem struct {
	Type         string  `json:"type"`
	HarvesterID  string  `json:"harvesterId"`
	CerebroScore float64 `json:"cerebro-score"`
	URL          string  `json:"url"`
	Title        string  `json:"title"`
	CleanImage   string  `json:"cleanImage"`
}

package apiserver

type Answer struct {
	ArticleItems  []ArticleItem `json:"article_items"`
	MarketingItem MarketingItem `json:"marketing_item"`
}

func Merge(article []ArticleItem, marketing []MarketingItem, target *[]Answer) {
	low := 0
	high := 5
	marketingIndex := 0

	for {
		if low >= len(article) {
			break
		}

		if high > len(article) {
			high = len(article)
		}

		var marketingItem = MarketingItem{}

		if marketingIndex < len(marketing) {
			marketingItem = marketing[marketingIndex]
			marketingIndex++
		} else {
			marketingItem = MarketingItem{
				Type: "Ad",
			}
		}

		*target = append(*target, Answer{
			ArticleItems:  article[low:high],
			MarketingItem: marketingItem,
		})

		low += 5
		high += 5
	}
}

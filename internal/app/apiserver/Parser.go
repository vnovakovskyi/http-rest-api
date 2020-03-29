package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseArticle(url string) Article {
	article := Article{}

	data := getJSON(url)
	err := json.Unmarshal(data, &article)

	if err != nil {
		fmt.Println("Error2: ", err)
	}

	return article
}

func ParseMarketing(url string) Marketing {
	marketing := Marketing{}

	data := getJSON(url)
	err := json.Unmarshal(data, &marketing)

	if err != nil {
		fmt.Println("Error2: ", err)
	}

	return marketing
}

func getJSON(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error1: ", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	return data
}

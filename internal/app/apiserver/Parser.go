package apiserver

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ParseArticle(url string) Article {
	article := Article{}

	data, err := getJSON(url)
	if err != nil {
		logrus.Info(err)
		return article
	}

	err = json.Unmarshal(data, &article)

	if err != nil {
		logrus.Info(err)
	}

	return article
}

func ParseMarketing(url string) Marketing {
	marketing := Marketing{}

	data, err := getJSON(url)
	if err != nil {
		logrus.Info(err)
		return marketing
	}

	err = json.Unmarshal(data, &marketing)

	if err != nil {
		logrus.Info(err)
	}

	return marketing
}

func getJSON(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		logrus.Info(err)
		return nil, err
	}

	data, _ := ioutil.ReadAll(res.Body)
	return data, nil
}

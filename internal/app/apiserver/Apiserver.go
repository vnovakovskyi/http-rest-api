package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/merger", s.HandleMerger())
}

func (s *APIServer) HandleMerger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		articleUrl := "https://storage.googleapis.com/aller-structure-task/articles.json"
		article := ParseArticle(articleUrl)

		marketingUrl := "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
		marketing := ParseMarketing(marketingUrl)

		var answer []Answer
		Merge(article.ArticleResponse.ArticleItems, marketing.MarketingResponse.MarketingItems, &answer)

		err := json.NewEncoder(w).Encode(answer)

		if err != nil {
			s.logger.Error(err)
		}
	}
}

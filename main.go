package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf := buildConfig()
	http.Handle("/slack-relay", &slackRelay{conf})
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", conf.ListenPort), nil)
}

type slackRelay struct {
	conf *config
}

func (s *slackRelay) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Printf("Could not parse form: %s", err)
		return
	}
	log.Println("Got form post:")
	for name, values := range req.PostForm {
		for _, value := range values {
			log.Printf("%s=%s", name, value)
		}
	}
}

type config struct {
	ListenPort int
}

func buildConfig() *config {
	config := &config{}
	flag.IntVar(&config.ListenPort, "port", 8080, "what http port to listen on")
	flag.Parse()
	return config
}

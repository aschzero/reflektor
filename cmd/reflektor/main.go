package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"reflektor"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	configFilePath := flag.String("config.file", "", "Relative path to config file yaml")

	flag.Parse()

	configData, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("unable to read config file")
		return
	}

	jobs, err := reflektor.NewJobCollection(configData)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("unable to parse config file")
	}

	r := &reflektor.Reflektor{
		Jobs: jobs,
	}

	r.RegisterJobs()

	<-sigc
	os.Exit(0)
}

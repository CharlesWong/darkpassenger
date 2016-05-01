package util

import (
	"log"
	"log/syslog"
	"os"
)

func SetupLogOptions(logTo string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
	if logTo == "syslog" {
		w, err := syslog.New(syslog.LOG_INFO, "darkpassenger")
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(w)
	}
}

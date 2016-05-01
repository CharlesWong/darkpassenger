package util

import (
	"log"
	"os"
)

func SetupLogOptions(logTo string) {
	log.SetOutput(os.Stdout)
}

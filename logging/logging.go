package logging

import (
	"appointments_scheduler/config"
	"log/syslog"

	log "github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// InitLogging initializes the logging.
func InitLogging(c config.LogConfig) {
	level, err := log.ParseLevel(c.Level)
	if err != nil {
		log.WithError(err).Fatal("failed to parse logging level")
	}

	log.SetLevel(level)
	log.SetReportCaller(c.EnableReportCaller)
	log.SetFormatter(&log.JSONFormatter{})

	if c.EnableSyslog {
		log.Info("initializing syslog hook")

		hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_DEBUG, c.AppName)

		if err != nil {
			log.WithError(err).Fatal(
				"failed to initialize syslog hook")
		} else {
			log.AddHook(hook)
		}
	}
}

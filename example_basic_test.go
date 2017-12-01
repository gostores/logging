package logging_test

import (
	"os"

	"github.com/gostores/logging"
)

func Example_basic() {
	var log = logging.New()
	log.Formatter = new(logging.JSONFormatter)
	log.Formatter = new(logging.TextFormatter)                     //default
	log.Formatter.(*logging.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logging.DebugLevel
	log.Out = os.Stdout

	// file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logging.Entry)
			log.WithFields(logging.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	log.WithFields(logging.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logging.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logging.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logging.Fields{
		"temperature": -4,
	}).Debug("Temperature changes")

	log.WithFields(logging.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")

	// Output:
	// level=debug msg="Started observing beach" animal=walrus number=8
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=debug msg="Temperature changes" temperature=-4
	// level=panic msg="It's over 9000!" animal=orca size=9009
	// level=error msg="The ice breaks!" err_animal=orca err_level=panic err_message="It's over 9000!" err_size=9009 number=100 omg=true
}

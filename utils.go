package tim_utils_log

import (
	"fmt"
	"log"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02  15:04:05") + " [DEBUG] " + string(bytes))
}

func StdOut(iString string) {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	log.Println(iString)

}

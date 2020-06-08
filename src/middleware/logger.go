package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/negroni"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")

	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

func cleanRequestStr(str []byte, err interface{}) string {
	// error, empty, or they're trying to flood us
	if err != nil || len(str) <= 0 || len(str) > 1024 {
		return "{}"
	}

	updatedStr := string(str)
	updatedStr = strings.ReplaceAll(updatedStr, "\n", "")
	updatedStr = strings.ReplaceAll(updatedStr, "\t", "")

	return updatedStr
}

func LoggingMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	current := time.Now()

	tempQueryStr, err := json.Marshal(req.URL.Query())
	queryStr := cleanRequestStr(tempQueryStr, err)

	body, err := ioutil.ReadAll(req.Body)
	bodyStr := cleanRequestStr(body, err)

	next(rw, req)

	res := rw.(negroni.ResponseWriter)

	finished := time.Since(current)

	log.Printf("{ \"User-Agent\": \"%s\", \"Url\": \"%s\", \"Host\": \"%s\", \"Uri\": \"%s\", \"Method\": \"%s\", \"Status Code\": %d, \"Query\": \"%s\", \"Body\": \"%s\", \"Response Time\": \"%s\" \"Content-Type\": \"%s\", \"Content-Length\": %d }",
		req.Header.Get("User-Agent"), req.Host+req.RequestURI, req.Host, req.RequestURI, req.Method, res.Status(), queryStr, bodyStr, finished, req.Header.Get("Content-Type"), req.ContentLength)

}

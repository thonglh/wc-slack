package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiBaseURL = "http://worldcup.sfg.io"

var commandMappings = map[string]string{
	"country": `matches/country?fifa_code=`,
	"matches": `matches/`,
}

func requestAPI(url string) []Match {
	rs, err := http.Get(url)

	// Process response
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return []Match{}
	}

	// var result map[string]interface{}
	var result []Match

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return []Match{}
	}

	return result
}

func buildSlackAttachments(matches []Match) []MatchAttachment {
	attachments := make([]MatchAttachment, len(matches))
	for idx, match := range matches {
		matchTime, _ := time.Parse(time.RFC3339, match.Datetime)
		timestamp := matchTime.Unix()

		attachments[idx] = MatchAttachment{
			Color:      "#36a64f",
			Title:      fmt.Sprintf("%s vs %s", match.HomeTeam.Country, match.AwayTeam.Country),
			Text:       fmt.Sprintf("Match status: %s", match.Status),
			Footer:     "Worlcup Bot",
			FooterIcon: "https://upload.wikimedia.org/wikipedia/en/thumb/6/67/2018_FIFA_World_Cup.svg/227px-2018_FIFA_World_Cup.svg.png",
			Timestamp:  timestamp,
			Fields: []Field{
				{
					Title: strconv.Itoa(match.HomeTeam.Goals),
					Value: match.HomeTeam.Country,
					Short: true,
				},
				{
					Title: strconv.Itoa(match.AwayTeam.Goals),
					Value: match.AwayTeam.Country,
					Short: true,
				},
			},
		}
	}
	return attachments
}

func writeResponse(w http.ResponseWriter, body []Match) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(buildSlackAttachments(body))

	if err != nil {
		fmt.Fprint(w, `{"error": "no match found"}`)
	}

	fmt.Fprintf(w, `{"attachments": %s}`, json)
}

func buildRequestURL(r *http.Request) string {
	command := r.FormValue("command")
	text := r.FormValue("text")

	subPath := strings.SplitN(text, " ", 2)[0]
	path, ok := commandMappings[command]

	if !ok {
		path = command + "/"
	}

	return apiBaseURL + "/" + path + subPath
}

func command(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	var body []Match

	if err == nil {
		url := buildRequestURL(r)

		body = requestAPI(url)

	} else {
		body = []Match{}
	}

	writeResponse(w, body)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", command)
	log.Fatal(http.ListenAndServe(addr, nil))
}

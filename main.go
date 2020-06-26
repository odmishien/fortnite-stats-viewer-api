package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var APIKEY = os.Getenv("APIKEY")

func main() {
	http.HandleFunc("/global-stats", globalStatsHandler)
	http.HandleFunc("/recent-matches", recentMatchesHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func globalStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		accountID := r.URL.Query().Get("account_id")
		res, err := getGlobalStats(accountID)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func recentMatchesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		accountID := r.URL.Query().Get("account_id")
		res, err := getRecentMatches(accountID)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(res))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func getGlobalStats(accountID string) (string, error) {
	var endpoint = "https://fortniteapi.io/stats"
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("account", accountID)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", APIKEY)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return "", err
	}

	return string(body), nil
}

func getRecentMatches(accountID string) (string, error) {
	var endpoint = "https://fortniteapi.io/matches"
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("account", accountID)
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", APIKEY)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return "", err
	}

	return string(body), nil
}

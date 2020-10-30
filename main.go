package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


type timeStamp struct{
	Unix string`json:"unix"`
	Natural string `json:"natural"`
}

var dateFormats = []string{
	"January 2 2006",
	"Jan 2 2006",
	"2 Jan 2006",
	"2 January 2006",

	"January 2, 2006",
	"Jan 2, 2006",
	"2 Jan, 2006",
	"2 January, 2006",

	"2006 January 2",
	"2006 Jan 2",
	"2006 2 Jan",
	"2006 2 January",

	"2006, January 2",
	"2006, Jan 2",
	"2006, 2 Jan",
	"2006, 2 January",
}

func main(){
	port := os.Getenv("PORT")
	
	if port == ""{
		port = "3000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/{date}", handler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handler(w http.ResponseWriter, r *http.Request){
	en := json.NewEncoder(w)
	date := mux.Vars(r)["date"]

	i , err := strconv.ParseInt(date, 10, 64)

	if err != nil{
		var res timeStamp
		for _, layout := range dateFormats{
			if getTime, err := time.Parse(layout, date); err == nil{
				res = timeStamp{
					Unix: strconv.FormatInt(getTime.UTC().Unix(), 10),
					Natural: getTime.Format("January 2, 2006"),
				}
			}
		}	
		en.Encode(res)
		return
	}

	t := time.Unix(i, 0)
	en.Encode(timeStamp{
		Unix:    date,
		Natural: t.Format("January 2, 2006"),
	})
}
package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


type timeStamp struct{
	Unix int64 `json:"unix"`
	Utc time.Time `json:"utc"`
}

const(
	layout = "01-02-2006"
)

func main(){
	port := os.Getenv("PORT")
	
	date := "12-25-2015"

	now, err := time.Parse(layout, date)
	if err != nil{
		fmt.Println(err)
	}

	timestamp := timeStamp{
		Unix: now.Unix(),
		Utc: now,
	}

	jsonUnix, _ := json.Marshal(timestamp)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		w.Write(jsonUnix)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
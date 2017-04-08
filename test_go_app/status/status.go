package status

import (
	"net/http"
	"time"
	"log"
	"database/sql"
	"fmt"
	"encoding/json"
	"github.com/ZephroC/test_go_app/config"
)

type StatusResponse struct {
	Time string `json:"time"`
	StatusCode uint32 `json:"status_code"`
	Messages []string `json:"messages"`
}

func StatusHandler(config config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		responses := make([]string, 0,5)
		time_now := time.Now().UTC()
		log.Println("Responding to status request")
		var dbResponse StatusResponse
		dbResponse.StatusCode = 200
		dbResponse.Time = time_now.Format(time.RFC1123)
		var dbConf = config.Database
		db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
			dbConf.Username,
			dbConf.Password,
			dbConf.Name,
			dbConf.Host,
			dbConf.Port))
		if err == nil {
			responses = append(responses,"Connected to DB")
		} else {
			responses = append(responses,err.Error())
			log.Fatal(err)
		}
		defer db.Close()
		dbResponse.Messages = responses
		fmt.Println(dbResponse)
		resp_str, err := json.MarshalIndent(dbResponse,"","\t")
		if err!=nil {
			log.Println("Error formatting json")
			log.Println(err.Error())
		}

		fmt.Println(string(resp_str))
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		fmt.Fprintln(response, string(resp_str))
	})
}

package books

import (
	"github.com/ZephroC/test_go_app/config"
	"net/http"
	"log"
	"database/sql"
	"fmt"
	"encoding/json"
)

type Book struct {
	Id uint32 `json:"id"`
	Title string `json:"title"`
}

func ListBooks(config config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var dbConf = config.Database
		db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
			dbConf.Username,
			dbConf.Password,
			dbConf.Name,
			dbConf.Host,
			dbConf.Port))
		if err != nil {
			log.Printf("Didn't connect");
			log.Fatal(err)
		}
		rows, err := db.Query("SELECT * FROM books_view;")
		if err !=nil {
			log.Printf("Didn't query");
			log.Fatal(err)
		}
		books := make([]Book, 0,5)
		for rows.Next() {
			var book_id uint32
			var title string
			if err := rows.Scan(&book_id,&title); err != nil {
				log.Fatal(err)
			}
			var book = Book{ book_id, title}
			books = append(books,book)
			fmt.Printf("%d is %s\n", book_id, title)
		}
		defer db.Close()


		resp_str, err := json.MarshalIndent(books,"","\t")
		if err!=nil {
			log.Println("Error formatting json")
			log.Println(err.Error())
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		fmt.Fprintln(response, string(resp_str))
	})
}
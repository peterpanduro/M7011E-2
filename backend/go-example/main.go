package go-example

import (
	"encoding/json"
	"net/http"
)

type Result struct {
    Message string `json:"message"`
    Status int `json:"status"`
}

func main() {
    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        result := Result {
            Message: "Success",
            Status: 200,
        }

        json.NewEncoder(w).Encode(result)
    })

    http.ListenAndServe(":8080", nil)
}
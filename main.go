package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}
type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handleStatusCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sample go server active")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "sample server working")
}

func handleSomeGetReq(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "not a valid request method")
		return
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "JSON object not valid")
		return
	}

	response := ResponseBody{
		Status:  "success",
		Message: "Request processed successfully.",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Failed to marshal response JSON.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	fmt.Println("server starts......")
	http.HandleFunc("/status", handleStatusCheck)
	http.ListenAndServe(":8080", nil)
}

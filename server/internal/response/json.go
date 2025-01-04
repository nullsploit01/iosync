package response

import (
	"encoding/json"
	"net/http"
)

type ResponsePagination struct {
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	TotalItems  int `json:"total_items,omitempty"`
}

type Response struct {
	Error      bool               `json:"error"`
	Message    string             `json:"message,omitempty"`
	Data       any                `json:"data,omitempty"`
	Pagination ResponsePagination `json:"pagination,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data Response) error {
	return JSONWithHeaders(w, status, data, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data Response, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

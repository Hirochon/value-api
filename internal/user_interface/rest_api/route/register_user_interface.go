package route

import (
	"fmt"
	"io"
	"net/http"
)

// unmarshal する際のフィールドの頭文字は大文字でなければいけない
type registerUser struct {
	Name     string `json:"name"`
	Gender   int    `json:"gender"`
	BirthDay string `json:"birthDay"`
}

// RegisterUser は、ユーザーを登録する
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error IO: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	parsedData, err := unmarshalToStruct[registerUser](data)
	if err != nil {
		fmt.Printf("Error parse data: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", parsedData)
	w.WriteHeader(http.StatusCreated)
}

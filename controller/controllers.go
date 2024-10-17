package controller

import (
	"encoding/json"
	"net/http"
)

type FindPiarReq struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type FindPairRes struct {
	Solutions [][]int `json:"solutions"`
	Error     string  `json:"error,omitempty"`
}

const ErrInvalidInput = "invalid input for api"

// controller function to handle endpoint
func FindPairController(writer http.ResponseWriter, request *http.Request) {
	var req FindPiarReq
	err := json.NewDecoder(request.Body).Decode(&req)
	var resp FindPairRes
	if err != nil {
		resp.Error = ErrInvalidInput
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(resp)
		return
	}

	resp.Solutions = findPiar(req.Numbers, req.Target)

	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

}

// function to implement logic to find solution
func findPiar(arr []int, target int) (res [][]int) {
	seen := make(map[int]int)

	for i, num := range arr {
		if id, ok := seen[target-num]; ok && id != i {
			res = append(res, []int{i, id})
		}
		seen[num] = i
	}
	return res

}

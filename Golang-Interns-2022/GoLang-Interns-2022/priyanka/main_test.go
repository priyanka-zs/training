package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func Test_Main(t *testing.T) {
	testcases := []struct {
		desc   string
		method string
		url    string
		body   []byte
		status int
	}{
		{"create", http.MethodPost, "http://localhost:8080/car",
			[]byte(`{"Name": "porsche 1","Year": 2021,"Brand": "porsche","Fuel": "diesel",
			"Engine": {"Displacement": 110,"NoOfCylinders": 2,"Ranges": 0}}`), http.StatusCreated},
		{"GetById", http.MethodGet, "http://localhost:8080/car/e24a702e-3ff0-414c-b08a-bdf0ab91c630",
			nil, http.StatusOK},
		{"GetByBrand", http.MethodGet, "http://localhost:8080/car?brand=bmw&&isEngine=true",
			nil, http.StatusOK},
		{"Update", http.MethodPut, "http://localhost:8080/car/e24a702e-3ff0-414c-b08a-bdf0ab91c630",
			[]byte(`{"name":"porsche 2","year":2021,"brand":"porsche","fuel":"diesel","engine":{
		"Displacement": 110,"NoOfCylinders": 2,"Ranges": 0}}`), http.StatusOK},
		{"create invalid parameters", http.MethodPost, "http://localhost:8080/car",
			[]byte(`{"Name": "bmw 2", "Year": 2000,"Brand":"tesla","Fuel":"diesel",
			"Engine":{"Displacement": 110,"NoOfCylinders": 2,"Ranges": 0}`), http.StatusBadRequest},
		{"GetById Invalid Id", http.MethodGet, "http://localhost:8080/car/f24a702e-3ff0-414c-b08a-bdf0ab91c630",
			nil, http.StatusBadRequest},
		{"GetByBrand with invalid brand", http.MethodGet, "http://localhost:8080/car?brand=bmm&&isEngine=true", nil, http.StatusBadRequest},
		{"Update invalid id", http.MethodPut, "http://localhost:8080/car/f24a702e-3ff0-414c-b08a-bdf0ab91c630",
			[]byte(`{"name":"porsche 2","year":2021,"brand":"porsche","fuel":"diesel","engine":{
			"Displacement": 110,"NoOfCylinders": 2,"Ranges": 0}}`), http.StatusBadRequest},
		{"Delete", http.MethodDelete, "http://localhost:8080/car/d50909c6-bd46-437f-b62f-6683e00d532c", nil, http.StatusNoContent},
		{"Delete with invalid id", http.MethodDelete,
			"http://localhost:8080/car/045b658e-9160-4f55-8e5a-be8ceb13fbf", nil, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.body))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("api-key", "123")

		h := http.Client{}

		res, err := h.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		err = res.Body.Close()
		if err != nil {
			return
		}

		if tc.status != res.StatusCode {
			t.Errorf("In method %v expected status %v got %v\n", tc.method, tc.status, res.StatusCode)
		}
	}
}

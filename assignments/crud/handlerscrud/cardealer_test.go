package handlerscrud

import (
	"assignments/crud/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"

	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func DbConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "priyanka"
	dbPass := "Hani@2001"
	dbName := "cardealer"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)"+"/"+dbName)
	if err != nil {
		//log.Fatal(fmt.Errorf("unexpected error %v", err.Error()))
		return nil, err
	}
	return db, nil
}

//TestPost is used to test post method
func TestPost(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)

	testCases := []struct {
		desc     string
		input    models.Car
		exOutput models.Car
	}{

		{desc: "success case", input: models.Car{Name: "johnny", Year: 2010, Brand: "mercedes", Fuel: "petrol", Engine: models.Engine{Displacement: 12, NoOfCylinders: 5}},
			exOutput: models.Car{Name: "johnny", Year: 2010, Brand: "mercedes", Fuel: "petrol", Engine: models.Engine{Displacement: 12, NoOfCylinders: 5}}},
		//{desc: "year less than 1980", input: models.Car{Name: "myCar1", Year: 1970, Brand: "bmw", Fuel: "electric", Engine: models.Engine{Displacement: 120, NoOfCylinders: 2, Ranges: 30}}, exOutput: models.Car{}},
		//{"year greater than 2022", models.Car{Name: "myCar2", Year: 2030, Brand: "Ferrari", Fuel: "electric", Engine: models.Engine{Displacement: 120, NoOfCylinders: 2, Ranges: 20}}, models.Car{}},
	}
	for _, v := range testCases {
		data, err := json.Marshal(v.input)
		if err != nil {
			t.Errorf("error:%v", err)
		}
		req := httptest.NewRequest("POST", "/cars", bytes.NewReader(data))
		w := httptest.NewRecorder()
		d.Post(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("wrong input")
		}
		body, er := io.ReadAll(resp.Body)
		if er != nil {
			t.Errorf("Unexpected Error got %v", er)
		}
		if assert.Equal(t, string(body), v.exOutput) {
			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", v.desc, v.exOutput, string(body))
		}
	}
}

//TestGet is used to test get method
/*func TestGet(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testCases := []struct {
		desc           string
		input          string
		expectedOutput []models.Car
	}{
		{desc: "success cse", input: "/ferrari&engine=included", expectedOutput: []models.Car{{uuid.MustParse("8f14a65f-3032-42c8-a196-1cf66d11b932"),
			"myCar", 2010, "/tesla", "electric", models.Engine{56, 2, 4}}}},
		{"url without brand name", "/", nil},
		{"success case with brand", "/tesla", []models.Car{{uuid.MustParse("8f14a65f-3032-42c8-a196-1cf66d11b932"),
			"john", 2000, "/tesla", "electric", models.Engine{}}}},
	}
	for _, v := range testCases {
		req := httptest.NewRequest("GET", v.input, nil)
		w := httptest.NewRecorder()
		d.Get(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("wrong input")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected Error got %v", err)
		}
		if assert.Equal(t, body, v.expectedOutput) {
			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", v.desc, v.expectedOutput, string(body))
		}
	}
}*/

//TestGetById is used to test get method
/*func TestGetById(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testCases := []struct {
		desc           string
		input          string
		expectedOutput string
	}{
		{"success testcase", "/6ba7b810-9dad-11d1-80b4-00c04fd430c8", "details displayed successfully"},
		{"url with wrong format for uuid", "/1", "invalid"},
		{"url without uuid", "/", "Status bad request"},
	}
	for _, v := range testCases {
		req := httptest.NewRequest("GET", v.input, nil)
		w := httptest.NewRecorder()
		d.GetById(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("wrong input")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected Error got %v", err)
		}
		if string(body) != v.expectedOutput {
			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", v.desc, v.expectedOutput, string(body))
		}
	}
}*/

//TestDelete is used to test the delete method
/*func TestDelete(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testCases := []struct {
		desc           string
		input          string
		expectedOutput string
	}{
		{"success case", "/6ba7b810-9dad-11d1-80b4-00c04fd430c8", "details deleted successfully"},
		{"url without id", "/", "Status Bad Request"},
		{"wrong format", "/l", "invalid"},
	}
	for _, v := range testCases {
		req := httptest.NewRequest("DELETE", v.input, nil)
		w := httptest.NewRecorder()
		d.Delete(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("wrong input")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected Error got %v", err)
		}
		if string(body) != v.expectedOutput {

			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", v.desc, v.expectedOutput, string(body))
		}
	}
}*/

//TestUpdate is used to test the update method
/*func TestUpdate(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testCases := []struct {
		desc     string
		url      string
		input    models.Car
		exOutput string
	}{
		{"success case", "/6ba7b810-9dad-11d1-80b4-00c04fd430c8", models.Car{Name: "myCar", Year: 2020, Brand: "tesla", Fuel: "electric", Engine: models.Engine{120, 2, 40}}, "details entered successfully"},
		{"url without id", "/", models.Car{Name: "myCar", Year: 2020, Brand: "tesla", Fuel: "electric", Engine: models.Engine{120, 2, 40}}, "status bad request"},
		{"invalid", "/p", models.Car{Name: "myCar", Year: 2020, Brand: "tesla", Fuel: "electric", Engine: models.Engine{120, 2, 40}}, "invalid"},
	}
	for _, v := range testCases {
		data, err := json.Marshal(v.input)
		if err != nil {
			t.Errorf("error:%v", err)
		}
		req := httptest.NewRequest("PUT", v.url, bytes.NewReader(data))
		w := httptest.NewRecorder()
		d.Update(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {

			t.Errorf("wrong input")
		}
		body, er := io.ReadAll(resp.Body)
		if er != nil {
			t.Errorf("Unexpected Error got %v", er)
		}
		if string(body) != v.exOutput {

			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", v.desc, v.exOutput, string(body))
		}
	}
}
*/

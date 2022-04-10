package handlerscrud

import (
	"assignments/crud/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type dB struct {
	DB *sql.DB
}

func New(db *sql.DB) dB {
	return dB{DB: db}
}
func (d dB) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	res, _ := ioutil.ReadAll(r.Body)
	data := models.Car{}
	err := json.Unmarshal(res, &data)
	if err != nil {
		log.Fatal(fmt.Errorf("unexpected error:%v", err))
	}

	if data.Year < 1980 || data.Year > 2022 {
		w.WriteHeader(http.StatusBadRequest)
		//w.Write([]byte("year should be between 1980 and 2022"))
		k, _ := json.Marshal(models.Car{})
		w.Write(k)
		return
	}
	data.Id = uuid.New()
	_, err = d.DB.Exec("INSERT INTO car(Id,Name,Year,Brand,FuelType)VALUES(?,?,?,?,?)", data.Id, data.Name, data.Year, data.Brand, data.Fuel)
	if err != nil {
		log.Fatal(fmt.Errorf("%v", err))
	}
	_, err = d.DB.Exec("INSERT INTO engine(Id,Displacement,No_of_cylinders,`Range`)VALUES ( ?,?,?,?)", data.Id, data.Engine.Displacement, data.Engine.NoOfCylinders, data.Engine.Ranges)
	if err != nil {
		log.Fatal(fmt.Errorf("%v", err))
	}
	k, _ := json.Marshal(res)
	w.Write(k)
	/*_, err = w.Write(res)
	if err != nil {
		log.Fatal(fmt.Errorf("%v", err))
	}*/

}

func (d dB) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	para := mux.Vars(r)
	nId := para["Id"]
	res := models.Car{}
	d.DB.QueryRow("SELECT * FROM car WHERE Id=?", nId).Scan(&res.Id, &res.Name, &res.Year, &res.Brand, &res.Fuel)
	d.DB.QueryRow("SELECT Displacement,No_of_cylinders,`Range` FROM engine WHERE Id=?", nId).Scan(&res.Engine.Displacement, &res.Engine.NoOfCylinders, &res.Engine.Ranges)
	//fmt.Println(res, err)
	data, err := json.Marshal(res)
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(fmt.Errorf("%v", err))
	}
}

func (d dB) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	brand := r.URL.Query().Get("brand")
	engine := r.URL.Query().Get("engine")
	if len(brand) != 0 && len(engine) == 0 {
		w.Write([]byte(" data with Brand " + brand))
		queryCar, err := d.DB.Query("Select * from car where Brand=?", brand)
		if err != nil {
			fmt.Errorf("Unexpected error %v", err)
		}
		var res []models.Car
		cars := models.Car{}
		for queryCar.Next() {
			err := queryCar.Scan(&cars.Id, &cars.Name, &cars.Year, &cars.Brand, &cars.Fuel)
			if err != nil {
				fmt.Errorf("Unexpected error %v", err)
			}
			res = append(res, cars)
		}
		output, err := json.Marshal(res)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		w.Write(output)
	} else if len(brand) != 0 && len(engine) != 0 {
		queryCar, err := d.DB.Query("Select * from car where Brand=?", brand)
		if err != nil {
			fmt.Errorf("Unexpected error %v", err)
		}
		cars := models.Car{}
		var res []models.Car
		for queryCar.Next() {
			err := queryCar.Scan(&cars.Id, &cars.Name, &cars.Year, &cars.Brand, &cars.Fuel)
			if err != nil {
				fmt.Errorf("Unexpected error %v", err)
			}
			d.DB.QueryRow("Select Displacement,No_of_cylinders,`Range` from engine where Id=?", cars.Id).
				Scan(&cars.Engine.Displacement, &cars.Engine.NoOfCylinders, &cars.Engine.Ranges)
			res = append(res, cars)
		}

		json.NewEncoder(w).Encode(cars)
	}
}

func (d dB) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	para := mux.Vars(r)
	Id := para["Id"]
	var value string
	err := d.DB.QueryRow("SELECT Id FROM car WHERE Id=?", Id).Scan(&value)
	if err != nil || value == "" {
		fmt.Errorf("Unexpected error %v", err)
		w.Write([]byte("Entity do not exists"))
		return
	}
	d.DB.QueryRow("Delete  from car where Id=?", Id)
	d.DB.QueryRow("Delete  from engine where Id=?", Id)
	w.Write([]byte("Row deleted successfully"))
}

func (d dB) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	data := models.Car{}
	para := mux.Vars(r)
	Id := para["Id"]
	var value string
	err := d.DB.QueryRow("select Id from car where Id=?", Id).Scan(&value)
	if err != nil || value == "" {
		fmt.Errorf("Unexpected error %v", err)
		w.Write([]byte("Entity do not exists"))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	data.Id = uuid.MustParse(Id)
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Errorf("unexpected error %v", err)
	}
	if data.Year < 1980 || data.Year > 2022 {
		w.Write([]byte("year must be between 1980 and 2022"))
		return
	}
	if data.Brand != "tesla" && data.Brand != "porsche" && data.Brand != "ferrari" && data.Brand != "mercedes" && data.Brand != "bmw" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if data.Fuel != "petrol" && data.Fuel != "diesel" && data.Fuel != "electric" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if data.Fuel == "electric" {
		data.Engine.Displacement = 0
		data.Engine.NoOfCylinders = 0
		if data.Engine.Ranges < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if data.Fuel == "petrol" || data.Fuel == "diesel" {
		if data.Engine.Displacement <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data.Engine.NoOfCylinders <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data.Engine.Ranges = 0
	}
	_, err = d.DB.Exec("UPDATE car SET Name=? , Year=? , Brand=? , FuelType=? WHERE Id=?", data.Name, data.Year, data.Brand, data.Fuel, data.Id)
	if err != nil {
		fmt.Errorf("unexpected error %v", err)
	}

	_, err = d.DB.Exec("UPDATE engine SET Displacement=? , No_of_cylinders=? , `Range`=? WHERE Id=?", data.Engine.Displacement, data.Engine.NoOfCylinders, data.Engine.Ranges, data.Id)
	if err != nil {
		fmt.Errorf("unexpected error %v", err)
	}
	json.NewEncoder(w).Encode(data)
}

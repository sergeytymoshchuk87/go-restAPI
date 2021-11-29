package main

// P.S: Procedures used more than once are encapsulated in helper functions in file helpers.go

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorilla/mux"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize the database connection pool
func init() {
	var err error
	db, err = sql.Open("mysql", "ffqnJuiViE:rqdDqpjfED@tcp(remotemysql.com:3306)/ffqnJuiViE")
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	} else {
		fmt.Println("Connected") // Note that this will still print even though there's no internet connection
	}
}

// Detail represents the food's info
type Detail struct {
	ID       int
	Name     string
	Price    float64
	MakeTime float64
}

// GetByID returns the details of food by and Id
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Check if id is numeric
	err := checkdata(&id)
	if err != nil {
		w.Write([]byte("ID value in URL Request is Invalid"))
		return
	}

	tmp := &Detail{}
	err = query("ID", tmp, id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	render(tmp, &w)
}

// GetByName returns the details of the food by searching name
// And assumes that no two foods will have the same exact Name
func GetByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if name == "" {
		w.Write([]byte("Data needed is empty in URL"))
		return
	}

	tmp := &Detail{}
	err := query("Name", tmp, name)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	render(tmp, &w)
}

// AddNew add new food to the Database
// This method assumes a form or cURL or a tool like postman will be used to send the data
func AddNew(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	price := r.FormValue("price")
	time := r.FormValue("time")

	if name == "" {
		w.Write([]byte("Data Provided Is Not Complete"))
		return
	}

	// Check if price and time are numeric
	err := checkdata(&price, &time)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = db.Exec("INSERT INTO Food(Name,Price,MakeTime) VALUES(?,?,?)", name, price, time)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write([]byte("Food Added Successfully"))
}

// Remove deletes the row with the given ID
func Remove(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	// Check if id is numeric
	err := checkdata(&id)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
	}

	_, err = db.Exec("DELETE FROM Food WHERE ID = ?", id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("Food Deleted Successfully"))
}

// ShowAll lists all books
func ShowAll(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * From Food")
	if err != nil {
		log.Fatal(err)
	}

	tmp := make([]*Detail, 0)
	for rows.Next() {
		fd := new(Detail)
		err := rows.Scan(&fd.ID, &fd.Name, &fd.Price, &fd.MakeTime)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Fatal(err)
			return
		}
		tmp = append(tmp, fd)
	}
	defer rows.Close()

	json, err := json.MarshalIndent(&tmp, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}

// UpdateName modifies the Name field for a given ID
func UpdateName(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	name := r.FormValue("name")

	err := update("Name", name, id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("Data Updated Succesfully"))
}

// UpdatePrice modifies the Price field for a given ID
func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	price := r.FormValue("price")

	err := update("Price", price, id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("Data Updated Succesfully"))
}

// UpdateTime modifies the MakeTime field for a given ID
func UpdateTime(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	time := r.FormValue("time")

	err := update("MakeTime", time, id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("Data Updated Succesfully"))
}

// UpdateAll takes all data at once if user wants to modify all columns for the data
func UpdateAll(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	price := r.FormValue("price")
	name := r.FormValue("name")
	time := r.FormValue("time")

	_, err := db.Exec("UPDATE Food SET Name = ?,Price = ?,MakeTime =? WHERE ID = ?", name, price, time, id)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("Data Updated Succesfully"))
}

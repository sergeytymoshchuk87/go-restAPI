package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// checkdata checks the values in the URL for integer equivalence
// Pointers used to save memory
func checkdata(data ...*string) error {
	for ix := range data {
		if *data[ix] != "" {
			// ParseFloat can handle both integers and floats hence,...
			_, err := strconv.ParseFloat((*data[ix]), 64)
			if err != nil {
				return err
			}
		} else {
			return errors.New("some needed data is empty in URL")
		}
	}
	return nil
}

// query queries the database and scans the result into the struct passed
func query(crit string, struc *Detail, variable string) error {

	qry := "SELECT * From Food WHERE " + crit + " = ?"
	row := db.QueryRow(qry, variable)

	return row.Scan(&struc.ID, &struc.Name, &struc.Price, &struc.MakeTime)
}

// updates modifies data in the database
func update(crit string, var1 string, var2 string) error {
	updqry := "UPDATE Food SET " + crit + " = ? WHERE ID = ?"
	_, err := db.Exec(updqry, var1, var2)
	if err != nil {
		return err
	}
	return nil
}

// render processes and send the JSON to the client
func render(tmp *Detail, w *http.ResponseWriter) {
	ref := *tmp
	json, err := json.MarshalIndent(&ref, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	res := *w
	res.Write(json)
}

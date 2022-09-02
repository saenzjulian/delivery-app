package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings" 
)

type Route struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool `json:"finished"`
}

func (r *Route) LoadPositions() error { // everytime r is modified, the (*) changes the values
	if r.ID == "" {
		return errors.New("Route ID is required")
	}
	file, err := os.Open("destinations/" + r.ID + ".txt") 
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		data := strings.Split(scanner.Text(), ",")
		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{
			Latitude: latitude,
			Longitude: longitude,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error){
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for i, v := range r.Positions { 
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Latitude, v.Longitude}
		route.Finished = false
		if total-1 == i {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
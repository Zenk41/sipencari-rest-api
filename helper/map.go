package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/util"
	"github.com/labstack/echo/v4"
)

type Response struct {
	FormatedAddress string `json:"formated_address"`
	Status          string `json:"status"`
}

func GetAdressFromLatLng(c echo.Context, lat float64, lng float64) (Response, error) {
	APIKey := util.GetEnv("MAP_API_KEY")
	URL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%.6f,%.6f&key=%s", lat, lng, APIKey)
	var defstatus string = "REQUEST_DENIED"
	var deformat string = "Unknown Place Name"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return Response{FormatedAddress: deformat, Status: defstatus}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{FormatedAddress: deformat, Status: defstatus}, err
	}
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Response{FormatedAddress: deformat, Status: defstatus}, err
	}

	data := make(map[string]interface{})
	var compoundCode string
	var status string
	err = json.Unmarshal(resData, &data)
	if err != nil {
		fmt.Println(err)
		return Response{FormatedAddress: deformat, Status: defstatus}, err
	}

	plscode := data["plus_code"].(map[string]interface{})
	compoundCode, ok := plscode["compound_code"].(string)
	if !ok {
		globeCode := plscode["global_code"].(string)
		return Response{FormatedAddress: globeCode, Status: "REPLACING ADDRESS WITH GLOBAL CODE"}, nil
	}

	status = data["status"].(string)

	if compoundCode == "" || plscode == nil || status == "REQUEST_DENIED" || status == "ZERO_RESULTS" {
		return Response{FormatedAddress: deformat, Status: status}, err
	}

	return Response{FormatedAddress: compoundCode, Status: status}, nil
}

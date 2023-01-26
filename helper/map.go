package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Zenk41/sipencari-rest-api/util"
)

type Response struct {
	FormatedAddress string `json:"formated_address"`
	Status          string `json:"status"`
}

func GetAdressFromLatLng(lat float64, lng float64) (Response, error) {
	client := &http.Client{}
	APIKey := util.GetEnv("MAP_API_KEY")
	URL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f, %f&key=%s", lat, lng, APIKey)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	response := string(resBody)

	resBytes := []byte(response)
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	plsCode := jsonRes["plus_code"].(map[string]interface{})
	compoundCode := plsCode["compound_code"].(string)
	status := jsonRes["status"].(string)

	if compoundCode != "" || plsCode != nil {
		return Response{FormatedAddress: "Unknown Place Name", Status: status}, err
	}

	return Response{FormatedAddress: compoundCode, Status: status}, err
}

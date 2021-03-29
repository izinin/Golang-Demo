package main

import "encoding/json"

func Unmarshal(data []byte) (HousingPrices, error) {
	var r HousingPrices
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *HousingPrices) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

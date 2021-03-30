package main

import (
	"github.com/buger/jsonparser"
)

type HousingPrice struct {
	ID                              uint `gorm:"autoIncrement;primaryKey"`
	FldSid                          string
	FldId                           string
	FldPosition                     string
	FldCreatedAt                    string
	FldCreatedMeta                  string
	FldUpdatedAt                    string
	FldUpdatedMeta                  string
	FldMeta                         string
	ApplicationReceiptDate          string
	ApplicationCompletionDate       string
	LocationOfProperty              string
	ZipCode                         string
	F8HourTrainingProvider          string
	Neighborhood                    string
	Numofunits                      string
	Ethnicity                       string
	Hispanic_yOr_n                  string
	ApplicantHtfdRes                string
	Firsttimehomebuyer              string
	HouseholdSize                   string
	HouseHoldIncome_0_30Percent     string
	HouseHoldIncome_31_50Percent    string
	HouseHoldIncome_51_60Percent    string
	HouseHoldIncome_61_80Percent    string
	HouseHoldIncome_81PercentOrMore string
	PurchasePriceTotalDevCost       string
	DownPaymentAmount               string
	ClosingCostAmount               string
	TotalAmountOfHomeFundsComm      string
	ClosingDate1                    string
	ClosingDate2                    string
	ClosingDate3                    string
	Volume1                         string
	Page1                           string
	Status1                         string
	Volume2                         string
	Page2                           string
	Status2                         string
	Volume3                         string
	Page3                           string
	Status3                         string
	Fiscal_year                     string
	Location                        string
	Xcoord                          string
	Ycoord                          string
	Geom                            string
	FldComputedRegion_2vdc_22if     string
	FldComputedRegion_35zh_8fi2     string
	FldComputedRegion_ugzy_ysqh     string
	FldComputedRegionHaf6_6xye      string
}

func newHousingPrice(data []byte) *HousingPrice {
	fields := make([]string, 50)
	index := 0
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fields[index] = string(value)
		index++
	})
	return fromFieldArray(fields)
}

func fromFieldArray(values []string) *HousingPrice {
	return &HousingPrice{
		FldSid:                          values[0],
		FldId:                           values[1],
		FldPosition:                     values[2],
		FldCreatedAt:                    values[3],
		FldCreatedMeta:                  values[4],
		FldUpdatedAt:                    values[5],
		FldUpdatedMeta:                  values[6],
		FldMeta:                         values[7],
		ApplicationReceiptDate:          values[8],
		ApplicationCompletionDate:       values[9],
		LocationOfProperty:              values[10],
		ZipCode:                         values[11],
		F8HourTrainingProvider:          values[12],
		Neighborhood:                    values[13],
		Numofunits:                      values[14],
		Ethnicity:                       values[15],
		Hispanic_yOr_n:                  values[16],
		ApplicantHtfdRes:                values[17],
		Firsttimehomebuyer:              values[18],
		HouseholdSize:                   values[19],
		HouseHoldIncome_0_30Percent:     values[20],
		HouseHoldIncome_31_50Percent:    values[21],
		HouseHoldIncome_51_60Percent:    values[22],
		HouseHoldIncome_61_80Percent:    values[23],
		HouseHoldIncome_81PercentOrMore: values[24],
		PurchasePriceTotalDevCost:       values[25],
		DownPaymentAmount:               values[26],
		ClosingCostAmount:               values[27],
		TotalAmountOfHomeFundsComm:      values[28],
		ClosingDate1:                    values[29],
		ClosingDate2:                    values[30],
		ClosingDate3:                    values[31],
		Volume1:                         values[32],
		Page1:                           values[33],
		Status1:                         values[34],
		Volume2:                         values[35],
		Page2:                           values[36],
		Status2:                         values[37],
		Volume3:                         values[38],
		Page3:                           values[39],
		Status3:                         values[40],
		Fiscal_year:                     values[41],
		Location:                        values[42],
		Xcoord:                          values[43],
		Ycoord:                          values[44],
		Geom:                            values[45],
		FldComputedRegion_2vdc_22if:     values[46],
		FldComputedRegion_35zh_8fi2:     values[47],
		FldComputedRegion_ugzy_ysqh:     values[48],
		FldComputedRegionHaf6_6xye:      values[49],
	}
}

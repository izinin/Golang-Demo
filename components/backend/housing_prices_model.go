package main

import (
	"github.com/buger/jsonparser"
)

type HousingPrice struct {
	ID                                   uint
	fld_sid                              string
	fld_id                               string
	fld_position                         string
	fld_created_at                       string
	fld_created_meta                     string
	fld_updated_at                       string
	fld_updated_meta                     string
	fld_meta                             string
	application_receipt_date             string
	application_completion_date          string
	location_of_property                 string
	zip_code                             string
	f8_hour_training_provider            string
	neighborhood                         string
	numofunits                           string
	ethnicity                            string
	hispanic_y_or_n                      string
	applicant_htfd_res                   string
	firsttimehomebuyer                   string
	household_size                       string
	house_hold_income_0_30_percent       string
	house_hold_income_31_50_percent      string
	house_hold_income_51_60_percent      string
	house_hold_income_61_80_percent      string
	house_hold_income_81_percent_or_more string
	purchase_price_total_dev_cost        string
	down_payment_amount                  string
	closing_cost_amount                  string
	total_amount_of_home_funds_comm      string
	closing_date1                        string
	closing_date2                        string
	closing_date3                        string
	volume1                              string
	page1                                string
	status1                              string
	volume2                              string
	page2                                string
	status2                              string
	volume3                              string
	page3                                string
	status3                              string
	fiscal_year                          string
	location                             string
	xcoord                               string
	ycoord                               string
	geom                                 string
	fld_computed_region_2vdc_22if        string
	fld_computed_region_35zh_8fi2        string
	fld_computed_region_ugzy_ysqh        string
	fld_computed_region_haf6_6xye        string
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
		fld_sid:                              values[0],
		fld_id:                               values[1],
		fld_position:                         values[2],
		fld_created_at:                       values[3],
		fld_created_meta:                     values[4],
		fld_updated_at:                       values[5],
		fld_updated_meta:                     values[6],
		fld_meta:                             values[7],
		application_receipt_date:             values[8],
		application_completion_date:          values[9],
		location_of_property:                 values[10],
		zip_code:                             values[11],
		f8_hour_training_provider:            values[12],
		neighborhood:                         values[13],
		numofunits:                           values[14],
		ethnicity:                            values[15],
		hispanic_y_or_n:                      values[16],
		applicant_htfd_res:                   values[17],
		firsttimehomebuyer:                   values[18],
		household_size:                       values[19],
		house_hold_income_0_30_percent:       values[20],
		house_hold_income_31_50_percent:      values[21],
		house_hold_income_51_60_percent:      values[22],
		house_hold_income_61_80_percent:      values[23],
		house_hold_income_81_percent_or_more: values[24],
		purchase_price_total_dev_cost:        values[25],
		down_payment_amount:                  values[26],
		closing_cost_amount:                  values[27],
		total_amount_of_home_funds_comm:      values[28],
		closing_date1:                        values[29],
		closing_date2:                        values[30],
		closing_date3:                        values[31],
		volume1:                              values[32],
		page1:                                values[33],
		status1:                              values[34],
		volume2:                              values[35],
		page2:                                values[36],
		status2:                              values[37],
		volume3:                              values[38],
		page3:                                values[39],
		status3:                              values[40],
		fiscal_year:                          values[41],
		location:                             values[42],
		xcoord:                               values[43],
		ycoord:                               values[44],
		geom:                                 values[45],
		fld_computed_region_2vdc_22if:        values[46],
		fld_computed_region_35zh_8fi2:        values[47],
		fld_computed_region_ugzy_ysqh:        values[48],
		fld_computed_region_haf6_6xye:        values[49],
	}
}

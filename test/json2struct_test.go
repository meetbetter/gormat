/*
@Time : 2020/4/14 13:29
@Software: GoLand
@File : json2struct_test
@Author : Bingo <airplayx@gmail.com>
@Desc: https://github.com/airplayx/gormat/issues/2
*/
package test

import (
	"fmt"
	"github.com/airplayx/json2go"
	"log"
	"strings"
	"testing"
)

func Test_json2struct(t *testing.T) {
	bytes, err := json2go.New([]byte(`[{
    "input_index": 0,
    "candidate_index": 0,
    "delivery_line_1": "1 N Rosedale St",
    "last_line": "Baltimore MD 21229-3737",
    "delivery_point_barcode": "212293737013",
    "components": {
      "primary_number": "1",
      "street_predirection": "N",
      "street_name": "Rosedale",
      "street_suffix": "St",
      "city_name": "Baltimore",
      "state_abbreviation": "MD",
      "zipcode": "21229",
      "plus4_code": "3737",
      "delivery_point": "01",
      "delivery_point_check_digit": "3"
    },
    "metadata": {
      "record_type": "S",
      "zip_type": "Standard",
      "county_fips": "24510",
      "county_name": "Baltimore City",
      "carrier_route": "C047",
      "congressional_district": "07",
      "rdi": "Residential",
      "elot_sequence": "0059",
      "elot_sort": "A",
      "latitude": 39.28602,
      "longitude": -76.6689,
      "precision": "Zip9",
      "time_zone": "Eastern",
      "utc_offset": -5,
      "dst": true
    },
    "analysis": {
      "dpv_match_code": "Y",
      "dpv_footnotes": "AABB",
      "dpv_cmra": "N",
      "dpv_vacant": "N",
      "active": "Y"
    }
  },
  {
    "input_index": 0,
    "candidate_index": 1,
    "delivery_line_1": "1 S Rosedale St",
    "last_line": "Baltimore MD 21229-3739",
    "delivery_point_barcode": "212293739011",
    "components": {
      "primary_number": "1",
      "street_predirection": "S",
      "street_name": "Rosedale",
      "street_suffix": "St",
      "city_name": "Baltimore",
      "state_abbreviation": "MD",
      "zipcode": "21229",
      "plus4_code": "3739",
      "delivery_point": "01",
      "delivery_point_check_digit": "1"
    },
    "metadata": {
      "record_type": "S",
      "zip_type": "Standard",
      "county_fips": "24510",
      "county_name": "Baltimore City",
      "carrier_route": "C047",
      "congressional_district": "07",
      "rdi": "Residential",
      "elot_sequence": "0064",
      "elot_sort": "A",
      "latitude": 39.2858,
      "longitude": -76.66889,
      "precision": "Zip9",
      "time_zone": "Eastern",
      "utc_offset": -5,
      "dst": true
    },
    "analysis": {
      "dpv_match_code": "Y",
      "dpv_footnotes": "AABB",
      "dpv_cmra": "N",
      "dpv_vacant": "N",
      "active": "Y"
    }
  }]`), "YourStruct")
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err := bytes.WriteGo()
	fmt.Println(strings.ReplaceAll(string(b), "\t", "    "), err)
}

func Test_json2struct_more(t *testing.T) {
	bytes, err := json2go.New([]byte(`[
		[
			{
				"id": 123
			},
			[
				{
					"test": false,
					"hello": "abc",
					"msg": []
				}
			]
		]
	]`), "YourStruct")
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err := bytes.WriteGo()
	fmt.Println(strings.ReplaceAll(string(b), "\t", "    "), err)
}

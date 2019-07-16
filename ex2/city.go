package main

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

type district struct {
	Id string `yaml:"id"`
	Name string `yaml:"name"`
	Wards []*ward `yaml:"wards"`
}
type ward struct {
	Id string `yaml:"id"`
	Name string `yaml:"name"`
}
type city struct {
	Id string `yaml:"id"`
	Name string `yaml:"name"`
	Districts []*district `yaml:"districts"`
}
var cities []*city

func main() {
	filename := "data.csv"

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		if _, err := strconv.Atoi(line[1]); err != nil {
			continue
		}
		var cty *city
		for _, c := range cities {
			if c.Id == line[1] {
				cty = c
				break
			}
		}
		if cty == nil {
			cty = &city{line[1], line[0], make([]*district,0)}
			cities = append(cities, cty)
		}
		var dis *district
		for _, c := range cty.Districts {
			if c.Id == line[3] {
				dis = c
				break
			}
		}
		if dis == nil {
			dis = &district{line[3], line[2], make([]*ward,0)}
			cty.Districts = append(cty.Districts, dis)
		}
		var war *ward
		for _, c := range dis.Wards {
			if c.Id == line[5] {
				war = c
				break
			}
		}
		if war == nil {
			war = &ward{line[5], line[4]}
			dis.Wards = append(dis.Wards, war)
		}
	}

	res, err := yaml.Marshal(cities)
	if err != nil {
		return
	}

	ioutil.WriteFile("data.yml", res, 0644)

	for _, c := range cities {
		fmt.Printf("- id: %s\n", c.Id)
		fmt.Printf("  name: %s\n", c.Name)
	}

}

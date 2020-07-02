package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

// const mul float64 = 1.6
// const bitMax float64 = 65535

// type car struct {
// 	gas_pedal      uint16
// 	brake_pedal    uint16
// 	steering_wheel int16
// 	top_speed      float64
// }

// func Add(a float64, b float64) float64 {
// 	return a + b
// }

// // a method of car struct
// func (c car) kmh() float64 {
// 	return float64(c.gas_pedal) * (c.top_speed / bitMax) * mul
// }

// func (c car) mph() float64 {
// 	return float64(c.gas_pedal) * (c.top_speed / bitMax)
// }

// func (c *car) new_top_speed(newspeed float64) {
// 	c.top_speed = newspeed
// }

// func index_handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, This is Index")
// }

// func about_handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
// 		<h1>title</h1>
// 		<p>Paragraph</p>
// 	`)
// }


// type SiteMapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

type Location struct {
	Loc string `xml:"loc"`
}

// // type modifier ? overdrive string return
// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyworld string
	Location string
}

func main() {
	// fmt.Println(Add(1, 2))
	// SamePackageFunc()
	// examplePackage.TestPack()
	// examplePackage.TestOther()

	// a_car := car{gas_pedal: 12, brake_pedal: 0, steering_wheel: 1024, top_speed: 200}
	// fmt.Println(a_car)
	// fmt.Println(a_car.kmh())
	// fmt.Println(a_car.mph())

	// a_car.new_top_speed(5000)
	// fmt.Println(a_car.kmh())
	// fmt.Println(a_car.mph())

	// http.HandleFunc("/", index_handle)
	// http.HandleFunc("/about/", about_handle)

	// http.ListenAndServe(":8000", nil)
	var s SiteMapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	resp.Body.Close()
	// fmt.Println(string_body)
	
	xml.Unmarshal(bytes,&s)
	// fmt.Println(s.Locations)

	// for i := 0;i<10;i++{
	// 	fmt.Println(i)
	// }

	// the same as for location in s.Location in python
	for _, Location := range s.Locations {
		// fmt.Printf("%s\n",Location)

		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes,&n)

		for idx,_ := range n.Titles{
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx],n.Locations[idx]}
		}

	}

	for idx, data := range news_map{
		fmt.Println("\n\n",idx)
		fmt.Println("\n",data.Keyworld)
		fmt.Println("\n",data.Location)
	}

	//map with key of string and value of float
	//var grades map[string]float32
	//grades := map[string]float32
	// grades := make(map[string]float32)

	// grades["John"] = 43
	// grades["Smith"] = 12
	// grades["Wick"] = 69

	
	// JohnGrade := grades["John"]
	// fmt.Println(JohnGrade)

	// // delete(grades,"Wick")
	// // fmt.Println(grades)

	// for key,value := range grades{
	// 	fmt.Println(key,":",value)
	// }

}

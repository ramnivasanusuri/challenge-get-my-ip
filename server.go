package main

import (
	"encoding/csv"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// <-----------------------------------Golbals------------------------------------------------->

var path string = "C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/input_fqdn.csv"
var FQDN [][]string
var wg sync.WaitGroup
var mut sync.RWMutex

// <------------------------------------------------------------------------------------------->

func checkError(e error) {
	if e != nil {
		fmt.Println("Error occured:", e)
	}
}

// func checkIPError(e error) {
// 	if e != nil {
// 		fmt.Printf("Could not get IP: %v\n", e)
// 	}
// }

func isIPv4(s string) bool {
	return strings.Count(s, ":") < 2
}

// <------------------------------------------------------------------------------------------->

func writeToCSV(a [][]string) {
	f, err := os.Create("C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/sampleop.csv")
	checkError(err)

	newWriter := csv.NewWriter(f)
	err = newWriter.WriteAll(a)
	checkError(err)

}

func wToCSV(v []string){
	defer wg.Done()
	f, err := os.Create("C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/sampleop.csv")
	checkError(err)
	mut.Lock()
	newWriter := csv.NewWriter(f)
	mut.Unlock()
	err = newWriter.Write(v)
	checkError(err)
}

func lookup_IP(v []string) {
	defer wg.Done()
	if v[0] == "FQDN" {
		return
	}

	var fl []string
	fl = append(fl, v[0])

	ip, err := net.LookupHost(v[0])
	if err != nil {
		// mut.Lock()
		fl = append(fl, err.Error())
		FQDN = append(FQDN, fl)
		// mut.Unlock()
		return
	}

	for _, i := range ip {
		if isIPv4(i) {
			ip[0] = i
		}
	}
	// mut.Lock()
	fl = append(fl, ip[0])
	FQDN = append(FQDN, fl)
	// mut.Unlock()
}

func main() {
	// Your Code Goes here !!

	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	var desline = []string{"FQDN", "IP"}
	FQDN = append(FQDN, desline)

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	// fmt.Print(data)
	for _, v := range data {

		wg.Add(1)
		lookup_IP(v)

	}
	fmt.Println(FQDN)
	wg.Wait()

	writeToCSV(FQDN)
	// for _,v:=range FQDN{
	// 	wg.Add(1)
	// 	go wToCSV(v)
	// }
	// wg.Wait()
}

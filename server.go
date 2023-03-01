package main

import (
	"encoding/csv"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

var path string = "C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/input_fqdn.csv"
var FQDN [][]string

var wg sync.WaitGroup
var mut sync.RWMutex

func checkError(e error) {
	if e != nil {
		fmt.Println("Error occured:", e)
	}
}

func isIPv4(s string) bool {
	return strings.Count(s, ":") < 2
}

func writeToCSV(a [][]string) {
	f, err := os.Create("C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/sampleop.csv")
	checkError(err)

	newWriter := csv.NewWriter(f)
	err = newWriter.WriteAll(a)
	checkError(err)

}

// func regexp_IPv4(s string){
// 	regexp.MatchString("[0-255]:[0-255]:")
// }

func lookupIP(ch chan string) {
	defer wg.Done()
	v := <-ch
	// n := <-nch
	var fl []string
	// fl = append(fl, n)
	fl = append(fl, v)

	ip, err := net.LookupHost(v)
	if err != nil {
		mut.RLock()
		fl = append(fl, "Invalid domain.")
		FQDN = append(FQDN, fl)
		mut.RUnlock()
		return
	}

	for _, i := range ip {
		if isIPv4(i) {
			ip[0] = i
			break
		}
	}
	mut.RLock()
	fl = append(fl, ip[0])
	FQDN = append(FQDN, fl)
	mut.RUnlock()
}

func main() {
	// Your Code Goes here !!
	// var f_data []string

	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	var desline = []string{"FQDN", "IP"}
	FQDN = append(FQDN, desline)

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	checkError(err)

	ch := make(chan string, len(data))
	// nch := make(chan string, len(data))

	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range data {
		if _, value := keys[entry[0]]; !value {
			keys[entry[0]] = true
			list = append(list, entry[0])
		}
	}
	fmt.Printf("Unique domains:%v", len(list))
	fmt.Println(list)

	for _, v := range list {
		if v == "FQDN" {
			continue
		}
		wg.Add(1)
		ch <- v
		// nch <- v[0]
		go lookupIP(ch)
	}
	wg.Wait()

	writeToCSV(FQDN)

	http.HandleFunc("/c12", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("C:/Users/ranusu737/Desktop/J Project/challenge-get-my-ip/sampleop.csv")
		checkError(err)
		defer file.Close()

		csvReader := csv.NewReader(file)
		data, err := csvReader.ReadAll()
		checkError(err)

		writer := csv.NewWriter(w)

		// Write the data to the CSV file
		for _, value := range data {
			err := writer.Write(value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		writer.Flush()
	})

	http.ListenAndServe(":8090", nil)
}

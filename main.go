package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var inputdatas []string

var inputdatafinal []string

func indexHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello User!") // write data to response

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		getdata := r.Form["name"]
		//===============process data input========
		process(getdata)

	}
}

func outputHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "The count is: %s/n", inputdatas)
	fmt.Fprintf(w, "The count is: %s", inputdatafinal)

}

func process(getdata []string) {

	str, err := json.Marshal(getdata)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	userinput := string(str[:])

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	userinput = re.ReplaceAllString(userinput, " ")
	userinput = strings.ToLower(userinput)

	fmt.Println(userinput)

	for word, occur := range countsimilarword(userinput) {

		occurance := strconv.Itoa(occur)

		inputdatas = append(inputdatas, "word:", word, "occurs", occurance, "times,")

	}

}

// ======================count similar word===========
func countsimilarword(st string) map[string]int {

	words := strings.Fields(st)

	m := make(map[string]int)
	for _, word := range words {
		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	processn(m)

	return m

}

func processn(m map[string]int) {
	words := make([]string, 0, len(m))
	for word := range m {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return m[words[i]] > m[words[j]]
	})

	for _, word := range words {
		//for i := 0; i < 10; i++ {
		fmt.Printf("%v %v\n", word, m[word])
		//inputdatas = append(inputdatas, "word:", word, "occurs", m[word], "times,")
		num := strconv.Itoa(m[word])
		inputdatafinal = append(inputdatafinal, word, num)

		//}
	}

	for i:=0; i<len()

    
}

func main() {

	http.HandleFunc("/index", indexHandler) // welcome page

	http.HandleFunc("/input", inputHandler)
	http.HandleFunc("/output", outputHandler)

	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println(inputdatas)
}

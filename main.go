package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/feeds"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	timeFormat = "02.01.2006"
	html       = "template.html"
	data       = "conf.yml"
	daysBefore = 10
)

type Conf struct {
	Title     string
	URL       string
	Startdate string
	CFPDate   string
	Location  string
}

func mkHTML(cnf *[]Conf) {
	//f, _ := filepath.Abs(html)
	htmltmpl, err := ioutil.ReadFile(html)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("tmpl").Parse(string(htmltmpl)))
	t.Execute(os.Stdout, *cnf)
}

func mkRSS(cnf *[]Conf) {

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "opensource events",
		Link:        &feeds.Link{Href: "https://bronevichok.ru/oss"},
		Description: "opensource, peace, software",
		Author:      &feeds.Author{Name: "Sergey Bronnikov", Email: "sergeyb@openvz.org"},
		Created:     now,
	}

	conf := feeds.Item{}
	for _, c := range *cnf {
		Link := &feeds.Link{Href: c.URL}
		Author := &feeds.Author{Name: "Sergey Bronnikov", Email: "sergeyb@openvz.org"}
		conf.Title = c.Title
		conf.Link = Link
		conf.Description = ""
		conf.Author = Author
		conf.Created = now
		feed.Add(&conf)
	}

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rss)
}

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:\n")
		//flag.PrintDefaults()
	}

	format := flag.String("out", "", "Output format (rss, html, icalendar)")
	flag.Parse()

	if *format == "" {
		fmt.Println("No parameters specified.")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(data); os.IsNotExist(err) {
		fmt.Printf("File %v doesn't exist.\n", data)
		os.Exit(1)
	}

	f, _ := filepath.Abs(data)
	yamlFile, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	confs := []Conf{}
	closestConfs := []Conf{} // conferences started after beforeDays
	closestCFPs := []Conf{}  // conferences which will finish CFP after beforeDays

	err = yaml.Unmarshal(yamlFile, &confs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	now := time.Now()
	for _, c := range confs {
		if c.Startdate == now.AddDate(0, 0, daysBefore).Format(timeFormat) {
			closestConfs = append(closestConfs, c)
		}
		if c.CFPDate == now.AddDate(0, 0, daysBefore).Format(timeFormat) {
			closestCFPs = append(closestCFPs, c)
		}
	}

	if *format == "rss" {
		mkRSS(&closestConfs)
	} else if *format == "html" {
		mkHTML(&closestConfs)
	} else if *format == "icalendar" {
		// mkiCal(&closestConfs)
		fmt.Printf("Not implemented.")
	}
}

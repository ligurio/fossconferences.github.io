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
	author     = "Sergey Bronnikov"
	email      = "sergeyb@openvz.org"
	url        = "https://bronevichok.ru/ose"
	daysBefore = 10
)

type Conf struct {
	Title     string
	URL       string
	Startdate string
	CFPDate   string
	CFPURL    string
	Location  string
}

func mkHTML(cnf *[]Conf) {
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
		Link:        &feeds.Link{Href: url},
		Description: "opensource, peace, software",
		Author:      &feeds.Author{Name: author, Email: email},
		Created:     now,
	}

	for _, c := range *cnf {
		conf := feeds.Item{}
		Link := &feeds.Link{Href: c.URL}
		Author := &feeds.Author{Name: author, Email: email}
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
		flag.PrintDefaults()
		fmt.Println()
	}

	format := flag.String("out", "", "Output format (rss, html)")
	flag.Parse()

	if *format == "" {
		fmt.Println("No parameters specified.")
		flag.Usage()
		//os.Exit(1)
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
		if c.Startdate != "none" {
			conftime, _ := time.Parse(timeFormat, c.Startdate)
			conf := int64(conftime.Sub(now).Hours()) / 24

			if (conf <= daysBefore) && (conf > 0) {
				closestConfs = append(closestConfs, c)
			}
			if now.Year() > conftime.Year() && *format == "" {
				fmt.Printf("WARNING: This conference was last time in previous year: %s - %s\n", c.Title, c.URL)
			}
		}
		/*
			if c.CFPDate == "none" && c.Startdate == "none" {
				fmt.Printf("WARNING: CFP and start dates are empty: %s - %s\n", c.Title, c.URL)
			}
		*/

		if c.CFPDate != "none" {
			cfptime, _ := time.Parse(timeFormat, c.CFPDate)
			cfp := int64(cfptime.Sub(now).Hours() / 24)

			if (cfp <= daysBefore) && (cfp > 0) {
				closestCFPs = append(closestCFPs, c)
			}
		} else {
			fmt.Printf("WARNING: CFP date is empty: %s - %s\n", c.Title, c.URL)
		}
	}

	if *format == "rss" {
		mkRSS(&closestConfs)
	} else if *format == "html" {
		mkHTML(&closestConfs)
	}
}

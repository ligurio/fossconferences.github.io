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

func mkFeed(cnf *[]Conf, f string) {

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
		conf.Title = c.Title
		if c.URL != "none" {
			conf.Link = &feeds.Link{Href: c.URL}
		}
		if c.CFPURL != "none" {
			conf.Description = c.CFPURL
		}
		conf.Author = &feeds.Author{Name: author, Email: email}
		conf.Created = now
		feed.Add(&conf)
	}

	if f == "rss" {
		lenta, _ := feed.ToRss()
		fmt.Println(lenta)
	} else {
		lenta, _ := feed.ToAtom()
		fmt.Println(lenta)
	}

}

func main() {

	flag.Usage = func() {
		fmt.Println("Usage:\n")
		flag.PrintDefaults()
		fmt.Println()
	}

	format := flag.String("out", "", "Output format (rss, atom, html)")
	flag.Parse()

	if *format == "" {
		fmt.Println("No parameters specified.")
		flag.Usage()
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
	closestConfs := []Conf{}

	err = yaml.Unmarshal(yamlFile, &confs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	now := time.Now()
	for _, c := range confs {
		// notify if 5 or 10 days before ending of CFP
		// warn if not: (startdate was this year) or (CFPdate or startdate will be in future)
		message := ""
		if c.Startdate != "none" {
			conftime, _ := time.Parse(timeFormat, c.Startdate)
			conf := int64(conftime.Sub(now).Hours()) / 24

			if now.Year() > conftime.Year() && *format == "" {
				fmt.Printf("WARNING: This conference was last time in previous year: %s - %s\n", c.Title, c.URL)
			}
			if (conf <= daysBefore) && (conf > 0) {
				message = "Conference will start soon: "
				closestConfs = append(closestConfs, c)
			}
		}

		if c.CFPDate != "none" {
			cfptime, _ := time.Parse(timeFormat, c.CFPDate)
			cfp := int64(cfptime.Sub(now).Hours() / 24)

			if (cfp <= daysBefore) && (cfp > 0) {
				message = "Conference will start soon: "
				closestConfs = append(closestConfs, c)
			}
			if now.Year() > cfptime.Year() && *format == "" {
				fmt.Printf("WARNING: CFP of this conference was last time in previous year: %s - %s\n", c.Title, c.URL)
			}
		} else {
			if *format == "" {
				fmt.Printf("WARNING: CFP date is empty: %s - %s\n", c.Title, c.URL)
			}
		}
		c.Title = message + c.Title
	}

	if *format == "rss" || *format == "atom" {
		mkFeed(&closestConfs, *format)
	} else if *format == "html" {
		mkHTML(&closestConfs)
	}
}

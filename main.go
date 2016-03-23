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

func mkCSV(cnf *[]Conf) {
	fmt.Printf("Title;URL;Begin date;CFP;CFP URL;Location\n")
	for _, c := range *cnf {
		fmt.Printf("%s;%s;%s;%s;%s;%s\n", c.Title, c.URL, c.Startdate, c.CFPDate, c.CFPURL, c.Location)
	}
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

func wasThisYear(date string) bool {
	now := time.Now()
	if date != "none" {
		ctime, _ := time.Parse(timeFormat, date)
		if (ctime.Year() == now.Year()) && int64(ctime.Sub(now)) < 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
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
		if c.CFPDate != "none" {
			cfptime, err := time.Parse(timeFormat, c.CFPDate)
			if err != nil && *format == "" {
				fmt.Printf("[WARN] Wrong date specified (%s): %s - %s\n", c.CFPDate, c.Title, c.URL)
			}
			cfp := int64(cfptime.Sub(now).Hours() / 24)

			if cfp == 5 || cfp == 10 {
				c.Title = "CFP will finish soon: " + c.Title
				closestConfs = append(closestConfs, c)
			}
		} else {
			if *format == "" && !wasThisYear(c.Startdate) {
				fmt.Printf("[WARN] CFP date is empty: %s - %s\n", c.Title, c.URL)
			}
		}
	}

	if *format == "rss" || *format == "atom" {
		mkFeed(&closestConfs, *format)
	} else if *format == "html" {
		mkHTML(&closestConfs)
	} else if *format == "csv" {
		mkCSV(&confs)
	}
}

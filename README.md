# FOSS events tracking

**[Subscribe to RSS feed](https://bronevichok.ru/ose/conf-rss.xml), [review page](https://bronevichok.ru/ose/)**

Crafted by [Sergey Bronnikov](https://bronevichok.ru/).

---

[![Build Status](https://travis-ci.org/ligurio/oss-events.svg?branch=master)](https://travis-ci.org/ligurio/oss-events)

## Sources about OS release dates

* CentOS: [Releases](https://wiki.centos.org/About/Product)
* Debian: [Releases](https://wiki.debian.org/DebianReleases)
* [DistroWatch](http://distrowatch.com/weekly.php?issue=20150727#upcoming)
* Fedora: [Releases](https://fedoraproject.org/wiki/Releases), [EOL](https://fedoraproject.org/wiki/End_of_life)
* FreeBSD: [Releases](), [Release Schedule](https://www.freebsd.org/releng/)
* OpenSUSE: [Releases](https://en.opensuse.org/openSUSE:Roadmap), [Lifetime](https://en.opensuse.org/Lifetime)
* RHEL: [Releases](https://access.redhat.com/articles/3078)
* Ubuntu: [Releases](https://wiki.ubuntu.com/Releases), [Development Codenames](https://wiki.ubuntu.com/DevelopmentCodeNames)


## Sources about opensource conferences

* [Linux Foundation](http://events.linuxfoundation.org/)
* [LWN](https://lwn.net/Calendar/)
* [OpenSource.com](https://opensource.com/resources/conferences-and-events-monthly)
* Twitter lists: [@estet](https://twitter.com/estet/lists/foss-conferences), [@jakerella](https://twitter.com/jakerella/lists/conferences/members)
* [Call to speakers](https://calltospeakers.com/)
* [Lanyrd](http://lanyrd.com/calls/)
* https://github.com/bamos/conference-tracker
* [BSD events](http://www.bsdevents.org/)
* [Wikipedia](https://en.wikipedia.org/wiki/List_of_computer_science_conferences)
* [RichardLitt/awesome-conferences](https://github.com/RichardLitt/awesome-conferences)
* [PlanetRuby/awesome-events](https://github.com/planetruby/awesome-events)
* [WikiCFP: CS](http://www.wikicfp.com/cfp/call?conference=computer%20science)

## Build

``
export GO386=387; export GOARCH=386; export GOOS=openbsd; go build -o oss
``

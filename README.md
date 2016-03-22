# FOSS events tracking

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

* https://en.wikipedia.org/wiki/List_of_free-software_events
* http://events.linuxfoundation.org/
* https://lwn.net/Calendar/
* http://www.wikicfp.com/cfp/home
* https://opensource.com/resources/conferences-and-events-monthly
* https://twitter.com/estet/lists/foss-conferences
* https://github.com/bamos/conference-tracker
* http://www.bsdevents.org/

## Build

   export GO386=387; GOARCH=386; export GOOS=openbsd; go build -o oss

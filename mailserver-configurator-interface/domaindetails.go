package main

import (
	"github.com/go-chi/chi"
	"gopkg.in/unrolled/render.v1"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"log"
)

func getDomainDetails(w http.ResponseWriter, r *http.Request)  {
	details := newDomainDetails(chi.URLParam(r, "domain"))
	ren := render.New()
	ren.JSON(w, http.StatusOK, details)
}

func newDomainDetails (domainName string) DomainDetails {
	d := DomainDetails{}
	//defaults
	d.MXRecordCheck = false
	d.SPFRecordCheck = false
	d.DMARCRecordCheck = false
	d.RecordChecked = true
	d.DKIMCheck = false

	//varieables
	d.DomainName = domainName

	//methods
	d.readPostfixConfig()
	d.checkMXRecord()
	d.checkSPFRecord()
	d.checkDMARCRecord()
	d.checkDKIMCRecord()

	return d
}

type DomainDetails struct {
	DomainName string `json:"domain_name"`

	hostname string
	MXRecordCheck bool
	SPFRecordCheck bool
	DMARCRecordCheck bool
	RecordChecked bool
	DKIMCheck bool
}

func (d *DomainDetails) readPostfixConfig() {
	//Get Hostname vom postfix config
	dat, err := ioutil.ReadFile("/etc/postfix/main.cf")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`myhostname\s?=.*`)
	hostnameline := re.FindAllString(string(dat), -1)[0]
	hostnamearray := strings.Split(hostnameline, "=")
	hostname := strings.Trim(hostnamearray[1], " ")
	log.Println("Hostname: "+hostname)
	d.hostname = hostname

}

func (d *DomainDetails) checkMXRecord() {
	log.Println("Check MX Record")
	mxrecords, _ := net.LookupMX(d.DomainName)
	for _, mx := range mxrecords {
		if(mx.Host == d.hostname+".") {
			log.Println("Found MX valide Record for Domain "+d.DomainName)
			d.MXRecordCheck = true
		}
	}
}

func (d *DomainDetails) checkSPFRecord() {
	log.Println("Check SPF Record")
	rs, err := net.LookupTXT(d.DomainName)
	if err != nil {
		log.Println("SPF Record check failed")
		return
		//panic(err)
	}
	for _, record := range rs {
		if record == "v=spf1 a:"+d.hostname+" ?all" {
			d.SPFRecordCheck = true
		}
	}
}

func (d *DomainDetails) checkDMARCRecord() {
	log.Println("Check DMARC Record")
	rs, err := net.LookupTXT("_dmarc."+d.DomainName)
	if err != nil {
		log.Println("DMARC Record check failed")
		return
	}

	for _, record := range rs {
		if record == "v=DMARC1; p=reject;" {
			d.DMARCRecordCheck = true
		}
	}
}

func (d *DomainDetails) checkDKIMCRecord() {
	log.Println("Check DKMI Record")

	rs, err := net.LookupTXT(getConfigVariable("DKIM_SELECTOR")+"._domainkey."+d.DomainName)
	if err != nil {
		log.Println("DMARC Record check failed")
		return
	}

	for _, record := range rs {
		if record == getConfigVariable("DKIM_VALUE") {
			d.DKIMCheck = true
		}
	}
}

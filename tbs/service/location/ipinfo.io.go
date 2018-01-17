package location

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var IPINFO_CHAN = make(chan IpInfo)

func getIPInfoIO(ip string) {
	var url = "http://ipinfo.io/" + ip + "/json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var i IpInfo
	json.Unmarshal(body, &i)

	IPINFO_CHAN <- i
}

type IpInfo struct {
	IP       string `bson:"ip" json:"ip"`
	Hostname string `bson:"hostname" json:"hostname"`
	City     string `bson:"city" json:"city"`
	Region   string `bson:"region" json:"region"`
	Country  string `bson:"country" json:"country"`
	Loc      string `bson:"loc" json:"loc"`
	Org      string `bson:"org" json:"org"`
}

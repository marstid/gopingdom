package gopingdom

type ChecksResult struct {
	Checks []Check `json:"checks"`
	Counts struct {
		Total    int `json:"total"`
		Limited  int `json:"limited"`
		Filtered int `json:"filtered"`
	} `json:"counts"`
}

type CheckResult struct {
	Checks Check `json:"check"`
}


type Check struct {
	Type struct {
		HTTP struct {
			Username         string `json:"username"`
			Password         string `json:"password"`
			URL              string `json:"url"`
			Encryption       bool   `json:"encryption"`
			Port             int    `json:"port"`
			Shouldcontain    string `json:"shouldcontain"`
			Shouldnotcontain string `json:"shouldnotcontain"`
			Postdata         string `json:"postdata"`
			Requestheaders   struct {
			} `json:"requestheaders"`
		} `json:"http"`
	} `json:"type"`
	Sendnotificationwhendown int      `json:"sendnotificationwhendown"`
	Notifyagainevery         int      `json:"notifyagainevery"`
	Notifywhenbackup         bool     `json:"notifywhenbackup"`
	ResponsetimeThreshold    int      `json:"responsetime_threshold"`
	CustomMessage            string   `json:"custom_message"`
	Integrationids           []int    `json:"integrationids"`
	ID                       int      `json:"id"`
	Name                     string   `json:"name"`
	Lasterrortime            int      `json:"lasterrortime"`
	Lasttesttime             int      `json:"lasttesttime"`
	Lastresponsetime         int      `json:"lastresponsetime"`
	Status                   string   `json:"status"`
	Resolution               int      `json:"resolution"`
	Hostname                 string   `json:"hostname"`
	Created                  int      `json:"created"`
	Tags                     []Tag    `json:"tags"`
	ProbeFilters             []string `json:"probe_filters"`
	Ipv6                     bool     `json:"ipv6"`
	VerifyCertificate        bool     `json:"verify_certificate"`
	SslDownDaysBefore        int      `json:"ssl_down_days_before"`
}

type Check1 struct {
	ID                int      `json:"id"`
	Created           int      `json:"created"`
	Name              string   `json:"name"`
	Hostname          string   `json:"hostname"`
	Resolution        int      `json:"resolution"`
	Type              string   `json:"type"`
	Ipv6              bool     `json:"ipv6"`
	VerifyCertificate bool     `json:"verify_certificate"`
	Lasterrortime     int      `json:"lasterrortime,omitempty"`
	Lasttesttime      int      `json:"lasttesttime"`
	Lastresponsetime  int      `json:"lastresponsetime"`
	Status            string   `json:"status"`
	Tags              []Tag    `json:"tags"`
	Maintenanceids    []string `json:"maintenanceids,omitempty"`
}

type Tag struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type ProbesResult struct {
	Probes []Probe `json:"probes"`
}

type Probe struct {
	ID         int    `json:"id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	Hostname   string `json:"hostname"`
	IP         string `json:"ip"`
	Countryiso string `json:"countryiso"`
	Region     string `json:"region"`
	Ipv6       string `json:"ipv6,omitempty"`
}

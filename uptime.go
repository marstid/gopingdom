package gopingdom

import (
	"encoding/json"
	"fmt"
)

func (r *RestClient) UptimeGetChecks() (cks []Check, error error) {
	uri := "/checks"

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	if r.debug{
		fmt.Printf("Response: %s", data)
	}

	var result ChecksResult
	json.Unmarshal(data, &result)
	return result.Checks, nil
}

func (r *RestClient) UptimeGetCheckDetails(id int) (cks Check, error error) {
	uri := "/checks/" + fmt.Sprintf("%d",id)

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	if r.debug{
		fmt.Printf("Response: %s", data)
	}

	var result CheckResult
	json.Unmarshal(data, &result)
	return result.Checks, nil
}

func (r *RestClient) UptimeGetProbes() (prb []Probe, err error) {
	uri := "/probes"

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	var result ProbesResult
	json.Unmarshal(data, &result)
	return result.Probes, nil
}

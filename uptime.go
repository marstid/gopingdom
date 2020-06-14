package gopingdom

import (
	"encoding/json"
	"fmt"
	"sort"
)

func (r *RestClient) UptimeGetChecks() (cks Checks, error error) {
	uri := "/checks"

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	if r.debug {
		fmt.Printf("Response: %s", data)
	}

	var result ChecksResult
	json.Unmarshal(data, &result)
	sort.Sort(result.Checks)
	return result.Checks, nil
}

func (r *RestClient) UptimeGetCheckDetails(id int) (cks Check, error error) {
	uri := "/checks/" + fmt.Sprintf("%d", id)

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	if r.debug {
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
		return nil, err
	}

	var result ProbesResult
	json.Unmarshal(data, &result)
	return result.Probes, nil
}

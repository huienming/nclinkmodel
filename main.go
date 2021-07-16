package main

import ( //"html/template"

	"encoding/json"
	"fmt"
	nclink "nclink/model"
	"time"
)

const VERSION = "1.0.0"

func main() {
	register := nclink.RegisterRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Device: "test"}
	JRegister, err := json.Marshal(register)
	if err == nil {
		fmt.Println(string(JRegister))
	}
	discoveryRequest := nclink.DiscoveryRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}}
	JdiscoveryRequest, err := json.Marshal(discoveryRequest)
	if err == nil {
		fmt.Println(string(JdiscoveryRequest))
	}
	discoveryResponse := nclink.DiscoveryResponse{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Deviceids: []string{"test", "test1"}}
	JdiscoveryResponse, err := json.Marshal(discoveryResponse)
	if err == nil {
		fmt.Println(string(JdiscoveryResponse))
	}
	probeVersion := nclink.ProbeVersion{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Version: "1.0.0"}
	JprobeVersion, err := json.Marshal(probeVersion)
	if err == nil {
		fmt.Println(string(JprobeVersion))
	}
	probeQueryRequest := nclink.ProbeQueryRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}}
	JprobeQueryRequest, err := json.Marshal(probeQueryRequest)
	if err == nil {
		fmt.Println(string(JprobeQueryRequest))
	}
	probeQueryResponse := nclink.ProbeQueryResponse{BaseCommand: nclink.BaseCommand{CommandID: "0", Code: "OK"}, Probe: string(nclink.DefaultModel)}
	JprobeQueryResponse, err := json.Marshal(probeQueryResponse)
	if err == nil {
		fmt.Println(string(JprobeQueryResponse))
	}
	probeSetRequest := nclink.ProbeSetRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Probe: string(nclink.DefaultModel)}
	JprobeSetRequest, err := json.Marshal(probeSetRequest)
	if err == nil {
		fmt.Println(string(JprobeSetRequest))
	}
	probeSetResponse := nclink.ProbeSetResponse{BaseCommand: nclink.BaseCommand{CommandID: "0", Code: "OK"}}
	JprobeSetResponse, err := json.Marshal(probeSetResponse)
	if err == nil {
		fmt.Println(string(JprobeSetResponse))
	}
	offset := 10
	length := 20
	queryRequest := nclink.QueryRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Ids: []nclink.QueryItem{{CommandItem: nclink.CommandItem{ID: "000123"}, Params: &nclink.QueryParameter{CommandParameter: nclink.CommandParameter{Operation: "get-value", Offset: &offset, Length: &length}, Indexes: []string{"10", "1-8", "100"}}}, {CommandItem: nclink.CommandItem{ID: "000123456"}}}}

	JqueryRequest, err := json.Marshal(queryRequest)
	if err == nil {
		fmt.Println(string(JqueryRequest))
	}
	queryResponse := nclink.QueryResponse{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Values: []nclink.QueryItem{{CommandItem: nclink.CommandItem{BaseCommand: nclink.BaseCommand{Code: "OK"}, ID: "000123"}, Params: queryRequest.Ids[0].Params, Values: []interface{}{1, 2, 3, 4, 5, 67, 8}}, {CommandItem: nclink.CommandItem{BaseCommand: nclink.BaseCommand{Code: "OK"}, ID: "000123456"}, Values: []interface{}{"1000"}}}}
	JqueryResponse, err := json.Marshal(queryResponse)
	if err == nil {
		fmt.Println(string(JqueryResponse))
	}
	index := 10
	setReqeust := nclink.SetRequest{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Values: []nclink.SetItem{{CommandItem: nclink.CommandItem{ID: "000123"}, Params: nclink.SetParameter{CommandParameter: nclink.CommandParameter{Operation: "set-value", Offset: &offset, Length: &length}, Index: &index, Value: 10}}, {CommandItem: nclink.CommandItem{ID: "000123456"}, Params: nclink.SetParameter{CommandParameter: nclink.CommandParameter{Operation: "set-value"}, Value: 10}}}}
	JsetReqeust, err := json.Marshal(setReqeust)
	if err == nil {
		fmt.Println(string(JsetReqeust))
	}
	setResponse := nclink.SetResponse{BaseCommand: nclink.BaseCommand{CommandID: "0"}, Results: []nclink.SetItem{{CommandItem: nclink.CommandItem{BaseCommand: nclink.BaseCommand{Code: "OK"}, ID: "000123"}, Params: setReqeust.Values[0].Params},

		{CommandItem: nclink.CommandItem{BaseCommand: nclink.BaseCommand{Code: "OK"}, ID: "000123456"}, Params: setReqeust.Values[1].Params}}}
	JsetResponse, err := json.Marshal(setResponse)
	if err == nil {
		fmt.Println(string(JsetResponse))
	}
	sampledata := nclink.SampleData{ID: "00210030", BeginTime: time.Now().Unix(), Data: []nclink.SampleDataItem{{Data: []interface{}{0, 1, 2, 3, 4}}, {Data: []interface{}{4, 5, 6, 7, 8}}}}
	Jsampledata, err := json.Marshal(sampledata)
	if err == nil {
		fmt.Println(string(Jsampledata))
	}
	model := nclink.NewModel(nclink.DefaultModel)
	for _, samplechannel := range model.Devices[0].SampleTasks {
		Jsamplechannel, _ := json.Marshal(samplechannel)
		fmt.Println(string(Jsamplechannel))
	}
}

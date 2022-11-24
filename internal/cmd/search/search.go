package search

import (
	"fmt"
	//	"log"
	"context"
	"encoding/json"
	"os"

	//	"path/filepath"

	"github.com/spf13/cobra"
	//	"gopkg.in/yaml.v3"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// Monitor struct... used for detailed monitor data
type Monitors struct {
	Monitors []Monitor `json:"monitors"`
}

type Monitor struct {
	Created       string   `json:"created"`
	Creator       Creator  `json:"creator"`
	Deleted       string   `json:"deleted"`
	Id            int32    `json:"id"`
	Message       string   `json:"message"`
	Modified      string   `json:"modified"`
	Multi         string   `json:"multi"`
	Name          string   `json:"name"`
	Options       Options  `json:"options"`
	Overall_State string   `json:"overall_state"`
	Priority      string   `json:"priority"`
	Query         string   `json:"query"`
	Tags          []string `json:"tags"`
	Type          string   `json:"type"`
}

type Creator struct {
	Email  string `json:"email"`
	Handle string `json:"handle"`
	Name   string `json:"name"`
}

type Options struct {
	Escalation_Message string     `json:"escalation_message"`
	Locked             string     `json:"locked"`
	New_Group_Delay    string     `json:"new_group_delay"`
	Notify_Audit       string     `json:"notify_audit"`
	Notify_No_Data     string     `json:"notify_no_data"`
	Renotify_Interval  string     `json:"renotify_interval":`
	Silenced           string     `json:"silenced"`
	Thresholds         Thresholds `json:"thresholds"`
	Timeout_h          string     `json:"timeout_h"`
}

type Thresholds struct {
	Critical string `json:"critical"`
	OK       string `json:"ok"`
	Warning  string `json:"warning"`
}

// Counts struct... used for count summaries

/*
type Counts struct {
    Muted Muted `json:"muted"`
}*/

// NewCmdSearch is the service catalog generate command.
func NewCmdSearch() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "searches monitors",
		Long:  "searches monitors",
		//Run:   search,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				monitorSummary()
			} else {
				argument := args[0]
				searchQuery(argument)
			}
			//fmt.Printf("argument: %T", argument)
		},
	}
}

func searchQuery(argument string) {

	query := argument
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.SearchMonitors(ctx, *datadogV1.NewSearchMonitorsOptionalParameters().WithQuery(query))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.SearchMonitors`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	//fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.SearchMonitors`:\n%s\n", responseContent)

	var monitors Monitors
	json.Unmarshal([]byte(responseContent), &monitors)

	for i := 0; i < len(monitors.Monitors); i++ {
		fmt.Println("Monitor Name:" + monitors.Monitors[i].Name)
		fmt.Println("Created By: " + monitors.Monitors[i].Creator.Email)
		fmt.Println("Created: " + monitors.Monitors[i].Created)
		fmt.Printf("ID: %d\n", monitors.Monitors[i].Id)
		for j := 0; j < len(monitors.Monitors[i].Tags); j++ {
			fmt.Printf("Tag: %s\n", monitors.Monitors[i].Tags[j])
		}
		fmt.Println()
	}
}

func monitorSummary() {

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMonitorsApi(apiClient)
	resp, r, err := api.SearchMonitors(ctx, *datadogV1.NewSearchMonitorsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.SearchMonitors`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.SearchMonitors`:\n%s\n", responseContent)
}

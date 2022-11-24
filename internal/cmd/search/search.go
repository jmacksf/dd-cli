package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

type Counts struct {
	Count []Count `json:"counts"`
}

type Count struct {
	Muted  Muted  `json:"muted"`
	Status Status `json:"status"`
	Tag    Tag    `json:"tag"`
	Type   Type   `json:"type"`
}

type Muted struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
}

type Status struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
}

type Tag struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
}

type Type struct {
	Count int32  `json:"count"`
	Name  string `json:"name"`
}

// NewCmdSearch is the dd-cli command.
func NewCmdSearch() *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitor",
		Short: "searches monitors",
		Long:  "searches monitors",
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("out")
			fmt.Printf(output)
			if len(args) < 1 {
				monitorSummary(output)
			} else {
				argument := args[0]
				searchQuery(argument, output)
			}
		},
	}
	addFlags(&cmd)
	return &cmd
}

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("out", "json", "output format")
}

func searchQuery(argument string, output string) {

	if !(checkOutputCmd(output)) {
		log.Fatal("output command not supported, must be either \"json\" or \"text\"")
	}

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
	switch output {
	case "json":
		monitorPrintJson(responseContent)
	case "text":
		monitorPrintTxt(responseContent)
	}
}

func monitorSummary(output string) {

	if !(checkOutputCmd(output)) {
		log.Fatal("output command not supported, must be either \"json\" or \"text\"")
	}

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
	//fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.SearchMonitors`:\n%s\n", responseContent)
	monitorPrintJson(responseContent)
}

func monitorPrintTxt(responseContent []byte) {

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

func monitorPrintJson(responseContent []byte) {
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.SearchMonitors`:\n%s\n", responseContent)
}

func checkOutputCmd(output string) bool {
	outputCommands := []string{
		"json",
		"text",
	}
	for _, item := range outputCommands {
		if item == output {
			return true
		}
	}
	return false
}

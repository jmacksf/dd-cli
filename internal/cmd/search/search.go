package generate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)


// NewCmdSearch is the service catalog generate command.
func NewCmdSearch() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "searches monitors",
		Long:  "searches monitors",
		Run:   search,
	}
}



func search(*cobra.Command, []string) {
	fmt.Printf("Running search")
}

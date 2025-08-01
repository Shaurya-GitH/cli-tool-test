package sub

import (
	"cli-tool/types"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var NameCmd = &cobra.Command{
	Use:   "name [id]",
	Short: "returns the name of id",
	Long:  "returns the name associated with the id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{}
		url := "https://688c48b7cd9d22dda5cc9337.mockapi.io/api/v1/comments/" + args[0]
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		resp, err := client.Do(r)
		responseBody := &types.Data{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(responseBody.Name)
	},
}

func init() {

}

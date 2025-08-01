package sub

import (
	"bytes"
	"cli-tool/types"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a record in db",
	Long:  "posts a random record in the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{}
		responseBody := &types.Data{}
		url := "https://688c48b7cd9d22dda5cc9337.mockapi.io/api/v1/comments"
		r, err := http.NewRequest("POST", url, bytes.NewBufferString(""))
		if err != nil {
			fmt.Println(err)
		}
		resp, err := client.Do(r)
		if err != nil {
			return
		}
		if resp.StatusCode == 201 {
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			fmt.Println(responseBody.Name)
		} else {
			fmt.Println(resp.StatusCode)
		}
	},
}

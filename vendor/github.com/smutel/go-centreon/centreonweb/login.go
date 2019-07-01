package centreonweb

import (
	"encoding/json"
	"net/http"

	"github.com/smutel/go-centreon/client"
)

type login struct {
	AuthToken string `json:"authToken"`
}

func (cli *ClientCentreonWeb) login() error {
	reqInputs := client.RequestInput{
		Method: http.MethodPost,
		Path:   centreonAPIPath,
		Query:  cli.AuthQuery,
		Header: cli.AuthHeader,
		Body:   cli.AuthInput,
	}

	respReader, err := cli.MainClient.ExecuteRequest(reqInputs)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var login login
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&login)

	cli.ConfigHeader.Set("centreon-auth-token", login.AuthToken)

	return nil
}

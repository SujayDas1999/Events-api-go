package secrets_reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Secrets struct {
	PostgresConn string `json:"postgres_conn"`
}

func SecretsReader(path string) Secrets {

	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return Secrets{}
	}

	var secrets Secrets

	err = json.Unmarshal(jsonData, &secrets)

	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return Secrets{}
	}

	return secrets

}

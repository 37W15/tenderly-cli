package call

import (
	"encoding/json"
	"github.com/tenderly/tenderly-cli/config"
	"github.com/tenderly/tenderly-cli/model"
	"github.com/tenderly/tenderly-cli/rest/client"
	"github.com/tenderly/tenderly-cli/rest/payloads"
	"strings"
)

type ContractCalls struct {
}

func NewContractCalls() *ContractCalls {
	return &ContractCalls{}
}

func (rest *ContractCalls) UploadContracts(request payloads.UploadContractsRequest, projectSlug string) (*payloads.UploadContractsResponse, error) {
	uploadJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	accountID := config.GetGlobalString(config.AccountID)
	if strings.Contains(projectSlug, "/") {
		projectInfo := strings.Split(projectSlug, "/")
		accountID = projectInfo[0]
		projectSlug = projectInfo[1]
	}

	var contracts *payloads.UploadContractsResponse

	response := client.Request(
		client.PostMethod,
		"api/v1/account/"+accountID+"/project/"+projectSlug+"/contracts",
		uploadJson,
	)

	err = json.NewDecoder(response).Decode(&contracts)
	return contracts, err
}

func (rest *ContractCalls) VerifyContracts(request payloads.UploadContractsRequest) (*payloads.UploadContractsResponse, error) {
	uploadJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	var contracts *payloads.UploadContractsResponse

	response := client.Request(
		client.PostMethod,
		"api/v1/account/me/verify-contracts",
		uploadJson,
	)

	err = json.NewDecoder(response).Decode(&contracts)
	return contracts, err
}

func (rest *ContractCalls) GetContracts(id string) ([]*model.Contract, error) {
	var contracts []*model.Contract

	response := client.Request(
		client.GetMethod,
		"api/v1/account/"+config.GetString("Username")+"/project/"+id,
		nil,
	)

	err := json.NewDecoder(response).Decode(contracts)
	return contracts, err
}

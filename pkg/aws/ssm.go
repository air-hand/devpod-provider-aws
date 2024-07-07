package aws

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ssmPortForwardingParameters struct {
	PortNumber      []string `json:"portNumber"`
	LocalPortNumber []string `json:"localPortNumber"`
}

func CommandArgsSSMTunneling(instanceID string, localPort string) ([]string, error) {
	parameters := &ssmPortForwardingParameters{
		PortNumber:      []string{"22"},
		LocalPortNumber: []string{localPort},
	}

	parameters_as_json, err := json.Marshal(parameters)
	if err != nil {
		return []string{}, err
	}

	//	parameters_escaped := strings.ReplaceAll(string(parameters_as_json), "\"", "\\\"")
	parameters_escaped := strings.ReplaceAll(string(parameters_as_json), "\"", "\"")

	return []string{
		"ssm", "start-session",
		"--target", instanceID,
		"--document-name", "AWS-StartPortForwardingSession",
		//		fmt.Sprintf("--parameters=\"%s\"", parameters_escaped),
		fmt.Sprintf("--parameters=%s", parameters_escaped),
	}, nil
}

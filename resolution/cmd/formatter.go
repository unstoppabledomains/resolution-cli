package cmd

import (
	"encoding/json"
	"strings"
)

func formatOutput(v interface{}) (string, error) {
	var outputBuilder strings.Builder
	output, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		return "", err
	}
	_, err = outputBuilder.Write(output)
	if err != nil {
		return "", err
	}
	return outputBuilder.String(), nil
}

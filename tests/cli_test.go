package tests

import (
	"fmt"
	"github.com/rendon/testcli"
	"os"
	"path"

	//"github.com/stretchr/testify"
	"testing"
)

var dir, _ = os.Getwd()
var commandPath = path.Join(dir, "resolution")

func TestCliResolve(t *testing.T) {
	testcli.Run(commandPath, "resolve", "-d", "brad.crypto")
	fmt.Println(testcli.Stdout())
	fmt.Println(testcli.Stderr())
}

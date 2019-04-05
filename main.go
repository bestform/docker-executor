package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"strings"
)

var CMD = []string{"php", "-v"}
var IDENTIFIER = "-php"

func main() {
	var err error

	defer func() {
		if err != nil {
			fmt.Println("Error running client:", err)
			os.Exit(1)
		}
	}()

	cli, err := client.NewEnvClient()
	if err != nil { return }

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { return }

	for _, c := range containers {
		if isTargetContainerByName(c.Names) {
			err = executeCommand(cli, c)
			if err != nil { return }
		}
	}

}

func executeCommand(cli *client.Client, c types.Container) error {
	fmt.Println("Executing on container", c.ID)
	idResp, err := cli.ContainerExecCreate(context.Background(), c.ID, types.ExecConfig{Cmd:CMD, AttachStdout:true, AttachStderr:true})
	if err != nil { return err }

	hjResp, err := cli.ContainerExecAttach(context.Background(), idResp.ID, types.ExecConfig{})
	if err != nil { return err }
	defer hjResp.Close()

	output, err := ioutil.ReadAll(hjResp.Reader)
	if err != nil { return err }

	fmt.Println(string(output))

	return nil
}

func isTargetContainerByName(names []string) bool {
	for _, n := range names {
		if strings.Contains(n, IDENTIFIER) {
			return true
		}
	}

	return false
}

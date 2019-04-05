package main

import (
	"fmt"
	"github.com/bestform/dockerExecutor/internal"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var config internal.Config

func main() {
	var err error

	config, err = internal.ConfigFromEnvironment()
	if err != nil { return }

	http.HandleFunc("/", handleRequest)

	if err = http.ListenAndServe(":" + strconv.Itoa(config.Port), nil); err != nil {
		return
	}
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("Error running client:", err)
			writer.WriteHeader(500)
		}
	}()

	writer.Header().Add("Content-Type", "text/html")

	cli, err := client.NewEnvClient()
	if err != nil { return }

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { return }
	for _, c := range containers {
		if isTargetContainerByName(c.Names, config) {
			output, err := executeCommand(cli, c, config)
			if err != nil { return }
			_, err = fmt.Fprint(writer, string(output))
			if err != nil { return }
		}
	}
}

func executeCommand(cli *client.Client, c types.Container, config internal.Config) ([]byte, error) {
	var err error
	var output []byte

	cmdParts := strings.Split(config.Cmd, " ")
	idResp, err := cli.ContainerExecCreate(context.Background(), c.ID, types.ExecConfig{Cmd:cmdParts, AttachStdout:true, AttachStderr:true})
	if err != nil { return output, err }

	hjResp, err := cli.ContainerExecAttach(context.Background(), idResp.ID, types.ExecConfig{})
	if err != nil { return output, err }
	defer hjResp.Close()

	output, err = ioutil.ReadAll(hjResp.Reader)
	if err != nil { return output, err }

	return output, nil
}

func isTargetContainerByName(names []string, config internal.Config) bool {
	for _, n := range names {
		if strings.Contains(n, config.Identifier) {
			return true
		}
	}

	return false
}

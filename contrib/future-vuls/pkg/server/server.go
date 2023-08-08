package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/future-architect/vuls/config"
	"github.com/future-architect/vuls/contrib/future-vuls/pkg/schema"
	"github.com/future-architect/vuls/util"
)

type createPseudoServerInput struct {
	APIKey     string `json:"api_key"`
	ServerName string `json:"serverName"`
}

type createPseudoServerOutput struct {
	UUID string `json:"serverUuid"`
}

func AddServerToFvuls(url string, token string, outputFile string) error {
	var servers map[string]*schema.ServerDetail
	_, err := toml.DecodeFile(outputFile, &servers)
	if err != nil {
		return fmt.Errorf("failed to read %s", outputFile)
	}

	targetServerCount := 0
	//This flag indicates whether at least one server has been added
	isAnyServerAdded := false
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for addr, params := range servers {
		if params.FvulsSync && params.UUID == "" {
			fmt.Printf("%v: Adding to Fvuls server...\n", addr)
			targetServerCount++
			uuid, err := createPseudoServer(ctx, url, token, addr)
			if err != nil {
				fmt.Printf("%s: Failed to add to Fvuls server. err: %v\n", addr, err)
				continue
			}
			servers[addr].UUID = uuid
			fmt.Printf("%s: Done.\n", addr)
			isAnyServerAdded = true
		}
	}
	if targetServerCount == 0 {
		fmt.Printf("The target of the command could not be found, please rewrite the FvulsSync to true.\n")
		return fmt.Errorf("command target not found error")
	}
	if !isAnyServerAdded {
		return fmt.Errorf("all server uploads failed")
	}

	f, err := os.OpenFile(outputFile, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("failed to open toml file. err: %v", err)
	}
	defer f.Close()
	encoder := toml.NewEncoder(f)
	if err := encoder.Encode(servers); err != nil {
		return fmt.Errorf("failed to write to %s. err: %v", outputFile, err)
	}
	fmt.Printf("Successfully wrote to %s\n", outputFile)
	return nil
}

func createPseudoServer(ctx context.Context, url string, token string, addr string) (string, error) {
	payload := createPseudoServerInput{
		ServerName: addr,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to Marshal to JSON: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/server/pseudo", url), bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	client, err := util.GetHTTPClient(config.Conf.HTTPProxy)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("falied to create pseudo server. statusCode: %d", resp.StatusCode)
	}
	t, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response data. err: %v", err)
	}
	var serverDetail createPseudoServerOutput
	if err := json.Unmarshal(t, &serverDetail); err != nil {
		if err.Error() == "invalid character 'A' looking for beginning of value" {
			return "", fmt.Errorf("invalid token")
		}
		return "", fmt.Errorf("failed to unmarshal serverDetail. err: %v", err)
	}
	return serverDetail.UUID, nil
}

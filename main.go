package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
)

// Instance はEC2インスタンスの情報を保持します。
type Instance struct {
	InstanceID   string `json:"InstanceID"`
	InstanceName string `json:"InstanceName"`
}

func main() {
	// 環境変数でAWS_PROFILEやAWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEYが設定されていることを前提とします
	cmd := exec.Command("aws", "ec2", "describe-instances", "--query", "Reservations[*].Instances[*].{InstanceID:InstanceId, InstanceName:Tags[?Key=='Name'] | [0].Value}", "--output", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		fmt.Printf("%s\n", output)
		os.Exit(1)
	}

	// JSONデータを解析します
	var instances [][]Instance
	if err := json.Unmarshal(output, &instances); err != nil {
		fmt.Printf("JSON parsing error: %s\n", err)
		os.Exit(1)
	}

	// インスタンスのリストを平坦化します
	var flatInstances []Instance
	for _, reservation := range instances {
		for _, instance := range reservation {
			if instance.InstanceName != "" {
				flatInstances = append(flatInstances, instance)
			}
		}
	}

	// fuzzyfinderを使用してインスタンスを選択します
	idx, err := fuzzyfinder.Find(
		flatInstances,
		func(i int) string {
			return flatInstances[i].InstanceName
		},
		fuzzyfinder.WithPromptString("Select an instance: "),
	)
	if err != nil {
		fmt.Printf("Fuzzy finder error: %s\n", err)
		os.Exit(1)
	}

	selectedInstance := flatInstances[idx]
	fmt.Printf("Selected Instance ID: %s, Name: %s\n", selectedInstance.InstanceID, selectedInstance.InstanceName)
}

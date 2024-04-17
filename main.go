package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"log/slog"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/ktr0731/go-fuzzyfinder"
)

// config
var profile string
var region string
var output string

func init() {
	flag.StringVar(&profile, "profile", "", "AWS profile name")
	flag.StringVar(&profile, "p", "", "AWS region"+"(shorthand)")
	flag.StringVar(&region, "region", "ap-northeast-1", "AWS region")
	flag.StringVar(&region, "r", "ap-northeast-1", "AWS region"+"(shorthand)")
	flag.StringVar(&output, "output", "text", "Output format (text or json)")
	flag.StringVar(&output, "o", "text", "Output format (text or json)"+"(shorthand)")
}

type InstanceInfo struct {
	InstanceID   string
	InstanceName string
}

func flattenInstance(instances []types.Reservation) []InstanceInfo {
	var flatInstances []InstanceInfo
	for _, reservation := range instances {
		for _, instance := range reservation.Instances {
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					flatInstances = append(flatInstances,
						InstanceInfo{
							InstanceID:   *instance.InstanceId,
							InstanceName: *tag.Value,
						})
				}
			}
		}
	}
	return flatInstances
}

func setLogLevel() *slog.Logger {
	var level slog.Level
	// Assume an environment variable `LOG_LEVEL` that dictates the desired log level
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default log level
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	return logger
}

func main() {
	flag.Parse()
	logger := setLogLevel()
	logger.Debug("Parameters", "profile", profile, "region", region)
	ctx := context.TODO()
	cfg, err := checkConfig(profile, region, ctx)
	logger.Debug("Config", "cfg", cfg)
	if err != nil {
		logger.Error("Failed to load config: %s\n", "error", err)
		os.Exit(1)
	}
	nameFilter := "tag:Name"
	filters := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   &nameFilter,
				Values: []string{"*"},
			},
		},
	}
	client := ec2.NewFromConfig(cfg)
	logger.Debug("Start describing instances")
	result, err := client.DescribeInstances(context.TODO(), filters)
	if err != nil {
		logger.Error("Failed to describe instances: %s\n", "error", err)
		os.Exit(1)
	}
	logger.Debug("Finish describing instances")
	instances := result.Reservations
	flatInstances := flattenInstance(instances)
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

	if output == "json" {
		if jsonBytes, err := json.Marshal(selectedInstance); err != nil {
			logger.Error("Failed to marshal JSON: %s\n", "error", err)
			os.Exit(1)
		} else {
			fmt.Println(string(jsonBytes))
		}
	} else {
		fmt.Printf("Selected Instance ID: %s, Name: %s\n", selectedInstance.InstanceID, selectedInstance.InstanceName)
	}
}

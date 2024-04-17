package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// profileが指定されている場合はそのprofileを使ってconfigを読み込む
// regionが指定されている場合はそのregionを使ってconfigを返す
// defaultのregionはap-northeast-1
func checkConfig(profile, region string, ctx context.Context) (aws.Config, error) {
	if profile != "" {
		cfg, err := config.LoadDefaultConfig(
			ctx,
			config.WithSharedConfigProfile(profile),
			config.WithDefaultRegion(region),
		)
		return cfg, err
	}
	return config.LoadDefaultConfig(ctx, config.WithDefaultRegion(region))
}

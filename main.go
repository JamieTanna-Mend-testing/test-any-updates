package main

import (
	"context"
	"fmt"

	m "github.com/JamieTanna-Mend-testing/private-go-module-child"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/oapi-codegen/nullable"
	_ "golang.org/x/net/http/httpproxy"
)

func main() {
	p := struct {
		N nullable.Nullable[int]
	}{}

	p.N = nullable.NewNullNullable[int]()

	fmt.Printf("Specified: %v\n", p.N.IsSpecified())
	fmt.Printf("Null: %v\n", p.N.IsNull())

	fmt.Println(m.Hello("JVT"))

	// Each AWS SDK for Go v2 module has its own large CHANGELOG.md, so updating
	// them exercises Renovate's changelog parsing/caching (renovatebot/renovate#47334)
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("eu-west-1"))
	if err != nil {
		fmt.Println("unable to load AWS config:", err)
		return
	}

	_ = s3.NewFromConfig(cfg)
	_ = dynamodb.NewFromConfig(cfg)
	_ = sqs.NewFromConfig(cfg)
	_ = sts.NewFromConfig(cfg)

	fmt.Println("Region:", aws.ToString(&cfg.Region))
}

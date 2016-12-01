package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "iamkeys IAMUSER")
		os.Exit(1)
	}
	userName := os.Args[1]

	sess, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set AWS session. error: %s", err)
		os.Exit(1)
	}

	svc := iam.New(sess, &aws.Config{})

	resp, err := svc.CreateAccessKey(&iam.CreateAccessKeyInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY. error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("AWS_ACCESS_KEY_ID:     %s\n", *resp.AccessKey.AccessKeyId)
	fmt.Printf("AWS_SECRET_ACCESS_KEY: %s\n", *resp.AccessKey.SecretAccessKey)
}

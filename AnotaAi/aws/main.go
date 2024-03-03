package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Message struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
	OwnerID     string  `json:"ownerID"`
}

type Catalog struct {
	Categories []Message `json:"categories"`
	Products   []Message `json:"products"`
}

type ContentBody struct {
	Message string `json:"message"`
}

func InsertOrUpdate(catalog *Catalog, data Message, isProduct bool) {

	if isProduct {
		for idx, content := range catalog.Products {
			if content.Id == data.Id {
				catalog.Products[idx] = data
				return
			}
		}
		catalog.Products = append(catalog.Products, data)

	} else {

		for idx, content := range catalog.Categories {
			if content.Id == data.Id {
				catalog.Categories[idx] = data
				return
			}
		}
		catalog.Categories = append(catalog.Categories, data)
	}

}

func GetObjectFromBucket(key string, s3Client *s3.S3) (*s3.GetObjectOutput, error) {
	data, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("anotaai-catalog-marketplace"),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Println("Erro ao encontrar item")
		return nil, err
	}

	return data, nil
}

func s3Client() *s3.S3 {
	sess := session.Must(session.NewSession())

	ddbCustResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		return endpoints.ResolvedEndpoint{
			URL:           "http://172.17.0.1:4566",
			SigningRegion: endpoints.UsEast1RegionID,
		}, nil
	}
	return s3.New(sess, &aws.Config{
		Region:           aws.String(endpoints.UsEast1RegionID),
		EndpointResolver: endpoints.ResolverFunc(ddbCustResolverFn),
		S3ForcePathStyle: aws.Bool(true),
	})
}

func HandleRequest(ctx context.Context, event events.SQSEvent) (string, error) {
	var timeout time.Duration = time.Hour

	internCtx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		internCtx, cancelFn = context.WithTimeout(internCtx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}

	s3Client := s3Client()

	for _, element := range event.Records {

		var (
			data    Catalog
			key     string
			msgData Message
		)

		var eventBody ContentBody

		if err := json.Unmarshal([]byte(element.Body), &eventBody); err != nil {
			fmt.Println("erro", err)
			os.Exit(0)
		}

		if err := json.Unmarshal([]byte(eventBody.Message), &msgData); err != nil {
			fmt.Println("erro", err)
			os.Exit(0)
		}

		key = fmt.Sprintf("%s-catalog.json", msgData.OwnerID)

		bucketContent, err := GetObjectFromBucket(key, s3Client)

		if err == nil {

			if err != nil {
				fmt.Println("Error:", err)
			}

			var catalog Catalog

			buf := new(bytes.Buffer)

			if _, err = buf.ReadFrom(bucketContent.Body); err != nil {
				fmt.Println("Error:", err)
				os.Exit(0)
			}

			if err = json.Unmarshal(buf.Bytes(), &catalog); err != nil {
				fmt.Println("Error:", err)
				os.Exit(0)
			}

			bucketContent.Body.Close()

			data = catalog

		}

		InsertOrUpdate(&data, msgData, strings.Contains(eventBody.Message, "Price"))

		content, err := json.Marshal(data)

		if err != nil {
			fmt.Println("erro", err)
		}

		stdinData := strings.NewReader(string(content))

		_, err = s3Client.PutObjectWithContext(internCtx, &s3.PutObjectInput{
			Bucket: aws.String("anotaai-catalog-marketplace"),
			Key:    aws.String(key),
			Body:   stdinData,
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
				fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
			}
			os.Exit(1)
		}

	}

	return "Ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}

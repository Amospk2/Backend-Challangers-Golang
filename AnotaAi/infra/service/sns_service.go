package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsService struct {
	snsClient *sns.SNS
}

func CreateSession() *SnsService {
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.

	ddbCustResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		return endpoints.ResolvedEndpoint{
			URL:           "http://172.17.0.1:4566",
			SigningRegion: endpoints.UsEast1RegionID,
		}, nil
	}
	snsClient := sns.New(sess, &aws.Config{
		Region:           aws.String(endpoints.UsEast1RegionID),
		EndpointResolver: endpoints.ResolverFunc(ddbCustResolverFn),
		S3ForcePathStyle: aws.Bool(true),
	})

	return &SnsService{snsClient: snsClient}

}

func (s *SnsService) PublishInTopic(message string) {
	s.snsClient.Publish(
		&sns.PublishInput{
			Message:  aws.String(message),
			TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:my-topic"),
		},
	)
}

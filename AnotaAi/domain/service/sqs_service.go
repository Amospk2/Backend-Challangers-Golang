package domain_service

type SqsService interface {
	PublishInTopic(msg string)
}

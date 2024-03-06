package mock

type SnsServiceMock struct {
	msg []string
}

func (s *SnsServiceMock) PublishInTopic(message string) {
	s.msg = append(s.msg, message)
}

func NewSqsServiceMock(msg []string) *SnsServiceMock {
	return &SnsServiceMock{msg: msg}
}

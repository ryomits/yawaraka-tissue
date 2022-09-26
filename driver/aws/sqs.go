package aws

type SQS struct {
	url                 string
	delaySecounds       int64
	maxNumberOfMessages int64
	waitTimeSeconds     int64
	backOffSeconds      int64
}

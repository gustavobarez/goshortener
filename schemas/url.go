package schemas

type URL struct {
	ID string `dynamodbav:"id"`
	OriginalURL string `dynamodbav:"original_url"`
}
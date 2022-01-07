package elasticsearch

import elastic "github.com/elastic/go-elasticsearch/v7"

func GetClient() (*elastic.Client, error) {
	client, err := elastic.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/microcmsio/microcms-go-sdk"
)

type Topic struct {
	ID          string    `json:"id,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	RevisedAt   time.Time `json:"revisedAt,omitempty"`
	PostDate    string    `json:"postDate,omitempty"`
	Title       string    `json:"title,omitempty"`
	Url         string    `json:"url,omitempty"`
}

type TopicList struct {
	Contents   []Topic
	TotalCount int
	Limit      int
	Offset     int
}

type Cms struct {
	serviceDomain string
	apiKey        string
	client        *microcms.Client
}

func (c *Cms) fetchList(endpoint string) (topics TopicList, err error) {
	err = c.client.List(microcms.ListParams{Endpoint: endpoint}, &topics)
	return
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	serviceDomain := os.Getenv("SERVICE_DOMAIN")
	apiKey := os.Getenv("API_KEY")
	client := microcms.New(serviceDomain, apiKey)
	cms := &Cms{serviceDomain, apiKey, client}
	res, _ := cms.fetchList("topics")

	fmt.Printf("%+v", res.Contents[0].Title)
}

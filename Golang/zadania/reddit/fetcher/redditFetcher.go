package fetcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type HTTPClient struct {
	c    *http.Client
	host string
}

type MyFetcher struct {
	data       *response
	httpClient *HTTPClient
	errors     []string
}

func (t *MyFetcher) Fetch() (outErr error) {
	defer func() {
		if t.errors == nil || len(t.errors) == 0 {
			return
		}
		outErr = fmt.Errorf("%d errors, Last: %w", len(t.errors), t.errors[len(t.errors)-1])
	}()

	if t.httpClient == nil {
		t.httpClient = &HTTPClient{
			host: "https://www.reddit.com/r/golang.json",
			c: &http.Client{
				Timeout: 5 * time.Second,
			},
		}
	}

	if t.errors == nil {
		t.errors = make([]string, 0)
	}

	req, err := http.NewRequest(http.MethodGet, t.httpClient.host, http.NoBody)
	if err != nil {
		t.errors = append(t.errors, fmt.Sprintf("cannot create request: %w", err))
		return
	}

	resp, err := t.httpClient.c.Do(req)
	if err != nil {
		t.errors = append(t.errors, fmt.Sprintf("cannot get data: %w", err))
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.errors = append(t.errors, fmt.Sprintf("Status code != 200: %d", resp.StatusCode))
		return
	}

	tmp, _ := io.ReadAll(resp.Body)
	buf := bytes.NewBuffer([]byte{})
	buf.Write(tmp)

	//t.data = response{}
	err = json.NewDecoder(buf).Decode(&t.data)
	if err != nil {
		t.errors = append(t.errors, fmt.Sprintf("cannot unmarshal data: %w", err))
		return
	}

	return
}

func (t *MyFetcher) Save(writer io.Writer) error {
	if t.httpClient == nil {
		return fmt.Errorf("Run the Fetch method first!")
	}
	if t.data == nil {
		return fmt.Errorf("No data!")
	}
	//fmt.Printf("%#v", t.data)
	for _, v := range t.data.Data.Children {
		//fmt.Printf("\n\n%#v", v)
		//fmt.Printf("\n\n%#v\n%#v", v.Data.Title, v.Data.URL)
		//var lala string = fmt.Sprintf("%v\n%v\n", v.Data.Title, v.Data.URL)
		//fmt.Printf("\nTest string: %v", lala)
		//writer.Write([]byte(v.Data.Title))
		//writer.Write([]byte("test"))
		//fmt.Fprintf(writer, "%v\n%v\n", v.Data.Title, v.Data.URL)
		//fmt.Fprintf(writer, v.Data.Title)
		writer.Write([]byte(fmt.Sprintf("%v\n%v\n", v.Data.Title, v.Data.URL)))
	}

	return nil
}

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func NewEhanceHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     time.Duration(2) * time.Second,
			DisableCompression:  false,
		},
	}

	return client
}

func HttpGetWithProcessor(ctx context.Context, client *http.Client, url string, processor func(r io.Reader) error) error {
	errch := make(chan error, 1)
	var wg sync.WaitGroup
	pr, pw := io.Pipe()

	// get
	resp, err := get(ctx, client, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read resp body stream from pipe
	// let the callback function prcoess the stream
	wg.Add(1)
	go func() {
		defer wg.Done()
		errch <- processor(pr)

		pr.Close()
	}()

	// write resp body stream to pipe
	io.Copy(pw, resp.Body)
	pw.Close()

	wg.Wait()

	return <-errch
}

func HttpGetJson(client *http.Client, url string, res interface{}) error {
	return HttpGetJsonWithContext(context.Background(), client, url, res)
}

func HttpGetJsonWithContext(ctx context.Context, client *http.Client, url string, res interface{}) error {
	return HttpGetWithProcessor(ctx, client, url, func(r io.Reader) error {
		return json.NewDecoder(r).Decode(res)
	})
}

func HttpGetBytes(client *http.Client, url string) ([]byte, error) {
	return HttpGetBytesWithContext(context.Background(), client, url)
}

func HttpGetBytesWithContext(ctx context.Context, client *http.Client, url string) ([]byte, error) {
	res := &bytes.Buffer{}

	err := HttpGetWithProcessor(ctx, client, url, func(r io.Reader) error {
		_, err := io.Copy(res, r)
		return err
	})

	return res.Bytes(), err

}

func header(req *http.Request) {
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0")

}

func Retry(retries int, fn func() error) error {
	var err error
	for retry := 0; retry < retries; retry++ {
		if retry > 0 {
			time.Sleep(time.Duration(retry) * time.Second)
		}

		err = fn()
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("fail with %v after %d retries", err, retries)

}

func get(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	header(req)

	return client.Do(req)
}

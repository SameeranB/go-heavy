package providers

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPProviderInterface interface {
	Get(ctx context.Context, baseURL string, path string, headers map[string]string, queryParams map[string]string) (http.Response, error)
}

type HTTPProvider struct {
	Client *http.Client
}

func NewHTTPProvider() HTTPProvider {
	return HTTPProvider{
		Client: &http.Client{},
	}
}

func (p HTTPProvider) Get(ctx context.Context, baseURL string, path string, headers map[string]string, queryParams map[string]string) (http.Response, error) {
	fullPath := baseURL + path
	if queryParams != nil {
		fullPath = fullPath + "?"
		for key, value := range queryParams {
			fullPath = fullPath + key + "=" + value + "&"
		}
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fullPath, nil)
	if err != nil {
		err = fmt.Errorf("HTTPProvider.GET - error creating request: %w", err)
		return http.Response{}, err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	res, err := p.Client.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTPProvider.GET - error making request: %w", err)
		return http.Response{}, err
	}

	return *res, nil

}

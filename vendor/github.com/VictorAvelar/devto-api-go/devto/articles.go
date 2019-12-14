package devto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

// ArticlesResource implements the APIResource interface
// for devto articles.
type ArticlesResource struct {
	API *Client
}

// List will return the articles uploaded to devto, the result
// can be narrowed down, filtered or enhanced using query
// parameters as specified on the documentation.
// See: https://docs.dev.to/api/#tag/articles/paths/~1articles/get
func (ar *ArticlesResource) List(ctx context.Context, opt ArticleListOptions) ([]ListedArticle, error) {
	q, err := query.Values(opt)
	if err != nil {
		return nil, err
	}
	req, err := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles?%s", q.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if nonSuccessfulResponse(res) {
		return nil, unmarshalErrorResponse(res)
	}
	var articles []ListedArticle
	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// ListForTag is a convenience method for retrieving articles
// for a particular tag, calling the base List method.
func (ar *ArticlesResource) ListForTag(ctx context.Context, tag string, page int) ([]ListedArticle, error) {
	return ar.List(ctx, ArticleListOptions{Tags: tag, Page: page})
}

// ListForUser is a convenience method for retrieving articles
// written by a particular user, calling the base List method.
func (ar *ArticlesResource) ListForUser(ctx context.Context, username string, page int) ([]ListedArticle, error) {
	return ar.List(ctx, ArticleListOptions{Username: username, Page: page})
}

// ListMyPublishedArticles lists all published articles
// written by the user authenticated with this client,
// erroring if the caller is not authenticated. Articles in
// the response will be listed in reverse chronological order
// by their publication times.
//
// If opts is nil, then no query parameters will be sent; the
// page number will be 1 and the page size will be 30
// articles.
func (ar *ArticlesResource) ListMyPublishedArticles(ctx context.Context, opts *MyArticlesOptions) ([]ListedArticle, error) {
	return ar.listMyArticles(ctx, "api/articles/me/published", opts)
}

// ListMyUnpublishedArticles lists all unpublished articles
// written by the user authenticated with this client,
// erroring if the caller is not authenticated. Articles in
// the response will be listed in reverse chronological order
// by their creation times.
//
// If opts is nil, then no query parameters will be sent; the
// page number will be 1 and the page size will be 30
// articles.
func (ar *ArticlesResource) ListMyUnpublishedArticles(ctx context.Context, opts *MyArticlesOptions) ([]ListedArticle, error) {
	return ar.listMyArticles(ctx, "api/articles/me/unpublished", opts)
}

// ListAllMyArticles lists all articles written by the user
// authenticated with this client, erroring if the caller is
// not authenticated. Articles in the response will be listed
// in reverse chronological order by their creation times,
// with unpublished articles listed before published articles.
//
// If opts is nil, then no query parameters will be sent; the
// page number will be 1 and the page size will be 30
// articles.
func (ar *ArticlesResource) ListAllMyArticles(ctx context.Context, opts *MyArticlesOptions) ([]ListedArticle, error) {
	return ar.listMyArticles(ctx, "api/articles/me/all", opts)
}

// listMyArticles serves for handling roundtrips to the
// /api/articles/me/* endpoints, requesting articles from the
// endpoint passed in, and returning a list of articles.
func (ar *ArticlesResource) listMyArticles(
	ctx context.Context,
	endpoint string,
	opts *MyArticlesOptions,
) ([]ListedArticle, error) {
	if ar.API.Config.InsecureOnly {
		return nil, ErrProtectedEndpoint
	}

	req, err := ar.API.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(APIKeyHeader, ar.API.Config.APIKey)

	if opts != nil {
		q, err := query.Values(opts)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if nonSuccessfulResponse(res) {
		return nil, unmarshalErrorResponse(res)
	}
	var articles []ListedArticle
	if err := json.NewDecoder(res.Body).Decode(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// Find will retrieve an Article matching the ID passed.
func (ar *ArticlesResource) Find(ctx context.Context, id uint32) (Article, error) {
	req, err := ar.API.NewRequest(http.MethodGet, fmt.Sprintf("api/articles/%d", id), nil)
	if err != nil {
		return Article{}, err
	}

	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer res.Body.Close()

	if nonSuccessfulResponse(res) {
		return Article{}, unmarshalErrorResponse(res)
	}
	var art Article
	if err := json.NewDecoder(res.Body).Decode(&art); err != nil {
		return Article{}, err
	}
	return art, nil
}

// New will create a new article on dev.to
func (ar *ArticlesResource) New(ctx context.Context, u ArticleUpdate) (Article, error) {
	if ar.API.Config.InsecureOnly {
		return Article{}, ErrProtectedEndpoint
	}
	cont, err := json.Marshal(&u)
	if err != nil {
		return Article{}, err
	}
	req, err := ar.API.NewRequest(http.MethodPost, "api/articles", strings.NewReader(string(cont)))
	if err != nil {
		return Article{}, err
	}
	req.Header.Add(APIKeyHeader, ar.API.Config.APIKey)
	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer res.Body.Close()

	if nonSuccessfulResponse(res) {
		return Article{}, unmarshalErrorResponse(res)
	}

	var a Article
	if err := json.NewDecoder(res.Body).Decode(&a); err != nil {
		return Article{}, err
	}
	return a, nil
}

// Update will mutate the resource by id, and all the changes
// performed to the Article will be applied, thus validation
// on the API side.
func (ar *ArticlesResource) Update(ctx context.Context, u ArticleUpdate, id uint32) (Article, error) {
	if ar.API.Config.InsecureOnly {
		return Article{}, ErrProtectedEndpoint
	}
	cont, err := json.Marshal(&u)
	if err != nil {
		return Article{}, err
	}
	req, err := ar.API.NewRequest(http.MethodPut, fmt.Sprintf("api/articles/%d", id), strings.NewReader(string(cont)))
	if err != nil {
		return Article{}, err
	}
	req.Header.Add(APIKeyHeader, ar.API.Config.APIKey)
	res, err := ar.API.HTTPClient.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer res.Body.Close()

	if nonSuccessfulResponse(res) {
		return Article{}, unmarshalErrorResponse(res)
	}

	var a Article
	if err := json.NewDecoder(res.Body).Decode(&a); err != nil {
		return Article{}, err
	}
	return a, nil
}

package zeptomail

import (
	"context"
	"net/http"
	"net/url"
)

type FileCache = Client

// FileCacheUploadAPI The API is used to upload files to File Cache
func (c *FileCache) FileCacheUploadAPI(ctx context.Context, req FileCacheUploadAPIReq) (*WrappedResponse[FileCacheUploadAPIRes], error) {
	path := "/files"
	query := url.Values{"name": []string{req.FileName}}
	header := http.Header{http.CanonicalHeaderKey("Content-Type"): {"text/plain"}}
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path, RawQuery: query.Encode()})
	return request[FileCacheUploadAPIReq, FileCacheUploadAPIRes](c, ctx, http.MethodPost, endpoint, header, req)
}

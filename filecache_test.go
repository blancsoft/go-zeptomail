package zeptomail_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/blancsoft/go-zeptomail"
)

func TestZeptoMailFileCache(t *testing.T) {
	zepto, err := zeptomail.NewZeptoMail(zeptoMailAgent, zeptoMailToken, "")
	require.NoError(t, err)

	rv, err := zepto.FileCache.FileCacheUploadAPI(t.Context(), zeptomail.FileCacheUploadAPIReq{
		FileName:    "test_filecache.ico",
		FileContent: fileAttachment,
	})
	require.NoError(t, err)

	t.Run("status code", func(t *testing.T) {
		assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
	})

	t.Run("valid response", func(t *testing.T) {
		assert.Nil(t, rv.Data.Error)

		assert.Equal(t, "OK", rv.Data.Message)
		assert.Equal(t, "file", rv.Data.Object)
		assert.NotEmpty(t, rv.Data.FileCacheKey)
		assert.NotNil(t, rv.Data.Data)
	})
}

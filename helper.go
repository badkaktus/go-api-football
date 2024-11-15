package gaf

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type HandlerHelper struct {
	Code            int
	ResponseBody    string
	ResponseHeaders map[string]int
}

func getHandler(t *testing.T, param *HandlerHelper) http.HandlerFunc {
	httpStatus := param.Code
	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers
		for key, value := range param.ResponseHeaders {
			w.Header().Set(key, strconv.Itoa(value))
		}

		w.WriteHeader(httpStatus)
		_, err := w.Write([]byte(param.ResponseBody))
		require.NoError(t, err)
	}
}

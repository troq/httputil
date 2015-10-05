package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type BadResponseStatusError struct {
	Resp *http.Response
}

func (e *BadResponseStatusError) Error() string {
	return fmt.Sprintf("GET response status not OK: %d", e.Resp.StatusCode)
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return body, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, &BadResponseStatusError{Resp: resp}
	}

	return body, nil
}

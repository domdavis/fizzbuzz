package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

func Get(urls []string, i int) (string, error) {
	s := ""

	for _, u := range urls {
		u, err := url.Parse(u)
		u.Path = path.Join(u.Path, strconv.Itoa(i))
		r, err := http.Get(u.String())

		if err != nil {
			return s, err
		}

		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			return s, err
		}

		s += string(b)
	}

	return s, nil
}

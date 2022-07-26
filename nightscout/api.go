package nightscout

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type NSRequest http.Request

type Instance struct {
	Url   *url.URL
	Token string
}

func GetInstance(instanceUrl string, token string) (*Instance, error) {
	parsedUrl, err := url.Parse(instanceUrl)
	if err != nil {
		return nil, err
	}

	return &Instance{
		Url:   parsedUrl,
		Token: token,
	}, nil
}

func GenerateFindKey(key string, operator string) string {
	var b strings.Builder
	b.WriteString("find[")
	b.WriteString(key)
	b.WriteString("]")

	if operator != "" {
		// ensure we're not doubling the $
		operator = strings.TrimPrefix(operator, "$")
		b.WriteString("[$")
		b.WriteString(operator)
		b.WriteString("]")
	}

	return b.String()
}

func (i *Instance) NewRequest(endpoint string) *NSRequest {
	api, _ := i.Url.Parse(endpoint)
	req, _ := http.NewRequest("GET", api.String(), nil)
	return (*NSRequest)(req)
}

func (r *NSRequest) AsRequest() *http.Request {
	return (*http.Request)(r)
}

func (r *NSRequest) Find(key, operator, value string) {
	fk := GenerateFindKey(key, operator)
	r.AddQuery(fk, value)
}

func (r *NSRequest) Count(count int) {
	r.SetQuery("count", strconv.Itoa(count))
}

func (r *NSRequest) AddToken(i *Instance) {
	if i.Token != "" {
		r.SetQuery("token", i.Token)
	}
}

func (r *NSRequest) SetQuery(key, value string) {
	query := r.URL.Query()
	query.Set(key, value)
	r.URL.RawQuery = query.Encode()
}

func (r *NSRequest) AddQuery(key, value string) {
	query := r.URL.Query()
	query.Add(key, value)
	r.URL.RawQuery = query.Encode()
}

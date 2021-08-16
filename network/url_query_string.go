package network

import "net/url"

//URLLinkQueryString URL拼接参数
func URLLinkQueryString(link string, data map[string]string) string {
	u, _ := url.Parse(link)
	q := u.Query()
	for k, v := range data {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

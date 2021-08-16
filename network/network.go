package network

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"pmo-test4.yz-intelligence.com/base/utils/network/constant"

	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type netWork struct {
	Link        string
	ContentType constant.ContentType
	Header      map[string]string
	Data        map[string]string
	Debug       bool
	sync.RWMutex
}

func New(link string) *netWork {
	return &netWork{
		Link:        link,
		ContentType: constant.FROM,
		Debug:       false,
	}
}

func (h *netWork) SetDebug(debug bool) *netWork {
	h.Debug = debug
	return h
}

//SetLink 更换请求地址
func (h *netWork) SetLink(link string) *netWork {
	h.Lock()
	defer h.Unlock()
	h.Link = link
	return h
}

//SetBody 设置数据包
func (h *netWork) SetBody(body map[string]string) *netWork {
	h.Lock()
	defer h.Unlock()
	h.Data = body
	return h
}

//SetHeader 设置请求头
func (h *netWork) SetHeader(header map[string]string) *netWork {
	h.Lock()
	defer h.Unlock()
	h.Header = header
	return h
}

//SetSendType 设置请求类型
func (h *netWork) SetSendType(sendType constant.ContentType) *netWork {
	h.Lock()
	defer h.Unlock()
	h.ContentType = sendType
	return h
}

func (h *netWork) Get() ([]byte, error) {
	return h.send(constant.GET.Value())
}

func (h *netWork) Post() ([]byte, error) {
	return h.send(constant.POST.Value())
}

func (h *netWork) send(method string) ([]byte, error) {
	var (
		httpRequest  *http.Request
		httpResponse *http.Response
		client       http.Client
		data         string
		err          error
	)

	if len(h.Data) > 0 {
		data, err = h.bodyHandler()
		if err != nil {
			return nil, err
		}
	}

	//忽略https的证书
	client.Transport = h.GetHttpTransport()

	httpRequest, err = http.NewRequest(method, h.Link, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer httpRequest.Body.Close()

	//if h.Debug {
	//
	//	if h.ContentType != constant.QueryString {
	//		requestPrintJSON, err := json.MarshalIndent(h.Data, "", "\t")
	//		if err != nil {
	//			return nil, err
	//		}
	//		h.Log.Debugf("content-Type: %s, requestURL: %s, data:\n %s", h.ContentType.Value(), h.Link, requestPrintJSON)
	//	}
	//
	//	if h.ContentType == constant.QueryString {
	//		h.Log.Debugf("content-Type: %s, requestURL: %s,  data: %s", h.ContentType, h.Link, data)
	//	}
	//
	//}

	h.setHeader(httpRequest)

	httpResponse, err = client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error http convert :%d", httpResponse.StatusCode))
	}

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		//h.Log.Errorf("parse json result failed. Exit by: %s", err.Error())
		return nil, err
	}

	//if h.Debug {
	//	var debugInterface interface{}
	//	if err := json.Unmarshal(responseBytes, &debugInterface); err != nil {
	//		return nil, err
	//	}
	//printJSON, err := json.MarshalIndent(debugInterface, "", "\t")
	//if err != nil {
	//	return nil, err
	//}
	//h.Log.Info(string(printJSON))
	//}

	return responseBytes, nil
}

func (h *netWork) setHeader(httpRequest *http.Request) {
	//设置默认header
	if len(h.Header) == 0 {
		if h.ContentType == constant.JSON {
			h.Header = map[string]string{
				"Content-Type": h.ContentType.Value(),
			}
		}
	}

	for k, v := range h.Header {
		if strings.ToLower(k) == "host" {
			httpRequest.Host = v
		} else {
			httpRequest.Header.Add(k, v)
		}
	}
}

//GetHttpTransport 设置https证书
func (h *netWork) GetHttpTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

//bodyHandler
func (h *netWork) bodyHandler() (string, error) {
	if h.ContentType == constant.JSON {
		sendBody, err := json.Marshal(h.Data)
		if err != nil {
			return "", err
		}
		return string(sendBody), nil
	} // 如果是JSON格式的

	if h.ContentType == constant.QueryString {
		h.SetLink(URLLinkQueryString(h.Link, h.Data))
		return "", nil
	}

	body := http.Request{}
	if err := body.ParseForm(); err != nil {
		return "", err
	}

	for k, v := range h.Data {
		body.Form.Add(k, v)
	}

	return body.Form.Encode(), nil
}

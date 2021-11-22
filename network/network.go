package network

//
//type netWork struct {
//	Link string
//
//	method      constant.Method
//	contentType constant.ContentType
//
//	Header      map[string]string
//	formData    map[string]string
//	requestBody map[string]interface{}
//	queryString map[string]string
//
//	data.I
//
//	Config Config
//
//	sync.RWMutex
//}
//
//type Config struct {
//	debug bool        // 是否打印
//	log   *zap.Logger // 日志
//}
//
////New 创建一个网络请求
//func New(link string, options ...Options) *netWork {
//	network := &netWork{
//		Link:        link,
//		contentType: constant.FORM,
//		Config: Config{
//			debug: false,
//			log:   DefaultLogger(),
//		},
//	}
//
//	// Run the options on it
//	for _, option := range options {
//		option(network)
//	}
//
//	return network
//}
//
////Options 配置项
//type Options func(*netWork)
//
////SetLogger 日志
//func SetLogger(logger *zap.Logger) Options {
//	return func(work *netWork) {
//		work.Config.log = logger
//	}
//}
//
////SetDebug 设置是否打印参数
//func SetDebug(debug bool) Options {
//	return func(work *netWork) {
//		work.Config.debug = debug
//	}
//}
//
////SetLink 更换请求地址
//func (h *netWork) SetLink(link string) *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.Link = link
//	return h
//}
//
////Form 设置数据包
//func (h *netWork) Form(data map[string]string) *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.formData = data
//	h.contentType = constant.FORM
//	return h
//}
//
////Body 设置
//func (h *netWork) Body(data map[string]interface{}) *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.requestBody = data
//	h.contentType = constant.JSON
//	return h
//}
//
//func (h *netWork) Post() *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.method = constant.POST
//	return h
//}
//
////QueryString URL查询字符
//func (h *netWork) QueryString(query map[string]string) *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.queryString = query
//	h.contentType = constant.QueryString
//	return h
//}
//
////SetHeader 设置请求头
//func (h *netWork) SetHeader(header map[string]string) *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.Header = header
//	return h
//}
//
////Get 请求方式为GET
//func (h *netWork) Get() *netWork {
//	h.Lock()
//	defer h.Unlock()
//	h.method = constant.GET
//	return h
//}
//
//func (h *netWork) Do() ([]byte, error) {
//	var (
//		httpRequest  *http.Request
//		httpResponse *http.Response
//		client       http.Client
//	)
//
//	time.Now()
//	parameter, err := h.data()
//	if err != nil {
//		return nil, err
//	}
//
//	//忽略https的证书
//	client.Transport = h.GetHttpTransport()
//
//	httpRequest, err = http.NewRequest(h.method.Value(), h.Link, parameter)
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		_ = httpRequest.Body.Close()
//	}()
//
//	if h.Config.debug {
//		if h.contentType == constant.FORM || h.contentType == constant.JSON {
//			var requestPrintJSON []byte
//
//			if h.contentType == constant.FORM {
//				requestPrintJSON, err = json.MarshalIndent(h.formData, "", "\t")
//				if err != nil {
//					return nil, err
//				}
//			} else if h.contentType == constant.JSON {
//				requestPrintJSON, err = json.MarshalIndent(h.requestBody, "", "\t")
//				if err != nil {
//					return nil, err
//				}
//			}
//
//			h.Config.log.Info("FormData", zap.String("content-Type:", h.contentType.Value()), zap.String("requestURL", h.Link), zap.String("parameter", string(requestPrintJSON)))
//		} else if h.contentType == constant.QueryString {
//			h.Config.log.Info("QueryString - ", zap.String("Content-Type", h.contentType.Value()), zap.String("requestURL", h.Link))
//		}
//
//	}
//	h.setHeader(httpRequest)
//
//	httpResponse, err = client.Do(httpRequest)
//	if err != nil {
//		return nil, err
//	}
//	defer func() {
//		_ = httpResponse.Body.Close()
//	}()
//
//	if httpResponse.StatusCode != http.StatusOK {
//		return nil, errors.New(fmt.Sprintf("error http convert :%d", httpResponse.StatusCode))
//	}
//
//	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	if h.Config.debug {
//		var debugInterface interface{}
//		if err := json.Unmarshal(responseBytes, &debugInterface); err != nil {
//			return nil, err
//		}
//		printJSON, err := json.MarshalIndent(debugInterface, "", "\t")
//		if err != nil {
//			return nil, err
//		}
//		h.Config.log.Info("RESULT = \n" + string(printJSON))
//	}
//
//	return responseBytes, nil
//}
//
//func (h *netWork) setHeader(httpRequest *http.Request) {
//	//设置默认header
//	if len(h.Header) == 0 {
//		if h.contentType == constant.JSON {
//			h.Header = map[string]string{
//				"Content-Type": h.contentType.Value(),
//			}
//		}
//	}
//
//	for k, v := range h.Header {
//		if strings.ToLower(k) == "host" {
//			httpRequest.Host = v
//		} else {
//			httpRequest.Header.Add(k, v)
//		}
//	}
//}
//
////GetHttpTransport 设置https证书
//func (h *netWork) GetHttpTransport() *http.Transport {
//	return &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//}
//
////data
//func (h *netWork) data() (*strings.Reader, error) {
//	if h.contentType == constant.JSON {
//		sendBody, err := json.Marshal(h.requestBody)
//		if err != nil {
//			return nil, err
//		}
//		return strings.NewReader(string(sendBody)), nil
//	} // 如果是JSON格式的
//
//	if h.contentType == constant.QueryString {
//		queryString := URLLinkQueryString(h.Link, h.queryString)
//		h.SetLink(queryString)
//		return strings.NewReader(""), nil
//	}
//
//	body := http.Request{}
//	if err := body.ParseForm(); err != nil {
//		return strings.NewReader(""), err
//	}
//
//	for k, v := range h.formData {
//		body.Form.Add(k, v)
//	}
//
//	return strings.NewReader(body.Form.Encode()), nil
//}

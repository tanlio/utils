package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

const defaultUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 18_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Mobile/15E148 Safari/604.1"

type RequestOption struct {
	Insecure  bool
	Timeout   time.Duration
	UserAgent string
}

type Option func(*RequestOption)

func WithInsecure() Option {
	return func(o *RequestOption) { o.Insecure = true }
}
func WithContextTimeout(d time.Duration) Option {
	return func(o *RequestOption) { o.Timeout = d }
}
func WithUserAgent(ua string) Option {
	return func(o *RequestOption) { o.UserAgent = ua }
}

type HTTPClient struct {
	secure   *http.Client
	insecure *http.Client
	ua       string
}

var (
	once       sync.Once
	httpClient *HTTPClient
)

func DefaultClient() *HTTPClient {
	once.Do(func() {
		secureTr := newTransport(false)
		insecureTr := newTransport(true)

		httpClient = &HTTPClient{
			secure: &http.Client{
				Timeout:   0,
				Transport: secureTr,
			},
			insecure: &http.Client{
				Timeout:   0,
				Transport: insecureTr,
			},
			ua: defaultUserAgent,
		}
	})
	return httpClient
}

func newTransport(insecure bool) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   20 * time.Second, // 连接超时
		KeepAlive: 30 * time.Second,
	}
	tr := &http.Transport{
		DialContext:           dialer.DialContext,
		IdleConnTimeout:       60 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if insecure {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return tr
}

func (h *HTTPClient) pick(opt RequestOption) *http.Client {
	if opt.Insecure {
		return h.insecure
	}
	return h.secure
}

func (h *HTTPClient) applyHeaders(req *http.Request, header map[string]string, ua string) {
	if ua == "" {
		ua = h.ua
	}
	req.Header.Set("User-Agent", ua)
	for k, v := range header {
		req.Header.Set(k, v)
	}
}

func normalizeParam(param any) (map[string]any, []byte, error) {
	if param == nil {
		return nil, nil, nil
	}

	switch p := param.(type) {
	case map[string]any:
		return p, nil, nil
	case map[string]string:
		m := make(map[string]any, len(p))
		for k, v := range p {
			m[k] = v
		}
		return m, nil, nil
	case []byte:
		return nil, param.([]byte), nil
	default:
		return nil, nil, errors.New("param must be map[string]any or map[string]string")
	}
}

func (h *HTTPClient) Do(ctx context.Context, req *http.Request, header map[string]string, opts ...Option) (int, []byte, error) {
	opt := RequestOption{UserAgent: h.ua, Timeout: 30 * time.Second}
	for _, fn := range opts {
		fn(&opt)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	if opt.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opt.Timeout)
		defer cancel()
	}
	req = req.WithContext(ctx)

	h.applyHeaders(req, header, opt.UserAgent)

	resp, err := h.pick(opt).Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	return resp.StatusCode, body, readErr
}

func (h *HTTPClient) RequestJSON(ctx context.Context, method, uri string, param any, header map[string]string, opts ...Option) (int, []byte, error) {
	pMap, pBody, err := normalizeParam(param)
	if err != nil {
		return 0, nil, err
	}
	var body io.Reader
	if pMap != nil {
		data, _ := json.Marshal(pMap)
		body = bytes.NewReader(data)
	} else if pBody != nil {
		body = bytes.NewReader(pBody)
	}

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return h.Do(ctx, req, header, opts...)
}

func (h *HTTPClient) RequestForm(ctx context.Context, method, uri string, param any, header map[string]string, opts ...Option) (int, []byte, error) {
	pMap, pBody, err := normalizeParam(param)
	if err != nil {
		return 0, nil, err
	}

	values := url.Values{}
	if pMap != nil {
		for k, v := range pMap {
			values.Set(k, fmt.Sprint(v))
		}
	} else if pBody != nil {
		paramBody := make(map[string]any)
		json.Unmarshal(pBody, &paramBody)
		for k, v := range paramBody {
			var tempStr string
			switch v.(type) {
			case float32:
				tempStr = decimal.NewFromFloat32(v.(float32)).String()
			case float64:
				tempStr = decimal.NewFromFloat(v.(float64)).String()
			default:
				tempStr = fmt.Sprint(v)
			}
			values.Set(k, tempStr)
		}
	}

	req, err := http.NewRequest(method, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return h.Do(ctx, req, header, opts...)
}

func (h *HTTPClient) RequestGet(ctx context.Context, method, rawURL string, param any, header map[string]string, opts ...Option) (int, []byte, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return 0, nil, err
	}

	pMap, _, err := normalizeParam(param)
	if err != nil {
		return 0, nil, err
	}

	q := u.Query()
	for k, v := range pMap {
		q.Set(k, fmt.Sprint(v))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return 0, nil, err
	}
	return h.Do(ctx, req, header, opts...)
}

func (h *HTTPClient) RequestFile(ctx context.Context, method, uri string, param any, fileFieldName string, file *os.File, header map[string]string, opts ...Option) (int, []byte, error) {
	buf := &bytes.Buffer{}
	bw := multipart.NewWriter(buf)

	pMap, pBody, err := normalizeParam(param)
	if err != nil {
		return 0, nil, err
	}
	for k, v := range pMap {
		if err := bw.WriteField(k, fmt.Sprint(v)); err != nil {
			return 0, nil, err
		}
	}
	paramBody := make(map[string]any)
	json.Unmarshal(pBody, &paramBody)
	for k, v := range paramBody {
		var tempStr string
		switch v.(type) {
		case float32:
			tempStr = decimal.NewFromFloat32(v.(float32)).String()
		case float64:
			tempStr = decimal.NewFromFloat(v.(float64)).String()
		default:
			tempStr = fmt.Sprint(v)
		}
		if err := bw.WriteField(k, tempStr); err != nil {
			return 0, nil, err
		}
	}

	if file != nil && fileFieldName != "" {
		fw, err := bw.CreateFormFile(fileFieldName, filepath.Base(file.Name()))
		if err != nil {
			return 0, nil, err
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return 0, nil, err
		}
		if _, err = io.Copy(fw, file); err != nil {
			return 0, nil, err
		}
	}

	if err := bw.Close(); err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", bw.FormDataContentType())
	return h.Do(ctx, req, header, opts...)
}

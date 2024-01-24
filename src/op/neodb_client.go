package op

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/golang/glog"

	"github.com/andrewrong/neodb-go-api/object"
	"github.com/andrewrong/neodb-go-api/util"
)

type NeoClient struct {
	client    *util.HttpClient
	token     string
	timeoutMs int
	url       string // 域名或者url
	MaxRetry  int
}

func NewNeoClient(url, token string, timeoutMs int) *NeoClient {
	tr := http.DefaultTransport
	return &NeoClient{
		client:    util.NewHttpClient(token, timeoutMs, tr),
		token:     token,
		timeoutMs: timeoutMs,
		url:       url,
		MaxRetry:  3,
	}
}

func (c *NeoClient) get(url string) (int, []byte, error, bool) {
	return func() (int, []byte, error, bool) {
		resp, err := c.client.Get(url)
		if err != nil {
			log.Errorf("请求失败:%++v", err)
			return util.HttpStatus500.Code, nil, err, true
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		if resp.StatusCode == util.HttpStatus401.Code {
			log.Error("请求失败,token无效")
			return util.HttpStatus401.Code,nil, err, false
		}

		if resp.StatusCode == util.HttpStatus400.Code {
			log.Error("请求失败,参数错误")
			return nil, err, false
		}

		if resp.StatusCode == util.HttpStatus403.Code {
			log.Error("请求失败,没有权限")
			return nil, err, false
		}

		if resp.StatusCode == util.HttpNotFound.Code {
			log.Error("请求失败,资源不存在")
			return nil, err, false
		}

		if resp.StatusCode == util.HttpStatus500.Code {
			log.Error("Service Unavailable, maybe retry")
			return nil, err,true
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("读取响应内容失败:", err)
			return nil, err, true
		}
		return body, nil, false
	}()
}

func (c *NeoClient) GetMe() (*object.User, error) {
	url := c.url + "/api/me"
	var err error = nil

	for i := 0; i < c.MaxRetry; i++ {

		if err != nil {
			if retry {
				continue
			}
			return nil, err
		}

		var user object.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Errorf("解析响应内容失败:%+v", err)
			return nil, err
		}
		return &user, nil
	}
	return nil, err
}

func (c *)

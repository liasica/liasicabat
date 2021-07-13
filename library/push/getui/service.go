package getui

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	UriToken = "/auth"
)

type service struct {
	appId           string
	appKey          string
	appSecret       string
	appMasterSecret string
	baseUrl         string
}

var serv *service

func Service() *service {
	if serv == nil {
		serv = &service{
			appId:           g.Cfg().GetString("getui.appId"),
			appKey:          g.Cfg().GetString("getui.appKey"),
			appSecret:       g.Cfg().GetString("getui.appSecret"),
			appMasterSecret: g.Cfg().GetString("getui.appMasterSecret"),
			baseUrl:         fmt.Sprintf("https://restapi.getui.com/v2/%s", g.Cfg().GetString("getui.appId")),
		}
	}
	return serv
}

func (s *service) Rpc(method, uri string, data interface{}) ([]byte, error) {
	resp, err := g.Client().SetHeader("Content-Type", "application/json").DoRequest(method, fmt.Sprintf("%s%s", s.baseUrl, uri), data)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}

func (s *service) Token() (string, error) {
	t := strconv.FormatInt(gtime.TimestampMilli(), 10)
	h := sha256.New()
	_, err := io.WriteString(h, s.appKey+t+s.appMasterSecret)
	if err != nil {
		return "", err
	}
	content, err := s.Rpc(http.MethodPost, UriToken, TokenRequest{
		Sign:      fmt.Sprintf("%x", h.Sum(nil)),
		Timestamp: t,
		AppKey:    s.appKey,
	})
	if err != nil {
		return "", err
	}
	var res TokenResponse
	if _err := json.Unmarshal(content, &res); _err != nil {
		return "", _err
	}
	if res.Code != 0 {
		return "", errors.New(res.Msg)
	}
	return res.Data.Token, nil
}

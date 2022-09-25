package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
)

func DownloadFile(fileUrl string) (fileName string, err error) {
	parsedUrl, err := url.Parse(fileUrl)
	if err != nil {
		return "", err
	}
	scheme := strings.ToLower(parsedUrl.Scheme)
	if scheme != "http" && scheme != "https" {
		err = errors.New("scheme is invalid. url: " + fileUrl)
		return "", err
	}
	if scheme == "https" {
		fileUrl = strings.Replace(fileUrl, "https", "http", -1)
	}

	endPoints, err := LookupIPAddr(parsedUrl.Host, 5*time.Minute)
	if err != nil {
		return "", err
	}

	index := rand.Intn(len(endPoints))
	ipInt := Ip2Int64(endPoints[index])
	if IP_127 == (ipInt>>24) ||
		IP_10 == (ipInt>>24) ||
		IP_172 == (ipInt>>20) ||
		IP_192 == (ipInt>>16) {
		err = errors.New("not found: " + fileUrl)
		return "", err
	}

	request := httplib.Get(fileUrl).SetProxy(func(request *http.Request) (*url.URL, error) {
		u, _ := url.ParseRequestURI("http://" + endPoints[index] + ":80")
		log.Println("Proxy ipAddr:", u)
		return u, nil
	})
	imageData, err := request.SetTimeout(time.Second*2, time.Second*20).Bytes()
	if err != nil {
		DeleteInvalidIpAddr(parsedUrl.Host, endPoints[index])
		return "", err
	}

	fileName = fmt.Sprintf("img-%s", path.Base(parsedUrl.Path))
	err = WriteFile(fileName, imageData)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func WriteFile(fileName string, data []byte) (err error) {
	fileName = strings.Replace(fileName, "\\", "/", -1)
	fileName = path.Join("/tmp/img/", fileName)
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	staticFolder = "static"
	jsonFolder   = "json"
	certFolder   = "cert"
	configFolder = "conf"
)

func SendRequest(url string, headers map[string]string, data interface{}) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	client := &http.Client{
		Timeout: time.Second * 50,
	}
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
	}
	// 设置请求头（如果有）
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCurrentPath 获取当前执行文件的路径
func GetCurrentPath() string {
	path := os.Getenv("BASE_PATH")
	if path == "" {
		if dir, err := os.Getwd(); err != nil {
			return ""
		} else {
			return dir
		}
	}
	return path
}

// GetStaticPath 获取静态文件路径
func GetStaticPath() string {
	return filepath.Join(GetCurrentPath(), staticFolder)
}

// IsExist 判断某个文件是否存在
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil || os.IsExist(err)
}

// IsDirEmpty checks if a directory is empty
func IsDirEmpty(dirPath string) bool {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false
	}
	return len(entries) == 0
}

// CreateDir create folder if not exist
func CreateDir(dirPath string) error {
	// check folder is existing
	if _, err := os.Stat(dirPath); err != nil {
		err := os.MkdirAll(dirPath, 0o755)
		if err != nil {
			return err
		}
	}
	return nil
}

func JsonPath() string {
	return path.Join(GetCurrentPath(), staticFolder, jsonFolder)
}

func CertPath() string {
	return path.Join(GetCurrentPath(), certFolder)
}

func ConfigPath() string {
	return path.Join(GetCurrentPath(), staticFolder, configFolder)
}

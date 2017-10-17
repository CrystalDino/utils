package utils

import (
	"bytes"
	"errors"
	"path"
	"strings"

	"github.com/dchest/captcha"
)

var (
	CaptchaWidth  int = 240
	CaptchaHeight int = 80
	CaptchaLength int = 6
	//support value; "en", "ja", "ru", "zh"
	CaptchaLang string = "en"
)

func InitCaptcha(width, height, defaultLen int, lang string) {
	if width > 0 {
		CaptchaWidth = width
	}
	if height > 0 {
		CaptchaHeight = height
	}
	if len(lang) == 2 {
		CaptchaLang = lang
	}
	if defaultLen >= 4 {
		CaptchaLength = defaultLen
	}
}

func CreateCaptchaId() string {
	return captcha.NewLen(CaptchaLength)
}

func GetCaptcha(name string) (ctt []byte, err error) {
	ctt = make([]byte, 0)
	if name == "" {
		err = errors.New("captcha name is empty")
		return
	}
	_, file := path.Split(name)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if id == "" {
		err = errors.New("captcha id is empty")
		return
	}
	var content bytes.Buffer
	captcha.Reload(id)
	switch strings.ToLower(ext) {
	case ".png":
		if err = captcha.WriteImage(&content, id, CaptchaWidth, CaptchaHeight); err != nil {
			return
		}
	case ".wav":
		if err = captcha.WriteAudio(&content, id, CaptchaLang); err != nil {
			return
		}
	default:
		err = errors.New("type of captcha name is error")
		return
	}
	ctt = content.Bytes()
	return
}

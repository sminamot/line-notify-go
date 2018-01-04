package linenotify

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiUrl = "https://notify-api.line.me/api/notify"

type Line struct {
	token string
}

func NewNotify() *Line {
	return &Line{}
}

func (l *Line) SetToken(t string) {
	l.token = t
}

func (l *Line) Notify(m string) error {
	if l.token == "" {
		return fmt.Errorf("must set token with SetToken()")
	}
	data := url.Values{"message": {m}}
	r, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", "Bearer "+l.token)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("request failed")
	}

	return nil
}

package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendSimpleAlert(topic, message string) error {
	url := "http://ntfy.sh/" + topic
	res, err := http.Post(url, "text/plain", bytes.NewBufferString(message))
	if err != nil {
		return fmt.Errorf("some error in ntfy")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("status code: %v", res.Status)
	}
	return nil
}

func SendRichAlert(topic, title, message string) error {
	url := "http://ntfy.sh/" + topic
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(message))
	if err != nil {
		return err
	}
	req.Header.Set("Title", title)
	req.Header.Set("Tags", "alarm_clock")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func FormatClassAlert(class Class, minutesUntil int) (string, string) {
	var title, message string
	title = fmt.Sprintf("Class in %d Minutes", minutesUntil)
	message = fmt.Sprintf("%s class starting soon\nRoom: %s\nProfessor: %s", class.subject, class.room, class.professor)
	return title, message
}

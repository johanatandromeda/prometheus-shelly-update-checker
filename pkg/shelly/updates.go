package shelly

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func UpdateNeeded(target string) (bool, int, error) {

	if strings.HasSuffix(target, "status") {
		slog.Debug(fmt.Sprintf("Checking update status for '%s' as Shelly V1", target))
		return updateNeededV1(target)
	} else {
		slog.Debug(fmt.Sprintf("Checking update status for '%s' as Shelly V2", target))
		return updateNeededV2(target)
	}
}

func updateNeededV1(target string) (bool, int, error) {

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := c.Get(target)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error making GET request to %s: %s", target, err))
		return false, 500, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		return false, resp.StatusCode, nil
	}
	j, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error reading response body: %s", err))
		return false, 500, err
	}
	var d interface{}
	err = json.Unmarshal(j, &d)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error unmarshalling response body: %s", err))
		return false, 500, err
	}
	update := d.(map[string]interface{})["update"].(map[string]interface{})["has_update"].(bool)
	slog.Debug(fmt.Sprintf("Update needed: %t", update))
	return update, resp.StatusCode, nil

}

func updateNeededV2(target string) (bool, int, error) {

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := c.Get(target)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error making GET request to %s: %s", target, err))
		return false, 500, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		return false, resp.StatusCode, nil
	}
	j, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error reading response body: %s", err))
		return false, 500, err
	}
	var d interface{}
	err = json.Unmarshal(j, &d)
	if err != nil {
		slog.Debug(fmt.Sprintf("Error unmarshalling response body: %s", err))
		return false, 500, err
	}
	update := d.(map[string]interface{})["available_updates"].(map[string]interface{})["stable"] != nil
	slog.Debug(fmt.Sprintf("Update needed: %t", update))
	return update, resp.StatusCode, nil

}

package btccli

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
)

func Call[T any](conf, command string, args ...string) (T, error) {
	args = append([]string{"--conf=" + conf, command}, args...)
	out, err := exec.Command("bitcoin-cli", args...).Output()
	var res T
	if err != nil {
		return res, err
	}

	switch any(res).(type) {
	case string:
		res = any(strings.Split(string(out), "\n")[0]).(T)
	case int:
		_res, err := strconv.Atoi(strings.Split(string(out), "\n")[0])
		if err != nil {
			return res, err
		}
		res = any(_res).(T)
	default:
		err = json.Unmarshal(out, &res)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

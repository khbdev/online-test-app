package discovery

import (
	"fmt"
	"geteway-service/internal/util"

	"strings"
)


func GetServiceAddress(serviceName string) (string, error) {
	envKey := strings.ToUpper(strings.ReplaceAll(serviceName, "-", "_"))

	addr := util.LoadEnv(envKey)
	if addr == "" {
		return "", fmt.Errorf("%s env topilmadi", envKey)
	}
	return addr, nil
}

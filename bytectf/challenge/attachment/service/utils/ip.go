package utils

import (
	"strconv"
	"strings"
)

const (
	IP_127 int64 = 127
	IP_10  int64 = 10
	IP_172 int64 = 2753
	IP_192 int64 = 49320
)

func Ip2Int64(ip string) (sum int64) {
	if ip == "" {
		return
	}
	bits := strings.Split(ip, ".")
	if len(bits) != 4 {
		return
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return
}

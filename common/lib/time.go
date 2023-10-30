package lib

import "time"

// StrToTime 解析 YYYY-mm-dd HH:ii:ss格式时间为13位时间戳
func StrToTime(str string) (int64, error) {
	parse, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return 0, err
	}
	return parse.Unix() * 1000, nil
}

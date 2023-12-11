package Utils

import (
	"encoding/base64"
	"fmt"
	"time"
)

func ID() string {
	dTime := time.Now()

	year := dTime.Year()
	month := int(dTime.Month())
	day := dTime.Day()
	hour := dTime.Hour()
	minute := dTime.Minute()
	second := dTime.Second()
	nanosecond := dTime.Nanosecond()

	data := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%09d", year, month, day, hour, minute, second, nanosecond)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
    ID := encoded+".zip" 
	return ID
}
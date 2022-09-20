package digitalclock

import (
	"fmt"
	"strconv"
	"time"
)

// AM defines the AM string format
const AM = "A"

// PM defines the PM string format
const PM = "P"
const minRowsForBorder = 3

// Converts integer to string along with makes sure the length of string is > 2
func intStrConv(val int) string {
	valStr := strconv.Itoa(val)

	if len(valStr) < 2 {
		valStr = "0" + valStr
	}
	return valStr
}

// Returns Hour + minute + AM/PM information based on the settings
func getHourMinute(hourFormat string) string {
	strHours := intStrConv(time.Now().Hour())
	AMPM := " "

	if hourFormat == "12" {
		hour := time.Now().Hour()
		strHours = intStrConv(hour % 12)
		if (hour % 12) == hour {
			AMPM = AM
		} else {
			AMPM = PM
		}

	}

	strMinutes := intStrConv(time.Now().Minute())
	strMinutes += AMPM
	return strHours + getColon() + strMinutes
}

// Returns the : with blinking based on the seconds
func getColon() string {
	if time.Now().Second()%2 == 0 {
		return ":"
	}
	return " "
}

func getDate(dateFormat string, withDatePrefix bool) string {
	if withDatePrefix {
		return fmt.Sprintf("Date: %s", time.Now().Format(dateFormat))
	}
	return time.Now().Format(dateFormat)
}

func getUTC() string {
	return fmt.Sprintf("UTC: %s", time.Now().UTC().Format(time.RFC3339))
}

func getEpoch() string {
	return fmt.Sprintf("Epoch: %d", time.Now().Unix())
}

// Renders the clock as string by accessing appropriate font from configured in settings
func renderClock(widgetSettings Settings) (string, bool) {
	var digFont ClockFont
	clockTime := getHourMinute(widgetSettings.hourFormat)
	digFont = getFont(widgetSettings)

	chars := [][]string{}
	for _, char := range clockTime {
		chars = append(chars, digFont.get(string(char)))
	}

	needBorder := digFont.fontRows <= minRowsForBorder
	return fontsJoin(chars, digFont.fontRows, widgetSettings.color), needBorder
}

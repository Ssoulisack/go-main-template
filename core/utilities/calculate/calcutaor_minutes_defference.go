package utilitiesCalculate

import (
	"fmt"
	"math"
	"time"
)

// CalculateMinutesDifference calculates the difference in minutes between two timestamps
func CalculateMinutesDifference(start, end string) (float64, error) {
	// Define the time format
	layout := "2006-01-02 15:04:05.000"
	fmt.Println(start)
	fmt.Println(end)
	// Parse the start and end times
	startTime, err := time.Parse(layout, start)
	if err != nil {
		return 0, err
	}

	endTime, err := time.Parse(layout, end)
	if err != nil {
		return 0, err
	}

	// Calculate the difference in seconds
	diff := endTime.Sub(startTime).Seconds()

	// Convert seconds into minutes and remaining seconds
	minutes := diff / 60

	// Round the seconds to 2 decimal places
	roundedSeconds := math.Round(minutes*100) / 100

	return roundedSeconds, nil
}

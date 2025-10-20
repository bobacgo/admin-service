package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timeStr := "08:00"
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		t.Errorf("Failed to parse time: %v", err)
	}
	fmt.Println(parsedTime)
}
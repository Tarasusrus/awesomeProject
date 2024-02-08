package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"testing"
	"time"
)

func TestGetNTPTime(t *testing.T) {
	ntpTime, err := getNTPTime(ntpServerURL)
	if err != nil {
		t.Fatalf("getNTPTime() returned error: %v", err)
	}

	if ntpTime.IsZero() {
		t.Fatal("getNTPTime() returned zero time")
	}

	if time.Since(ntpTime) > time.Minute {
		t.Fatal("getNTPTime() returned time that was too far from current time")
	}
}

func getNTPTime(server string) (time.Time, error) {
	ntpTime, err := ntp.Time(server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NTP error: %v\n", err)
		os.Exit(1)
	}

	return ntpTime, nil
}

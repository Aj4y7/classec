package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Class struct {
	section    string
	day        string
	subject    string
	start_time string
	end_time   string
	room       string
	professor  string
}

type Timetable []Class

func LoadTimetableFromCSV(filepath string) (Timetable, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}

	var timetable Timetable

	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 7 {
			return nil, fmt.Errorf("record %d has %d fields, want at least 7", i+1, len(record))
		}
		class := Class{
			section:    record[0],
			day:        record[1],
			subject:    record[2],
			start_time: record[3],
			end_time:   record[4],
			room:       record[5],
			professor:  record[6],
		}

		timetable = append(timetable, class)

	}

	return timetable, nil
}

func GetTodayClassesForSection(tt Timetable, section string) []Class {
	dayName := time.Now().Weekday().String()
	dayNameFirst3 := dayName[0:3]
	var todayClasses []Class
	for i := 0; i < len(tt); i++ {
		if tt[i].day == dayNameFirst3 && tt[i].section == section {
			todayClasses = append(todayClasses, tt[i])
		}
	}

	return todayClasses
}

func ParseClassTime(timeStr string) (time.Time, error) {
	now := time.Now()
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location()), nil
}

func GetNextClass(classes []Class) (*Class, int, error) {
	if len(classes) == 0 {
		return nil, 0, fmt.Errorf("no classes provided")
	}
	now := time.Now()
	var nextClass *Class
	minDiff := 9999

	for i := 0; i < len(classes); i++ {
		classTime, err := ParseClassTime(classes[i].start_time)
		if err != nil {
			return nil, 0, fmt.Errorf("invalid class")
		}
		timeDiffInMin := int(classTime.Sub(now).Minutes())
		if timeDiffInMin < 13 || timeDiffInMin > 16 {
			continue
		}
		minDiff = min(minDiff, timeDiffInMin)
		if minDiff == timeDiffInMin {
			nextClass = &classes[i]
		}
	}

	if nextClass == nil {
		return nil, 0, fmt.Errorf("no upcoming class")
	}
	return nextClass, minDiff, nil
}

package main

import (
	"log"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	sentAlerts = make(map[string]time.Time)
	alertMutex sync.Mutex
)

func StartScheduler() {
	log.Println("Starting College Alert System...")
	log.Println("Checking for classes every minute")

	c := cron.New()
	c.AddFunc("* * * * *", func() {
		// Clean old alerts from memory (older than 2 hours)
		alertMutex.Lock()
		for key, sentTime := range sentAlerts {
			if time.Since(sentTime) > 2*time.Hour {
				delete(sentAlerts, key)
			}
		}
		alertMutex.Unlock()

		tt, err := LoadTimetableFromCSV("timetable.csv")
		if err != nil {
			log.Printf("ERROR: Failed to load timetable: %v", err)
			return
		}
		config := GetConfig()
		for _, section := range config.Sections {
			todayClasses := GetTodayClassesForSection(tt, section)
			nextClass, minutesUntil, err := GetNextClass(todayClasses)
			if err != nil {
				continue
			}

			// Check if we already sent an alert for this class
			alertKey := section + "-" + nextClass.subject + "-" + nextClass.start_time
			alertMutex.Lock()
			lastSent, exists := sentAlerts[alertKey]
			if exists && time.Since(lastSent) < 10*time.Minute {
				alertMutex.Unlock()
				continue // Already sent alert for this class recently
			}
			sentAlerts[alertKey] = time.Now()
			alertMutex.Unlock()

			topic := config.GetTopicForSection(section)
			title, message := FormatClassAlert(*nextClass, minutesUntil)
			err = SendRichAlert(topic, title, message)
			if err != nil {
				log.Printf("ERROR: Failed to send alert to %s: %v", topic, err)
			} else {
				log.Printf("✓ Sent alert to %s: %s in %d minutes", topic, nextClass.subject, minutesUntil)
			}

		}
	})
	c.Start()
	log.Println("✓ Scheduler started successfully!")
	select {}

}

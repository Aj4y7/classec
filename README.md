# ClassEc- College Class Alert System

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> Never miss a college class again! Automated notifications sent 15 minutes before each class starts.

ClassEc is a lightweight Go application that monitors your college timetable and sends push notifications to students' phones via [ntfy.sh](https://ntfy.sh) - a free, open-source notification service that requires **no account or setup**.

## Features

- **Automatic Notifications** - Alerts sent 15 minutes before class
- **Mobile & Desktop** - Works on iOS, Android, and web browsers
- **Section-Specific** - Each section gets its own notification channel
- **No Duplicates** - Smart deduplication prevents spam
- **Completely Free** - No API keys, no credit card, no signup required
- **Memory Efficient** - Automatic cleanup of old data
- **Easy to Deploy** - Single binary, runs anywhere

## How It Works

1. **Load Timetable** - Reads class schedule from CSV file
2. **Monitor** - Checks every minute for upcoming classes
3. **Notify** - Sends push notification to section-specific topics
4. **Students Subscribe** - Students subscribe to their section (e.g., `classec-A1`)

## Quick Start

### Prerequisites

- Go 1.20 or higher
- A timetable CSV file (format below)

### Installation

```bash
# Clone the repository
git clone https://github.com/Aj4y7/classec.git
cd classec

# Install dependencies
go mod download

# Run the application
go run .
```

The scheduler will start and check for classes every minute!

## Timetable CSV Format

Create a `timetable.csv` file with the following columns:

```csv
section,day,subject,start_time,end_time,room,professor
A1,Mon,MPWS,08:30,10:30,DG-04,MPS
A1,Mon,CAMD,10:30,11:30,CG-09,SS
A2,Mon,CAMD,08:30,10:30,CS-17,SS
B1,Tue,BEE,08:30,09:30,CG-09,NKK
```

**Important:**

- Days must be 3-letter abbreviations: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`, `Sun`
- Times must be in 24-hour format: `HH:MM`
- Section names must match what you configure in `config.go`

## For Students: How to Subscribe

### On Mobile (Android/iOS)

1. **Install ntfy app**

   - Android: [Google Play Store](https://play.google.com/store/apps/details?id=io.heckel.ntfy)
   - iOS: [App Store](https://apps.apple.com/app/ntfy/id1625396347)

2. **Subscribe to your section**
   - Open the app
   - Tap "+" to add subscription
   - Enter topic: `classec-A1` (replace A1 with your section)
   - Done! You'll now receive notifications

### On Web Browser

Visit [ntfy.sh](https://ntfy.sh) and subscribe to `classec-A1` (replace with your section)

### Example Topics

- Section A1: `classec-A1`
- Section A2: `classec-A2`
- Section B1: `classec-B1`
- And so on...

## Configuration

Edit `config.go` to customize:

```go
func GetConfig() Config {
    return Config{
        NtfyServerURL:    "https://ntfy.sh",        // Change to self-hosted server if needed
        NtfyTopicPrefix:  "classec",                // Change topic prefix
        Sections:         []string{"A1", "A2"...},  // Add/remove sections
        AlertIntervals:   []int{15},                // Minutes before class
        TimetableCSVPath: "timetable.csv",          // Path to timetable
    }
}
```

## Development

### Project Structure

```
classec/
├── main.go          # Entry point
├── scheduler.go     # Cron job & alert logic
├── timetable.go     # CSV parsing & time calculations
├── notifier.go      # ntfy.sh integration
├── config.go        # Configuration
├── timetable.csv    # Class schedule
└── go.mod           # Dependencies
```

### Testing

1. **Add a test class:**

   ```csv
   A1,Tue,TEST,19:50,20:50,TEST-ROOM,TEST-PROF
   ```

   (Set time to 15 minutes from now)

2. **Subscribe to test topic:**
   Subscribe to `classec-A1` on ntfy

3. **Run the app:**

   ```bash
   go run .
   ```

4. **Wait:** You should receive a notification within 1-2 minutes

## Troubleshooting

### No notifications received?

- ✅ Check that your section matches the CSV (case-sensitive)
- ✅ Verify the day format is 3 letters (Mon, Tue, Wed...)
- ✅ Ensure time format is HH:MM in 24-hour format
- ✅ Check server timezone matches your college timezone
- ✅ Confirm you're subscribed to the correct topic

### Duplicate notifications?

This shouldn't happen due to built-in deduplication, but if it does:

- Check that only one instance of the app is running
- Verify the time window in `timetable.go` (should be 13-16 minutes)

### App crashes or stops?

- Check logs for error messages
- Ensure CSV file is properly formatted
- Verify Go version is 1.20 or higher

## Technical Details

- **Language:** Go 1.20+
- **Dependencies:**
  - `github.com/robfig/cron/v3` - Cron scheduler
- **Notification Service:** ntfy.sh (free, open-source)
- **Memory Usage:** ~5-10 MB
- **CPU Usage:** Minimal (checks run once per minute)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [ntfy.sh](https://ntfy.sh) - For the amazing free notification service
- [robfig/cron](https://github.com/robfig/cron) - For the cron scheduler library

## Support

If you have any questions or issues, please open an issue on GitHub.

---

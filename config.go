package main

type Config struct {
	NtfyServerURL    string
	NtfyTopicPrefix  string
	Sections         []string
	AlertIntervals   []int
	TimetableCSVPath string
}

func GetConfig() Config {
	return Config{
		NtfyServerURL:    "https://ntfy.sh",
		NtfyTopicPrefix:  "classec",
		Sections:         []string{"A1", "A2", "B1", "B2", "C1", "C2", "D1", "D2", "E1", "E2", "F1", "F2", "G1", "G2", "H1", "H2", "I1", "I2", "J1", "J2", "K1", "K2", "L1", "L2", "M1", "M2", "N1", "N2", "O1", "O2", "P1", "P2", "Q1", "Q2", "R1", "R2", "S1", "S2", "T1", "T2", "U1", "U2", "V1", "V2", "W1", "W2", "X1", "X2", "Y1", "Y2", "Z1", "Z2", "ZA1", "ZA2"},
		AlertIntervals:   []int{15},
		TimetableCSVPath: "timetable.csv",
	}
}

func (c Config) GetTopicForSection(section string) string {
	return ("classec-" + section)
}

package domain

import "time"

type UserStats struct {
	UserID          UserID
	AttendanceCount int
	Punctuality     float64
	Streak          int
	LastUpdate      time.Time
}

func NewUserStats(userID UserID) *UserStats {
	return &UserStats{
		UserID:      userID,
		Punctuality: 0,
		Streak:      0,
		LastUpdate:  time.Now(),
	}
}

// RecordActivity actualiza la lógica de rachas y puntualidad
func (s *UserStats) RecordActivity(onTime bool) {
	s.AttendanceCount++
	if onTime {
		s.Streak++
	} else {
		s.Streak = 0
	}

	// Lógica de actualización de puntualidad (ejemplo simple)
	s.Punctuality = (s.Punctuality*float64(s.AttendanceCount-1) + map[bool]float64{true: 100, false: 0}[onTime]) / float64(s.AttendanceCount)
	s.LastUpdate = time.Now()
}

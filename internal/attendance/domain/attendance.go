package domain

import (
	"time"
)

type Attendance struct {
	ID           AttendanceID
	UserID       UserID
	Timestamp    time.Time
	SignatureURL SignatureURL
	IP           IPAddress
	Status       AttendanceStatus
	Audit        AuditLog
}

func NewAttendance(id string, userID UserID, sigURLStr, ipStr string) (*Attendance, error) {
	sigURL, err := NewSignatureURL(sigURLStr)
	if err != nil {
		return nil, err
	}

	ip, err := NewIPAddress(ipStr)
	if err != nil {
		return nil, err
	}

	return &Attendance{
		ID:           AttendanceID(id),
		UserID:       userID,
		Timestamp:    time.Now(),
		SignatureURL: sigURL,
		IP:           ip,
		Status:       StatusPending,
		Audit:        NewAuditLog(userID),
	}, nil
}

// TransitionToStatus implementa una pequeña máquina de estados
func (a *Attendance) TransitionToStatus(newStatus AttendanceStatus) error {
	if a.Status == StatusVerified || a.Status == StatusRejected {
		return ErrAlreadyProcessed
	}
	a.Status = newStatus
	a.Audit.UpdatedAt = time.Now()
	return nil
}

// IsFromToday verifica si el registro pertenece al día actual (útil para la regla de unicidad)
func (a *Attendance) IsFromToday() bool {
	now := time.Now()
	return a.Timestamp.Year() == now.Year() &&
		a.Timestamp.Month() == now.Month() &&
		a.Timestamp.Day() == now.Day()
}

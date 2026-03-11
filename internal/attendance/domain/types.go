package domain

import (
	"errors"
	"net"
	"net/mail"
	"strings"
	"time"
)

var (
	ErrInvalidEmail     = errors.New("formato de email inválido")
	ErrInvalidIP        = errors.New("dirección IP inválida")
	ErrInvalidSignature = errors.New("la URL de la firma debe ser de un almacenamiento seguro")
	ErrEmptyID          = errors.New("el identificador no puede estar vacío")
)

type (
	UserID       string
	AttendanceID string
	ShiftID      string
)

func NewUserID(v string) (UserID, error) {
	if strings.TrimSpace(v) == "" {
		return "", ErrEmptyID
	}
	return UserID(v), nil
}

// --- Datos de Identidad ---
type Email string

func NewEmail(v string) (Email, error) {
	v = strings.TrimSpace(strings.ToLower(v))
	_, err := mail.ParseAddress(v)
	if err != nil {
		return "", ErrInvalidEmail
	}
	return Email(v), nil
}

type PasswordHash string

// --- Datos de Logística y Firma ---
type SignatureURL string

func NewSignatureURL(v string) (SignatureURL, error) {
	// Validación: Solo permitimos nuestro bucket de GCS por seguridad
	if !strings.HasPrefix(v, "https://storage.googleapis.com/") {
		return "", ErrInvalidSignature
	}
	return SignatureURL(v), nil
}

type IPAddress string

func NewIPAddress(v string) (IPAddress, error) {
	if net.ParseIP(v) == nil {
		return "", ErrInvalidIP
	}
	return IPAddress(v), nil
}

// --- Roles y Estados ---
type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleStaff    UserRole = "staff"
	RoleEmployee UserRole = "employee"
)

func (r UserRole) IsAdmin() bool {
	return r == RoleAdmin
}

type AttendanceStatus string

const (
	StatusPending  AttendanceStatus = "pending"
	StatusVerified AttendanceStatus = "verified"
	StatusRejected AttendanceStatus = "rejected"
)

// --- Estructuras de Soporte (Value Objects Complejos) ---

// AuditLog encapsula la trazabilidad de cualquier entidad
type AuditLog struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy UserID
}

func NewAuditLog(creator UserID) AuditLog {
	now := time.Now()
	return AuditLog{
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: creator,
	}
}

type GeoLocation struct {
	Latitude  float64
	Longitude float64
}

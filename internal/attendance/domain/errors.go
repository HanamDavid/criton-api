package domain

import "errors"

var (
	ErrAccessDenied     = errors.New("no tienes permisos para realizar esta operación")
	ErrAlreadyProcessed = errors.New("esta asistencia ya fue verificada o rechazada")
	ErrShiftMismatch    = errors.New("el registro no coincide con el horario asignado")
	ErrAlreadyCheckedIn = errors.New("ya existe un registro de asistencia para el día de hoy")
)

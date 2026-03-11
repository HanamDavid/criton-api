package domain

import (
	"time"
)

type User struct {
	ID       UserID
	Name     string
	Email    Email
	Password PasswordHash
	Role     UserRole
	IsActive bool
	Audit    AuditLog
}

// NewUser es el constructor principal con validación de tipos
func NewUser(id, name, emailStr, pass string, role UserRole) (*User, error) {
	uID, err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	uEmail, err := NewEmail(emailStr)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uID,
		Name:     name,
		Email:    uEmail,
		Password: PasswordHash(pass),
		Role:     role,
		IsActive: true,
		Audit:    NewAuditLog(uID),
	}, nil
}

// Authenticate permite validar el acceso sin acoplar el dominio a bcrypt
func (u *User) Authenticate(plainPassword string, checkFn func(hash, plain string) bool) bool {
	return u.IsActive && checkFn(string(u.Password), plainPassword)
}

func (u *User) Deactivate() {
	u.IsActive = false
	u.Audit.UpdatedAt = time.Now()
}

func (u *User) CanManageSystem() bool {
	return u.IsActive && u.Role.IsAdmin()
}

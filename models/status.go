package models

// Status represents the status of an entity
type Status int

const (
	StatusActive Status = iota
	StatusInactive
	StatusPending
	StatusDeleted
)

// String returns the string representation
func (s Status) String() string {
	switch s {
	case StatusActive:
		return "ACTIVE"
	case StatusInactive:
		return "INACTIVE"
	case StatusPending:
		return "PENDING"
	case StatusDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}

// IsValid checks if the status is valid
func (s Status) IsValid() bool {
	return s >= StatusActive && s <= StatusDeleted
}

package repository

import "errors"

var (
	// ErrNotFound is returned when a requested entity does not exist.
	ErrNotFound = errors.New("not found")

	// ErrAlreadyClosed is returned when attempting to close an already-closed hanchan.
	ErrAlreadyClosed = errors.New("hanchan is already closed")

	// ErrDuplicateEntry is returned on unique constraint violations.
	ErrDuplicateEntry = errors.New("duplicate entry")
)

package service

import "fmt"

type NotFoundError struct {
	message       string
	originalError error
}

func (n *NotFoundError) Error() string {
	return formatError(n.message, n.originalError)
}

func (n *NotFoundError) Unwrap() error {
	return n.originalError
}

func (n *NotFoundError) Wrap(err error) error {
	n.originalError = err
	return n
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message: message}
}

type InternalServerError struct {
	message       string
	originalError error
}

func (i *InternalServerError) Error() string {
	return formatError(i.message, i.originalError)
}

func formatError(message string, originalError error) string {
	if originalError != nil {
		return fmt.Sprintf("%s: %v", message, originalError)
	}

	return message
}

func (i *InternalServerError) Unwrap() error {
	return i.originalError
}

func (i *InternalServerError) Wrap(err error) error {
	i.originalError = err
	return i
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message: message}
}

type BadRequestError struct {
	message       string
	originalError error
}

func (b *BadRequestError) Error() string {
	return formatError(b.message, b.originalError)
}

func (b *BadRequestError) Unwrap() error {
	return b.originalError
}

func (b *BadRequestError) Wrap(err error) error {
	b.originalError = err
	return b
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{message: message}
}

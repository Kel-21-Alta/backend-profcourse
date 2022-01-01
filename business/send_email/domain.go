package send_email

import (
	"context"
)

type Repository interface {
	SendEmail(ctx context.Context, to, subject, message string) error
}

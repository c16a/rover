package drivers

import "context"

// Driver interface specifies the properties of a driver.
type Driver interface {
	GetName(ctx context.Context) string
	GetVersion(ctx context.Context) string
	IsHealthy(ctx context.Context) error
}

package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func ToDuration(ctx context.Context, d *transform.TransformData) (any, error) {
	if value, ok := d.Value.(int); ok {
		return time.Duration(value * 1_000_000_000).Round(time.Second).String(), nil
	}
	return nil, fmt.Errorf("invalid type for input value: %T", d.Value)
}

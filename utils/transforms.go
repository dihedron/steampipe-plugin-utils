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

func UnixTimeToDate(ctx context.Context, d *transform.TransformData) (any, error) {
	var err error
	switch t := d.Value.(type) {
	case int64:
		if t == 0 {
			return nil, nil
		}
		return time.Unix(t, 0).Format(time.RFC3339), nil
	default:
		err = fmt.Errorf("invalid type: %T", d.Value)
	}
	return nil, err
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlptext // import "go.opentelemetry.io/collector/exporter/jsonloggingexporter/internal/otlptext"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// NewJSONLogsMarshaler returns a plog.Marshaler to encode to OTLP JSON text bytes.
func NewJSONLogsMarshaler() plog.Marshaler {
	return &plog.JSONMarshaler{}
}

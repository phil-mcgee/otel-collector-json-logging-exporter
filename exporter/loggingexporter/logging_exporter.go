// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loggingexporter // import "go.opentelemetry.io/collector/exporter/loggingexporter"

import (
	"context"
	"errors"
	"os"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/exporter/loggingexporter/internal/otlptext"
	"go.opentelemetry.io/collector/pdata/plog"
)

type loggingExporter struct {
	verbosity     configtelemetry.Level
	logger        *zap.SugaredLogger
	logsMarshaler plog.Marshaler
}

func (s *loggingExporter) pushLogs(_ context.Context, ld plog.Logs) error {
	s.logger.Info("received log counts",
		"resource log count", ld.ResourceLogs().Len(),
		"log record count", ld.LogRecordCount(),
	)
	if s.verbosity != configtelemetry.LevelDetailed {
		return nil
	}

	s.logger.Info("received logs",
		"logs", ld,
	)
	return nil
}

func newLoggingExporter(logger *zap.SugaredLogger, verbosity configtelemetry.Level) *loggingExporter {
	return &loggingExporter{
		verbosity:     verbosity,
		logger:        logger,
		logsMarshaler: otlptext.NewLogsMarshaler(),
	}
}

func loggerSync(logger *zap.SugaredLogger) func(context.Context) error {
	return func(context.Context) error {
		// Currently Sync() return a different error depending on the OS.
		// Since these are not actionable ignore them.
		err := logger.Sync()
		osErr := &os.PathError{}
		if errors.As(err, &osErr) {
			wrappedErr := osErr.Unwrap()
			if knownSyncError(wrappedErr) {
				err = nil
			}
		}
		return err
	}
}

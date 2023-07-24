// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package loggingexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/internal/testdata"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestLoggingLogsExporterNoErrors(t *testing.T) {
	f := NewFactory()
	lle, err := f.CreateLogsExporter(context.Background(), exportertest.NewNopCreateSettings(), f.CreateDefaultConfig())
	require.NotNil(t, lle)
	assert.NoError(t, err)

	assert.NoError(t, lle.ConsumeLogs(context.Background(), plog.NewLogs()))
	assert.NoError(t, lle.ConsumeLogs(context.Background(), testdata.GenerateLogs(10)))

	assert.NoError(t, lle.Shutdown(context.Background()))
}

// func TestLoggingExporterErrors(t *testing.T) {
// 	var config = createDefaultConfig()
// 	var logger, err = DefaultLoggerConfig(config.(*Config)).Build()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	le := newLoggingExporter(logger.Sugar(), configtelemetry.LevelDetailed)
// 	require.NotNil(t, le)
//
// 	errWant := errors.New("my error")
// 	le.logsMarshaler = &errMarshaler{err: errWant}
// 	le.pushLogs(context.Background(), plog.NewLogs())
// 	assert.Equal(t, errWant, le.pushLogs(context.Background(), plog.NewLogs()))
// }

// type errMarshaler struct {
// 	err error
// }

// func (e errMarshaler) MarshalLogs(plog.Logs) ([]byte, error) {
// 	return nil, e.err
// }

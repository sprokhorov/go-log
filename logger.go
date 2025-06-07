package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Level string `koanf:"level"`
}

func (o *Options) SetDefaults() {
	if o.Level == "" {
		o.Level = "info"
	}
}

func (o *Options) Validate() error { return nil }

func NewZapLogger(w io.Writer, opts *Options) (*zap.Logger, error) {
	// Parse log level from string
	l, err := zap.ParseAtomicLevel(opts.Level)
	if err != nil {
		return &zap.Logger{}, err
	}

	// Configure encoder (JSON or console format)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create a core that writes to the custom writer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(w),
		l,
	)

	return zap.New(core), nil
}

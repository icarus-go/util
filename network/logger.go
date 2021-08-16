package network

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = _pool.Get
)

type MyEncoder struct {
	*zapcore.MapObjectEncoder
}

func (enc MyEncoder) Clone() zapcore.Encoder {
	return MyEncoder{
		MapObjectEncoder: zapcore.NewMapObjectEncoder(),
	}
}

func (enc MyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	myEnc := enc.Clone().(MyEncoder)
	buf := _pool.Get()

	buf.AppendString(entry.Message)
	buf.AppendString(" ")

	for _, field := range fields {
		field.AddTo(myEnc)
		value := myEnc.MapObjectEncoder.Fields[field.Key]
		buf.AppendString(field.Key)
		buf.AppendString("=")
		if value == "" {
			buf.AppendString(" ''")
		} else {
			buf.AppendString(fmt.Sprintf("%v ", value))
		}
	}

	buf.AppendByte('\n')

	if entry.Stack != "" {
		buf.AppendString(entry.Stack)
		buf.AppendByte('\n')
	}
	return buf, nil
}

func DefaultLogger() *zap.Logger {
	_ = zap.RegisterEncoder("my-encoder", func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return MyEncoder{
			MapObjectEncoder: zapcore.NewMapObjectEncoder(),
		}, nil
	})
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "my-encoder"
	logger, _ := cfg.Build()
	return logger
}

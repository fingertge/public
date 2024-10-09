// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 17:42:13
// * Proj: public
// * Pack: logs
// * File: zerolog.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package logs

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	_zero    zerolog.Logger
	_zeroPwd string
	_zeroDir string
)

func zeroCallerMarshler(pc uintptr, file string, line int) string {
	l := strings.TrimPrefix(file, _zeroPwd)
	return _zeroDir + l + ":" + strconv.Itoa(line)
}

func consoleCaller(pc uintptr, file string, line int) string {
	l := strings.TrimPrefix(file, _zeroPwd)
	return _zeroDir + l + "(" + strconv.Itoa(line) + "):"
}

func consoleDefaultPartsOrder() []string {
	return []string{
		zerolog.LevelFieldName,
		zerolog.TimestampFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
}

func init() {
	_zeroPwd, _ = os.Getwd()
	_zeroPwd = strings.ReplaceAll(_zeroPwd, "\\", "/")
	_zeroDir = filepath.Base(_zeroPwd)
	zerolog.LevelInfoValue = "infos"
	zerolog.LevelWarnValue = "warns"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func InitZeroLog(writer io.Writer, level zerolog.Level, hooks ...zerolog.Hook) {
	zerolog.CallerMarshalFunc = zeroCallerMarshler
	zerolog.SetGlobalLevel(level)
	_zero = zerolog.New(writer)

	for _, hook := range hooks {
		_zero.Hook(hook)
	}
}

func initConsoleLog(level zerolog.Level) {
	zerolog.CallerMarshalFunc = consoleCaller
	zerolog.SetGlobalLevel(level)
	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}
	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("<%s>", i)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("[%s]", i)
	}
	output.FormatCaller = func(i interface{}) string {
		return i.(string)
	}
	output.PartsOrder = consoleDefaultPartsOrder()
	_zero = zerolog.New(output)
}

func Trace() *zerolog.Event {
	return _zero.Trace().Timestamp().Caller(1).Stack()
}

func Debug() *zerolog.Event {
	return _zero.Debug().Timestamp().Caller(1).Stack()
}

func Infos() *zerolog.Event {
	return _zero.Info().Timestamp().Caller(1).Stack()
}

func Warns() *zerolog.Event {
	return _zero.Warn().Timestamp().Caller(1).Stack()
}

func Error() *zerolog.Event {
	return _zero.Error().Timestamp().Caller(1).Stack()
}

func Fatal() *zerolog.Event {
	return _zero.Fatal().Timestamp().Caller(1).Stack()
}

func Panic() *zerolog.Event {
	return _zero.Panic().Timestamp().Caller(1).Stack()
}

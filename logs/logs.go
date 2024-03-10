// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/10/25 23:06:57                                                                                         *
// * Proj: work                                                                                                        *
// * Pack: logs                                                                                                        *
// * File: logs.go                                                                                                     *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// * 	- some init work.
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * 	- InitLogs(parttern string, maxSize uint64, maxLevel Level): init file logger.                                 *
// * 	- InitConsoleLogs(maxLevel Level): init console logger.                                                        *
// * 	- SetLogLevel(level Level): set log level.                                                                     *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *
package logs

import (
	"github.com/rs/zerolog"
)

type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.use FatalLevel will exit the program.
	FatalLevel
	// PanicLevel defines panic log level.use PanicLevel will panic and exit the program.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
	// Values less than TraceLevel are handled as numbers.
)

// parttern group by ./directory/timeformat/filename like ./log/20060102/zer.log
// *********************************************************************************************************************
// * SUMMARY:                                                                                                          *
// * 	- init logger, this logger will write log to file.															   *
// * 	- inputs:                                                                                                      *
// *        - parttern string: log file save info,include directory etc... exp: ./log/20060102/zer.log                 *
// *        - maxSizxe uint64: log file max size, bigger than maxsize log file will split.                             *
// *        - maxLevel Level(int8): level less than max level log will save.                                           *
// * WARNING: none.                                                                                                    *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 22:59:07 ColeCai.                                                                          *
// *********************************************************************************************************************
func InitLogs(parttern string, maxSize uint64, maxLevel Level) {
	rot := NewRotator(WithPattern(parttern), WithMaxSize(maxSize))
	InitZeroLog(rot, zerolog.Level(maxLevel))
}

// *********************************************************************************************************************
// * SUMMARY:                                                                                                          *
// * 	- init logger, this logger will write to console.                                                              *
// * 	- inputs:                                                                                                      *
// *        - maxLevel Level(int8): level less than max level log will save.                                           *
// * WARNING:                                                                                                          *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 23:04:01 ColeCai.                                                                          *
// *********************************************************************************************************************
func InitConsoleLogs(maxLevel Level) {
	initConsoleLog(zerolog.Level(maxLevel))
}

// *********************************************************************************************************************
// * SUMMARY: none.                                                                                                    *
// * WARNING: none.                                                                                                    *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 23:06:11 ColeCai.                                                                          *
// *********************************************************************************************************************
func SetLogLevel(level Level) {
	zerolog.SetGlobalLevel(zerolog.Level(level))
}

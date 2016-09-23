package servicetest

import (
	"github.com/jabong/florest-core/src/common/logger"
)

func initTestLogger() {
	logger.Initialise("testdata/testloggerAsync.json")
}

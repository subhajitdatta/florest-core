package servicetest

import (
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/logger/impls"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

func loggerTest() {
	gk.Describe("test async logger with string format", func() {
		aerr := logger.Initialise("testdata/testLoggerAsync.json")
		gk.Context("then log manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(aerr).To(gm.BeNil())
			})
		})
		gk.Context("then debug,info,trace,profile,warning,error logging", func() {
			logger.Debug("florest-debug")
			logger.Info("florest-info")
			logger.Trace("florest-trace")
			logger.Profile("florest-profile")
			logger.Warning("florest-warning")
			logger.Error("florest-error")
			gk.It("should not return error", func() {
			})
		})
		gk.Context("then flush,destroy", func() {
			handle, _ := logger.GetLoggerHandle(logger.GetDefaultLogTypeKey())
			asyncLogr, _ := handle.(*impls.AsynchLogger)
			asyncLogr.Flush()
			asyncLogr.Destroy()
			gk.It("should not return error", func() {
			})
		})

	})
	gk.Describe("test sync logger with json format", func() {
		serr := logger.Initialise("testdata/testLoggerSync.json")
		gk.Context("then log manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(serr).To(gm.BeNil())
			})
		})
		gk.Context("then debug,info,trace,profile,warning,error logging", func() {
			logger.Debug("florest-debug")
			logger.Info("florest-info")
			logger.Trace("florest-trace")
			logger.Profile("florest-profile")
			logger.Warning("florest-warning")
			logger.Error("florest-error")
			gk.It("should not return error", func() {
			})
		})
	})
}

package servicetest

import (
	"github.com/jabong/florest-core/src/components/sqldb"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

func sqldbTest() {
	gk.Describe("test invalid sqldb driver", func() {
		dbObj, err := sqldb.Get("invalid")
		gk.Context("then sql manager", func() {
			gk.It("should return nil impl", func() {
				gm.Expect(dbObj).To(gm.BeNil())
			})
			gk.It("should return error", func() {
				gm.Expect(err.ErrCode).To(gm.Equal(sqldb.ErrKeyNotPresent))
			})
		})
	})

	gk.Describe("test mysql impl", func() {
		testConf, err := getTestAPPConfig()
		gk.Context("then test app config", func() {
			gk.It("should not return error", func() {
				gm.Expect(err).To(gm.BeNil())
			})
		})
		conf := testConf.MySQL
		serr := sqldb.Set("mysdb", conf, new(sqldb.MysqlDriver))
		gk.Context("then sql manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(serr).To(gm.BeNil())
			})
		})
		dbObj, terr := sqldb.Get("mysdb")
		gk.Context("then sql manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(terr).To(gm.BeNil())
			})
		})

		gk.Context("then db impl", func() {
			perr := dbObj.Ping()
			gk.It("should not give error on ping", func() {
				gm.Expect(perr).To(gm.BeNil())
			})
		})

		gk.Context("then db impl", func() {
			_, verr := dbObj.Execute("drop table if exists florest_mysql")
			gk.It("should not give error on valid execute", func() {
				gm.Expect(verr).To(gm.BeNil())
			})
		})

		gk.Context("then db impl", func() {
			_, iverr := dbObj.Execute("invalid execute")
			gk.It("should give error on invalid execute", func() {
				gm.Expect(iverr.ErrCode).To(gm.Equal(sqldb.ErrExecuteFailure))
			})
		})

		gk.Context("then db impl", func() {
			_, qerr := dbObj.Query("show tables")
			gk.It("should not give error on valid query", func() {
				gm.Expect(qerr).To(gm.BeNil())
			})
		})

		gk.Context("then db impl", func() {
			_, qierr := dbObj.Query("invalid query")
			gk.It("should give error on invalid query", func() {
				gm.Expect(qierr.ErrCode).To(gm.Equal(sqldb.ErrQueryFailure))
			})
		})

		gk.Context("then db impl", func() {
			_, terr := dbObj.GetTxnObj()
			gk.It("should not give error on starting txn", func() {
				gm.Expect(terr).To(gm.BeNil())
			})
		})
		gk.Context("then db impl", func() {
			cerr := dbObj.Close()
			gk.It("should not give error on closing", func() {
				gm.Expect(cerr).To(gm.BeNil())
			})
		})
	})

}

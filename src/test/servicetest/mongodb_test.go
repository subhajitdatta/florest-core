package servicetest

import (
	"encoding/json"
	"github.com/jabong/florest-core/src/components/mongodb"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

// an example for mongo document
type employeeInfo struct {
	ID   string
	Type string
}

func mongodbTest() {
	gk.Describe("test invalid mongodb details", func() {
		_, err := mongodb.Get("invalid")
		gk.Context("then db manager", func() {
			gk.It("should return error", func() {
				gm.Expect(err.ErrCode).To(gm.Equal(mongodb.ErrKeyNotPresent))
			})
		})
	})

	gk.Describe("test mongodb impl", func() {
		collection := "employee"
		conf := new(mongodb.MDBConfig)
		conf.URL = "mongodb://localhost:27017"
		conf.DbName = "florest"
		serr := mongodb.Set("mymongo", conf, new(mongodb.MongoDriver))
		gk.Context("then db manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(serr).To(gm.BeNil())
			})
		})
		db, terr := mongodb.Get("mymongo")
		gk.Context("then db manager", func() {
			gk.It("should not return error", func() {
				gm.Expect(terr).To(gm.BeNil())
			})
		})
		gk.Context("then mongodb impl", func() {
			ierr := db.Insert(collection, &employeeInfo{ID: "123", Type: "Manager"})
			gk.It("should not return error on insert", func() {
				gm.Expect(ierr).To(gm.BeNil())
			})
		})
		query := make(map[string]interface{}, 1)
		query["id"] = "123"
		gk.Context("then mongodb impl", func() {
			uerr := db.Update(collection, query, &employeeInfo{ID: "123", Type: "Director"})
			gk.It("should not return error on update", func() {
				gm.Expect(uerr).To(gm.BeNil())
			})
		})
		gk.Context("then mongodb impl", func() {
			ret, ferr := db.FindOne(collection, query)
			gk.It("should not return error on findone", func() {
				gm.Expect(ferr).To(gm.BeNil())
			})
			gk.It("should not return correct result", func() {
				emp1 := new(employeeInfo)
				byt, _ := json.Marshal(ret)
				json.Unmarshal(byt, emp1)
				gm.Expect(emp1.ID).To(gm.Equal("123"))
			})

		})
		gk.Context("then mongodb impl", func() {
			_, faerr := db.FindAll(collection, query)
			gk.It("should not return error on findall", func() {
				gm.Expect(faerr).To(gm.BeNil())
			})
		})
		gk.Context("then mongodb impl", func() {
			raerr := db.Remove(collection, query)
			gk.It("should not return error on remove", func() {
				gm.Expect(raerr).To(gm.BeNil())
			})
		})

	})
}

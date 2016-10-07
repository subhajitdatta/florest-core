package servicetest

import (
	"github.com/jabong/florest-core/src/components/cache"
	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
)

func cacheTest() {
	gk.Describe("test invalid cache platform", func() {
		obj, err := cache.Get("invalid")
		gk.Context("then cache manager", func() {
			gk.It("should return nil impl", func() {
				gm.Expect(obj).To(gm.BeNil())
			})
			gk.It("should return error", func() {
				gm.Expect(err).Should(gm.HaveOccurred())
			})
		})
	})
	gk.Describe("test redis cache", func() {
		conf := new(cache.Config)
		conf.Platform = "redis"
		conf.Cluster = false
		conf.ConnStr = "localhost:6379"
		conf.BucketHashes = []string{"florestHash"}
		serr := cache.Set("myredis", conf, new(cache.RedisClientAdapter))
		gk.Context("then cache manager", func() {
			gk.It("should return no error", func() {
				gm.Expect(serr).To(gm.BeNil())
			})
		})
		obj, terr := cache.Get("myredis")
		gk.Context("then cache manager", func() {
			gk.It("should return no error", func() {
				gm.Expect(terr).To(gm.BeNil())
			})
		})
		item := cache.Item{Key: "florest_key", Value: "florest_value", Error: ""}
		itemExp := cache.Item{Key: "florest_expiry_key", Value: "florest_expiry_value", Error: ""}
		gk.Context("then redis impl", func() {
			serr := obj.Set(item, false, false)
			gk.It("should not return error on set", func() {
				gm.Expect(serr).To(gm.BeNil())
			})
		})
		gk.Context("then redis impl", func() {
			sxerr := obj.Set(itemExp, false, false)
			gk.It("should not return error on set with expiry", func() {
				gm.Expect(sxerr).To(gm.BeNil())
			})
		})
		gk.Context("then redis impl", func() {
			ret, gerr := obj.Get(item.Key, false, false)
			gk.It("should not return error on get", func() {
				gm.Expect(gerr).To(gm.BeNil())
			})
			gk.It("should return correct item on get", func() {
				gm.Expect(ret.Value).To(gm.Equal(item.Value))
			})
		})
		gk.Context("then redis impl", func() {
			derr := obj.Delete(item.Key)
			gk.It("should not return error on delete", func() {
				gm.Expect(derr).To(gm.BeNil())
			})
		})
		gk.Context("then redis impl", func() {
			retb, gberr := obj.GetBatch([]string{itemExp.Key}, false, false)
			gk.It("should not return error on get batch", func() {
				gm.Expect(gberr).To(gm.BeNil())
			})
			gk.It("should return correct item on get", func() {
				gm.Expect(retb[itemExp.Key].Value).To(gm.Equal(itemExp.Value))
			})
		})
		gk.Context("then redis impl", func() {
			dberr := obj.DeleteBatch([]string{itemExp.Key})
			gk.It("should not return error on delete batch", func() {
				gm.Expect(dberr).To(gm.BeNil())
			})
		})

	})
}

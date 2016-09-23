package cache

import ()

// Config contains the details about the cache platform to be used
type Config struct {
	Platform     string
	ConnStr      string
	Password     string
	BucketHashes []string
	Cluster      bool
}

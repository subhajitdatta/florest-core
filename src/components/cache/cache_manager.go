package cache

// Get() - Creates, initializes and returns the appropriate instance based on the Platform in the config
func Get(conf Config) (ret CInterface, err error) {
	switch conf.Platform {
	case Redis:
		ret = new(redisClientAdapter)
		if err = ret.Init(&conf); err != nil {
			return nil, err
		}
		return ret, nil
	default:
		return nil, getErrObj(ErrNoPlatform, conf.Platform+" is not supported")
	}
}

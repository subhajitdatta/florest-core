Contains the external libs.

Developer should manually

1. `go get <their_external_dependency>`
2. copy the folder inside _libs/src/
3. Till the time we find a better dependency management, the external libs should be checked in florest-core repository.
4. And delete the .git/ and .gitignore

The list of external dependencies:
```
github.com/go-sql-driver/mysql		version 1.3
github.com/imdario/mergo			version 0.2.0
github.com/onsi/ginkgo				version 1.2.0
github.com/onsi/gomega				version 1.0
github.com/ooyala/go-dogstatsd		NA
github.com/twinj/uuid				version prior to 0.10.0

gopkg.in/mgo.v2						version v2
gopkg.in/redis.v3					version v3
gopkg.in/bsm/ratelimit.v1           version v1

```

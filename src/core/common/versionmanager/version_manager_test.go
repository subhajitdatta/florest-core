package versionmanager

import (
	"reflect"
	"testing"
)

/*
Test Versionable interface implementation
*/
type testVersionableImpl struct {
}

func (o testVersionableImpl) GetInstance() interface{} {
	return o
}

/*
Test version manager
*/
func TestVersionManager(t *testing.T) {
	resource := "TEST_RESOURCE"
	version := "TEST_VERSION"
	action := "TEST_ACTION"
	bucketID := "TEST_BUCKET_ID"

	testParam := NewParam()
	testParam.Update("", *new(testVersionableImpl), nil)

	vmap := VersionMap{
		BasicVersion{
			Resource: resource,
			Version:  version,
			Action:   action,
			BucketID: bucketID,
		}: testParam,
	}

	version1 := Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
		Path:     "buckets/keys",
	}

	version2 := Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
		Path:     "buckets/{bucketId}/keys/{keyId}",
	}

	version3 := Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
		Path:     "buckets/{bucketId}/{keyId}/keys",
	}

	version4 := Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
		Path:     "{bucketId}/{keyId}",
	}

	version5 := Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
		Path:     "{bucketId}/buckets/{keyId}/keys",
	}

	addTestVersions(version1, vmap, *new(testVersionableImpl))
	addTestVersions(version2, vmap, *new(testVersionableImpl))
	addTestVersions(version3, vmap, *new(testVersionableImpl))
	addTestVersions(version4, vmap, *new(testVersionableImpl))
	addTestVersions(version5, vmap, *new(testVersionableImpl))

	Initialize(vmap)

	// Valid URLs
	getVersionable(resource, version, action, bucketID, "", map[string]string{}, t)
	getVersionable(resource, version, action, bucketID, "buckets/keys", map[string]string{}, t)
	targetPmts := map[string]string{"bucketId": "1", "keyId": "2"}
	getVersionable(resource, version, action, bucketID, "buckets/1/keys/2", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "100", "keyId": "200"}
	getVersionable(resource, version, action, bucketID, "buckets/100/keys/200", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "bucketId", "keyId": "keyId"}
	getVersionable(resource, version, action, bucketID, "buckets/bucketId/keys/keyId", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "1", "keyId": "2"}
	getVersionable(resource, version, action, bucketID, "buckets/1/2/keys", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "1000", "keyId": "200"}
	getVersionable(resource, version, action, bucketID, "buckets/1000/200/keys", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "bucket", "keyId": "key"}
	getVersionable(resource, version, action, bucketID, "bucket/key", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "1", "keyId": "2"}
	getVersionable(resource, version, action, bucketID, "1/buckets/2/keys", targetPmts, t)
	targetPmts = map[string]string{"bucketId": "buckets", "keyId": "keys"}
	getVersionable(resource, version, action, bucketID, "buckets/buckets/keys/keys", targetPmts, t)

	// Invalid URLs
	getVersionableExpectingErrors(resource, version, action, bucketID, "buckets/key", t)
	getVersionableExpectingErrors(resource, version, action, bucketID, "somepath", t)
	getVersionableExpectingErrors(resource, version, action, bucketID, "buckets/1/2", t)
	getVersionableExpectingErrors(resource, version, action, bucketID, "1/buckets/2", t)
	getVersionableExpectingErrors(resource, version, action, bucketID, "bucket/1/keys/2", t)
}

func getVersionable(resource string, version string, action string, bucketID string, path string, targetPmts map[string]string, t *testing.T) {
	versionableInstance, _, pmts, verr := Get(resource, version, action, bucketID, path)
	if verr != nil {
		t.Error("Failed to get versionable from version manager for this path - " + path)
	}

	_, ok := versionableInstance.(testVersionableImpl)

	if !ok {
		t.Error("Returned versionable instance mismatch for this path - " + path)
	}

	eq := reflect.DeepEqual(pmts, targetPmts)
	if !eq {
		t.Error("Returned parameters mismatch for this path - " + path)
	}
}

func getVersionableExpectingErrors(resource string, version string, action string, bucketID string, path string, t *testing.T) {
	_, _, _, verr := Get(resource, version, action, bucketID, path)
	if verr == nil {
		t.Error("Expected error, but got a valid versionable for this path - " + path)
	}
}

func addTestVersions(version Version, vmap VersionMap, testVersionableImpl Versionable) {
	param := vmap[version.GetBasicVersion()]
	if param == nil {
		param = NewParam()
		vmap[version.GetBasicVersion()] = param
	}
	param.Update(version.Path, testVersionableImpl, nil)
}

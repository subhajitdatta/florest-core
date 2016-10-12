package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
)

func getAppRequest(r *http.Request) (*utilhttp.Request, error) {
	req, rerr := utilhttp.GetRequest(r)
	if rerr != nil {
		return nil, rerr
	}

	return &req, nil
}

func getBucketsMap(bucketsList string) map[string]string {
	buckets := strings.Split(bucketsList, constants.FieldSeperator)
	bucketMap := make(map[string]string, len(buckets))
	for _, v := range buckets {
		bKV := strings.Split(v, constants.KeyValueSeperator)
		if len(bKV) < 2 { //invalid bucket
			continue
		}
		bucketMap[bKV[0]] = bKV[1]
	}
	return bucketMap
}

//Get the Service WorkFlow Data
func GetData(r *http.Request) (*orchestrator.WorkFlowData, error) {
	serviceInputOutput := new(orchestrator.WorkFlowIOInMemoryImpl)
	appReq, rerr := getAppRequest(r)
	if rerr != nil {
		return nil, rerr
	}

	serviceInputOutput.Set(constants.URI, appReq.URI)
	serviceInputOutput.Set(constants.HTTPVerb, appReq.HTTPVerb)
	serviceInputOutput.Set(constants.Request, appReq)

	logger.Info(fmt.Sprintf("Service Input Output %v", serviceInputOutput))

	serviceEcContext := new(orchestrator.WorkFlowECInMemoryImpl)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.UserID], appReq.Headers.UserID)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.SessionID], appReq.Headers.SessionID)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.TokenID], appReq.Headers.AuthToken)
	serviceEcContext.Set(constants.UserAgent, appReq.Headers.UserAgent)
	serviceEcContext.Set(constants.HTTPReferrer, appReq.Headers.Referrer)
	serviceEcContext.Set(utilhttp.CustomHeaderMap[utilhttp.RequestID], appReq.Headers.RequestID)
	serviceEcContext.SetBuckets(getBucketsMap(appReq.Headers.BucketsList))
	serviceEcContext.SetDebugFlag(appReq.Headers.Debug)

	serviceEcContext.Set(constants.RequestContext,
		utilhttp.RequestContext{
			AppName:       config.GlobalAppConfig.AppName,
			UserID:        appReq.Headers.UserID,
			SessionID:     appReq.Headers.SessionID,
			RequestID:     appReq.Headers.RequestID,
			TransactionID: appReq.Headers.TransactionID,
			URI:           appReq.URI,
			ClientAppID:   appReq.Headers.ClientAppID,
		})

	logger.Info(fmt.Sprintf("Service Execution Context %v", serviceInputOutput))

	serviceWorkFlowData := new(orchestrator.WorkFlowData)
	serviceWorkFlowData.Create(serviceInputOutput, serviceEcContext)

	return serviceWorkFlowData, nil
}

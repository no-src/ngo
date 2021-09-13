package _agent

import (
	"github.com/no-src/log"
	"github.com/no-src/ngo/cache"
	"github.com/no-src/ngo/contract"
	"github.com/no-src/ngo/encoding/json"
	"github.com/no-src/ngo/http"
	"time"
)

var (
	apiAgentLogger log.Logger
	cache          _cache.Cache
	httpClient     *_http.HttpClient
)

const (
	RequestCacheSeconds   = 60                      //请求缓存时间（秒）
	ResponseCacheSeconds  = 60 * 20                 //响应缓存时间（秒）
	ApiAgentRequestHeader = "ApiAgentRequestHeader" //ApiAgent请求头标识
	ApiAgentRequestValue  = "1"                     //ApiAgent请求头标识值
)

func init() {
	apiAgentLogger = log.NewConsoleLogger(log.DebugLevel)
	cache = &_cache.EmptyCache{}
	httpClient = _http.NewHttpClient()
}

// Post 发起Post请求，默认启用responseCached缓存，和RequestCacheSeconds秒的requestCached缓存，如果强制不使用任何缓存，则请调用PostNoCache方法
// url 请求完整地址
// request 请求参数
// response 响应结果
func Post(url string, request *_contract.ApiRequest) (response *_contract.ApiResponse, err error) {
	response, err = post(url, request, true, true)
	if err != nil {
		apiAgentLogger.Error(err, "【Post】【API调用】【异常】 URL=%s  Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	} else {
		apiAgentLogger.Debug("【Post】【API调用】【正常】URL=%s Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	}
	return response, err
}

// PostWithResponseCache 发起Post请求，启用responseCached缓存
// url 请求完整地址
// request 请求参数
// response 响应结果
func PostWithResponseCache(url string, request *_contract.ApiRequest) (response *_contract.ApiResponse, err error) {
	response, err = post(url, request, false, true)
	if err != nil {
		apiAgentLogger.Error(err, "【PostWithResponseCache】【API调用】【异常】 URL=%s  Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	} else {
		apiAgentLogger.Debug("【PostWithResponseCache】【API调用】【正常】URL=%s Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	}
	return response, err
}

// PostNoCache 发起Post请求，强制不使用任何缓存
// url 请求完整地址
// request 请求参数
// response 响应结果
func PostNoCache(url string, request *_contract.ApiRequest) (response *_contract.ApiResponse, err error) {
	response, err = post(url, request, false, false)
	if err != nil {
		apiAgentLogger.Error(err, "【PostNoCache】【API调用】【异常】 URL=%s  Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	} else {
		apiAgentLogger.Debug("【PostNoCache】【API调用】【正常】URL=%s Request=%s  Response=%s", url, _json.SerializeObject(request), _json.SerializeObject(response))
	}
	return response, err
}

func getCache(key string) (value interface{}, exists bool) {
	value, err := cache.Get(key)
	if err != nil || value == nil {
		return nil, false
	} else {
		return value, true
	}
}

func removeCache(key string) {
	cache.Delete(key)
}

func setCache(key string, value _contract.ApiResponse) error {
	cacheData := _contract.ApiResponseCache{}
	cacheData.CreateTime = time.Now()
	cacheData.CacheExpireTime = time.Now().Add(time.Second * RequestCacheSeconds)
	cacheData.StoreExpireTime = time.Now().Add(time.Second * ResponseCacheSeconds)
	cacheData.Data = value
	//Cache物理存储时间比StoreExpireTime多5分钟，防止数据接口返回时找不到本地的Cache数据
	return cache.Set(key, cacheData, ResponseCacheSeconds+60*5)
}

//post 发起post请求
//url 请求地址
//request 请求参数
//requestCached 是否缓存请求 如果为true，会先检查本地是否有缓存 有则直接返回缓存数据
//responseCached 是否缓存响应结果 如果为true，如果api服务器检测到服务端的响应结果与api调用客户端自己缓存的数据一致，则不会在响应中包含具体的数据
func post(url string, request *_contract.ApiRequest, requestCached bool, responseCached bool) (response *_contract.ApiResponse, err error) {

	//发起请求前检查缓存是否存在，如果存在并且没有过期则不请求接口
	cacheData := _contract.ApiResponseCache{}
	requestKey := ""
	if requestCached || responseCached {
		requestKey = _contract.GetRequestHash(url, request.RequestData)
		obj, exist := getCache(requestKey)
		//类型转换是否成功 如果类型转换失败 则请求接口 并清除本地缓存
		parseSuccess := false
		if exist {
			cacheData, parseSuccess = obj.(_contract.ApiResponseCache)
		}
		if parseSuccess == false || exist == false || cacheData.StoreExpireTime.Before(time.Now()) {
			//如果类型转换失败，或者缓存不存在，或者缓存存储时间已经过期，则清除缓存，重新请求API
			removeCache(requestKey)
		} else if requestCached && cacheData.CacheExpireTime.After(time.Now()) {
			//如果缓存过期时间还没到，并且启用了requestCached，则直接返回数据，不请求API
			response = &cacheData.Data
			return response, nil
		} else {
			//缓存过期，但缓存没被删除，则重新请求API服务器，并告知服务器本地数据的Hash值
			//设置本地保存的响应数据Hash值,如果开启了客户端缓存
			if responseCached {
				request.Hash = cacheData.Data.Hash
			}
		}

		if responseCached {
			//开启客户端缓存
			request.AllowClientCache = true
		}
	}

	postJson := _json.SerializeObject(request)
	apiAgentLogger.Info("【HTTP POST】URL=【%s】DATA=【%s】", url, postJson)
	header := make(map[string]string)
	header[ApiAgentRequestHeader] = ApiAgentRequestValue
	body, errReturn := httpClient.Post(url, request, "application/json", header, nil)
	if errReturn != nil {
		apiAgentLogger.Error(errReturn, "【API接口请求异常】URL=【%s】DATA=【%s】", url, postJson)
		return nil, errReturn
	}
	err = _json.Unmarshal(body, &response)
	if err != nil {
		apiAgentLogger.Error(err, "【API接口响应数据反序列化失败】URL=【%s】DATA=【%s】 ResponseData=【%s】", url, postJson, body)
		return nil, err
	}

	//如果服务器端的数据与客户端一致，则从本地客户端缓存查询
	if responseCached && response.ClientCached {
		response.Message = cacheData.Data.Message
		//这里使用绝对过期，暂时不延长缓存过期时间
	} else {
		//如果启用了客户端缓存，并且响应数据正常则写入服务器缓存
		if (requestCached || responseCached) && response.Code == 0 {
			setCache(requestKey, *response)
		}
	}
	return response, nil
}

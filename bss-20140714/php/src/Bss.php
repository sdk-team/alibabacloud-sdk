<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Bss\V20140714;

use AlibabaCloud\Tea\Tea;
use AlibabaCloud\Tea\Model;
use AlibabaCloud\Tea\Request;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Tea\Exception\TeaUnableRetryError;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Rpc\Rpc;
use AlibabaCloud\Endpoint\Endpoint;
use AlibabaCloud\Tea\Utils\Utils\RuntimeOptions;

use AlibabaCloud\SDK\Bss\V20140714\Models\OpenCallbackRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\OpenCallbackResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\QueryForCssOrderRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\QueryForCssOrderResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\CreateOrderRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\CreateOrderResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\vnoPayCallBackNotifyRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\vnoPayCallBackNotifyResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\vnoBatchRefundOrderRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\vnoBatchRefundOrderResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\SubscriptionCreateOrderApiRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\SubscriptionCreateOrderApiResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\SetResourceBusinessStatusRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\SetResourceBusinessStatusResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCouponListRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCouponListResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCouponDetailRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCouponDetailResponse;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCashDetailRequest;
use AlibabaCloud\SDK\Bss\V20140714\Models\DescribeCashDetailResponse;

class Bss {
    public function __construct($config){
        parent::__construct($config);
        $this->_endpointRule = "";
        $this->checkConfig($config);
        $this->_endpoint = $this->getEndpoint($this->_productId, $this->_regionId, $this->_endpointRule, $this->_network, $this->_suffix, $this->_endpointMap, $this->_endpoint);
    }

    /**
     * @param OpenCallbackRequest $request
     * @param RuntimeOptions $runtime
     * @return OpenCallbackResponse
     * @throws \Exception
     */
    public function openCallbackEx(OpenCallbackRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return OpenCallbackResponse::fromMap($this->doRequest("OpenCallback", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param OpenCallbackRequest $request
     * @return OpenCallbackResponse
     * @throws \Exception
     */
    public function openCallback(OpenCallbackRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->openCallbackEx($request, $runtime);
    }

    /**
     * @param QueryForCssOrderRequest $request
     * @param RuntimeOptions $runtime
     * @return QueryForCssOrderResponse
     * @throws \Exception
     */
    public function queryForCssOrderEx(QueryForCssOrderRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return QueryForCssOrderResponse::fromMap($this->doRequest("QueryForCssOrder", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param QueryForCssOrderRequest $request
     * @return QueryForCssOrderResponse
     * @throws \Exception
     */
    public function queryForCssOrder(QueryForCssOrderRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->queryForCssOrderEx($request, $runtime);
    }

    /**
     * @param CreateOrderRequest $request
     * @param RuntimeOptions $runtime
     * @return CreateOrderResponse
     * @throws \Exception
     */
    public function createOrderEx(CreateOrderRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return CreateOrderResponse::fromMap($this->doRequest("CreateOrder", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param CreateOrderRequest $request
     * @return CreateOrderResponse
     * @throws \Exception
     */
    public function createOrder(CreateOrderRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->createOrderEx($request, $runtime);
    }

    /**
     * @param vnoPayCallBackNotifyRequest $request
     * @param RuntimeOptions $runtime
     * @return vnoPayCallBackNotifyResponse
     * @throws \Exception
     */
    public function vnoPayCallBackNotifyEx(vnoPayCallBackNotifyRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return vnoPayCallBackNotifyResponse::fromMap($this->doRequest("vnoPayCallBackNotify", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param vnoPayCallBackNotifyRequest $request
     * @return vnoPayCallBackNotifyResponse
     * @throws \Exception
     */
    public function vnoPayCallBackNotify(vnoPayCallBackNotifyRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->vnoPayCallBackNotifyEx($request, $runtime);
    }

    /**
     * @param vnoBatchRefundOrderRequest $request
     * @param RuntimeOptions $runtime
     * @return vnoBatchRefundOrderResponse
     * @throws \Exception
     */
    public function vnoBatchRefundOrderEx(vnoBatchRefundOrderRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return vnoBatchRefundOrderResponse::fromMap($this->doRequest("vnoBatchRefundOrder", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param vnoBatchRefundOrderRequest $request
     * @return vnoBatchRefundOrderResponse
     * @throws \Exception
     */
    public function vnoBatchRefundOrder(vnoBatchRefundOrderRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->vnoBatchRefundOrderEx($request, $runtime);
    }

    /**
     * @param SubscriptionCreateOrderApiRequest $request
     * @param RuntimeOptions $runtime
     * @return SubscriptionCreateOrderApiResponse
     * @throws \Exception
     */
    public function subscriptionCreateOrderApiEx(SubscriptionCreateOrderApiRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return SubscriptionCreateOrderApiResponse::fromMap($this->doRequest("SubscriptionCreateOrderApi", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param SubscriptionCreateOrderApiRequest $request
     * @return SubscriptionCreateOrderApiResponse
     * @throws \Exception
     */
    public function subscriptionCreateOrderApi(SubscriptionCreateOrderApiRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->subscriptionCreateOrderApiEx($request, $runtime);
    }

    /**
     * @param SetResourceBusinessStatusRequest $request
     * @param RuntimeOptions $runtime
     * @return SetResourceBusinessStatusResponse
     * @throws \Exception
     */
    public function setResourceBusinessStatusEx(SetResourceBusinessStatusRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return SetResourceBusinessStatusResponse::fromMap($this->doRequest("SetResourceBusinessStatus", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param SetResourceBusinessStatusRequest $request
     * @return SetResourceBusinessStatusResponse
     * @throws \Exception
     */
    public function setResourceBusinessStatus(SetResourceBusinessStatusRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->setResourceBusinessStatusEx($request, $runtime);
    }

    /**
     * @param DescribeCouponListRequest $request
     * @param RuntimeOptions $runtime
     * @return DescribeCouponListResponse
     * @throws \Exception
     */
    public function describeCouponListEx(DescribeCouponListRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return DescribeCouponListResponse::fromMap($this->doRequest("DescribeCouponList", "HTTPS", "POST", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param DescribeCouponListRequest $request
     * @return DescribeCouponListResponse
     * @throws \Exception
     */
    public function describeCouponList(DescribeCouponListRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->describeCouponListEx($request, $runtime);
    }

    /**
     * @param DescribeCouponDetailRequest $request
     * @param RuntimeOptions $runtime
     * @return DescribeCouponDetailResponse
     * @throws \Exception
     */
    public function describeCouponDetailEx(DescribeCouponDetailRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return DescribeCouponDetailResponse::fromMap($this->doRequest("DescribeCouponDetail", "HTTPS", "GET", "2014-07-14", "AK", $request, null, $runtime))
;
    }

    /**
     * @param DescribeCouponDetailRequest $request
     * @return DescribeCouponDetailResponse
     * @throws \Exception
     */
    public function describeCouponDetail(DescribeCouponDetailRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->describeCouponDetailEx($request, $runtime);
    }

    /**
     * @param DescribeCashDetailRequest $request
     * @param RuntimeOptions $runtime
     * @return DescribeCashDetailResponse
     * @throws \Exception
     */
    public function describeCashDetailEx(DescribeCashDetailRequest $request, RuntimeOptions $runtime){
        Utils::validateModel($request);
        return DescribeCashDetailResponse::fromMap($this->doRequest("DescribeCashDetail", "HTTPS", "POST", "2014-07-14", "AK", null, $request, $runtime))
;
    }

    /**
     * @param DescribeCashDetailRequest $request
     * @return DescribeCashDetailResponse
     * @throws \Exception
     */
    public function describeCashDetail(DescribeCashDetailRequest $request){
        $runtime = new \AlibabaCloud\Tea\Utils\Utils\RuntimeOptions([]);
        return $this->describeCashDetailEx($request, $runtime);
    }

    /**
     * @param string $productId
     * @param string $regionId
     * @param string $endpointRule
     * @param string $network
     * @param string $suffix
     * @param array $endpointMap
     * @param string $endpoint
     * @return string
     * @throws \Exception
     */
    public function getEndpoint($productId, $regionId, $endpointRule, $network, $suffix, $endpointMap, $endpoint){
        if (!Utils::empty_($endpoint)) {
            return $endpoint;
        }
        if (!Utils::isUnset($endpointMap) && !Utils::empty_($endpointMap["regionId"])) {
            return $endpointMap["regionId"];
        }
        return Endpoint::getEndpointRules($productId, $regionId, $endpointRule, $network, $suffix);
    }
}

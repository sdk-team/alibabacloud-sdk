<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\ModifyInstanceRequest\parameter;

class ModifyInstanceRequest extends Model {
    protected $_name = [
        'productCode' => 'ProductCode',
        'ownerId' => 'OwnerId',
        'productType' => 'ProductType',
        'subscriptionType' => 'SubscriptionType',
        'modifyType' => 'ModifyType',
        'instanceId' => 'InstanceId',
        'parameter' => 'Parameter',
        'clientToken' => 'ClientToken',
    ];
    public function validate() {
        Model::validateRequired('productCode', $this->productCode, true);
        Model::validateRequired('subscriptionType', $this->subscriptionType, true);
        Model::validateRequired('modifyType', $this->modifyType, true);
    }
    public function toMap() {
        $res = [];
        $res['ProductCode'] = $this->productCode;
        $res['OwnerId'] = $this->ownerId;
        $res['ProductType'] = $this->productType;
        $res['SubscriptionType'] = $this->subscriptionType;
        $res['ModifyType'] = $this->modifyType;
        $res['InstanceId'] = $this->instanceId;
        $res['Parameter'] = [];
        if(null !== $this->parameter && is_array($this->parameter)){
            $n = 0;
            foreach($this->parameter as $item){
                $res['Parameter'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        $res['ClientToken'] = $this->clientToken;
        return $res;
    }
    /**
     * @param array $map
     * @return ModifyInstanceRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['ProductCode'])){
            $model->productCode = $map['ProductCode'];
        }
        if(isset($map['OwnerId'])){
            $model->ownerId = $map['OwnerId'];
        }
        if(isset($map['ProductType'])){
            $model->productType = $map['ProductType'];
        }
        if(isset($map['SubscriptionType'])){
            $model->subscriptionType = $map['SubscriptionType'];
        }
        if(isset($map['ModifyType'])){
            $model->modifyType = $map['ModifyType'];
        }
        if(isset($map['InstanceId'])){
            $model->instanceId = $map['InstanceId'];
        }
        if(isset($map['Parameter'])){
            if(!empty($map['Parameter'])){
                $model->parameter = [];
                $n = 0;
                foreach($map['Parameter'] as $item) {
                    $model->parameter[$n++] = null !== $item ? parameter::fromMap($item) : $item;
                }
            }
        }
        if(isset($map['ClientToken'])){
            $model->clientToken = $map['ClientToken'];
        }
        return $model;
    }
    /**
     * @description productCode
     * @var string
     */
    public $productCode;

    /**
     * @description ownerId
     * @var integer
     */
    public $ownerId;

    /**
     * @description productType
     * @var string
     */
    public $productType;

    /**
     * @description subscriptionType
     * @var string
     */
    public $subscriptionType;

    /**
     * @description modifyType
     * @var string
     */
    public $modifyType;

    /**
     * @description instanceId
     * @var string
     */
    public $instanceId;

    /**
     * @description parameter
     * @var array
     */
    public $parameter;

    /**
     * @description clientToken
     * @var string
     */
    public $clientToken;

}

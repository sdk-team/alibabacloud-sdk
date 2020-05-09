<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryMonthlyInstanceConsumptionResponse\data\items;

use AlibabaCloud\Tea\Model;

class item extends Model {
    protected $_name = [
        'instanceID' => 'InstanceID',
        'productCode' => 'ProductCode',
        'productType' => 'ProductType',
        'subscriptionType' => 'SubscriptionType',
        'tag' => 'Tag',
        'resourceGroup' => 'ResourceGroup',
        'payerAccount' => 'PayerAccount',
        'ownerID' => 'OwnerID',
        'region' => 'Region',
        'pretaxGrossAmount' => 'PretaxGrossAmount',
        'discountAmount' => 'DiscountAmount',
        'pretaxAmount' => 'PretaxAmount',
        'currency' => 'Currency',
        'pretaxAmountLocal' => 'PretaxAmountLocal',
        'tax' => 'Tax',
        'afterTaxAmount' => 'AfterTaxAmount',
        'paymentCurrency' => 'PaymentCurrency',
    ];
    public function validate() {
        Model::validateRequired('instanceID', $this->instanceID, true);
        Model::validateRequired('productCode', $this->productCode, true);
        Model::validateRequired('productType', $this->productType, true);
        Model::validateRequired('subscriptionType', $this->subscriptionType, true);
        Model::validateRequired('tag', $this->tag, true);
        Model::validateRequired('resourceGroup', $this->resourceGroup, true);
        Model::validateRequired('payerAccount', $this->payerAccount, true);
        Model::validateRequired('ownerID', $this->ownerID, true);
        Model::validateRequired('region', $this->region, true);
        Model::validateRequired('pretaxGrossAmount', $this->pretaxGrossAmount, true);
        Model::validateRequired('discountAmount', $this->discountAmount, true);
        Model::validateRequired('pretaxAmount', $this->pretaxAmount, true);
        Model::validateRequired('currency', $this->currency, true);
        Model::validateRequired('pretaxAmountLocal', $this->pretaxAmountLocal, true);
        Model::validateRequired('tax', $this->tax, true);
        Model::validateRequired('afterTaxAmount', $this->afterTaxAmount, true);
        Model::validateRequired('paymentCurrency', $this->paymentCurrency, true);
    }
    public function toMap() {
        $res = [];
        $res['InstanceID'] = $this->instanceID;
        $res['ProductCode'] = $this->productCode;
        $res['ProductType'] = $this->productType;
        $res['SubscriptionType'] = $this->subscriptionType;
        $res['Tag'] = $this->tag;
        $res['ResourceGroup'] = $this->resourceGroup;
        $res['PayerAccount'] = $this->payerAccount;
        $res['OwnerID'] = $this->ownerID;
        $res['Region'] = $this->region;
        $res['PretaxGrossAmount'] = $this->pretaxGrossAmount;
        $res['DiscountAmount'] = $this->discountAmount;
        $res['PretaxAmount'] = $this->pretaxAmount;
        $res['Currency'] = $this->currency;
        $res['PretaxAmountLocal'] = $this->pretaxAmountLocal;
        $res['Tax'] = $this->tax;
        $res['AfterTaxAmount'] = $this->afterTaxAmount;
        $res['PaymentCurrency'] = $this->paymentCurrency;
        return $res;
    }
    /**
     * @param array $map
     * @return item
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['InstanceID'])){
            $model->instanceID = $map['InstanceID'];
        }
        if(isset($map['ProductCode'])){
            $model->productCode = $map['ProductCode'];
        }
        if(isset($map['ProductType'])){
            $model->productType = $map['ProductType'];
        }
        if(isset($map['SubscriptionType'])){
            $model->subscriptionType = $map['SubscriptionType'];
        }
        if(isset($map['Tag'])){
            $model->tag = $map['Tag'];
        }
        if(isset($map['ResourceGroup'])){
            $model->resourceGroup = $map['ResourceGroup'];
        }
        if(isset($map['PayerAccount'])){
            $model->payerAccount = $map['PayerAccount'];
        }
        if(isset($map['OwnerID'])){
            $model->ownerID = $map['OwnerID'];
        }
        if(isset($map['Region'])){
            $model->region = $map['Region'];
        }
        if(isset($map['PretaxGrossAmount'])){
            $model->pretaxGrossAmount = $map['PretaxGrossAmount'];
        }
        if(isset($map['DiscountAmount'])){
            $model->discountAmount = $map['DiscountAmount'];
        }
        if(isset($map['PretaxAmount'])){
            $model->pretaxAmount = $map['PretaxAmount'];
        }
        if(isset($map['Currency'])){
            $model->currency = $map['Currency'];
        }
        if(isset($map['PretaxAmountLocal'])){
            $model->pretaxAmountLocal = $map['PretaxAmountLocal'];
        }
        if(isset($map['Tax'])){
            $model->tax = $map['Tax'];
        }
        if(isset($map['AfterTaxAmount'])){
            $model->afterTaxAmount = $map['AfterTaxAmount'];
        }
        if(isset($map['PaymentCurrency'])){
            $model->paymentCurrency = $map['PaymentCurrency'];
        }
        return $model;
    }
    /**
     * @description instanceID
     * @var string
     */
    public $instanceID;

    /**
     * @description productCode
     * @var string
     */
    public $productCode;

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
     * @description tag
     * @var string
     */
    public $tag;

    /**
     * @description resourceGroup
     * @var string
     */
    public $resourceGroup;

    /**
     * @description payerAccount
     * @var string
     */
    public $payerAccount;

    /**
     * @description ownerID
     * @var string
     */
    public $ownerID;

    /**
     * @description region
     * @var string
     */
    public $region;

    /**
     * @description pretaxGrossAmount
     * @var float
     */
    public $pretaxGrossAmount;

    /**
     * @description discountAmount
     * @var float
     */
    public $discountAmount;

    /**
     * @description pretaxAmount
     * @var float
     */
    public $pretaxAmount;

    /**
     * @description currency
     * @var string
     */
    public $currency;

    /**
     * @description pretaxAmountLocal
     * @var float
     */
    public $pretaxAmountLocal;

    /**
     * @description tax
     * @var float
     */
    public $tax;

    /**
     * @description afterTaxAmount
     * @var float
     */
    public $afterTaxAmount;

    /**
     * @description paymentCurrency
     * @var string
     */
    public $paymentCurrency;

}

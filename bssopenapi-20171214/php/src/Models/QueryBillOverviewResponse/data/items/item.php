<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryBillOverviewResponse\data\items;

use AlibabaCloud\Tea\Model;

class item extends Model {
    protected $_name = [
        'item' => 'Item',
        'productCode' => 'ProductCode',
        'productType' => 'ProductType',
        'subscriptionType' => 'SubscriptionType',
        'productName' => 'ProductName',
        'productDetail' => 'ProductDetail',
        'pretaxGrossAmount' => 'PretaxGrossAmount',
        'invoiceDiscount' => 'InvoiceDiscount',
        'deductedByCoupons' => 'DeductedByCoupons',
        'pretaxAmount' => 'PretaxAmount',
        'currency' => 'Currency',
        'paymentAmount' => 'PaymentAmount',
        'outstandingAmount' => 'OutstandingAmount',
        'deductedByCashCoupons' => 'DeductedByCashCoupons',
        'deductedByPrepaidCard' => 'DeductedByPrepaidCard',
        'pretaxAmountLocal' => 'PretaxAmountLocal',
        'tax' => 'Tax',
        'afterTaxAmount' => 'AfterTaxAmount',
        'paymentCurrency' => 'PaymentCurrency',
        'roundDownDiscount' => 'RoundDownDiscount',
    ];
    public function validate() {
        Model::validateRequired('item', $this->item, true);
        Model::validateRequired('productCode', $this->productCode, true);
        Model::validateRequired('productType', $this->productType, true);
        Model::validateRequired('subscriptionType', $this->subscriptionType, true);
        Model::validateRequired('productName', $this->productName, true);
        Model::validateRequired('productDetail', $this->productDetail, true);
        Model::validateRequired('pretaxGrossAmount', $this->pretaxGrossAmount, true);
        Model::validateRequired('invoiceDiscount', $this->invoiceDiscount, true);
        Model::validateRequired('deductedByCoupons', $this->deductedByCoupons, true);
        Model::validateRequired('pretaxAmount', $this->pretaxAmount, true);
        Model::validateRequired('currency', $this->currency, true);
        Model::validateRequired('paymentAmount', $this->paymentAmount, true);
        Model::validateRequired('outstandingAmount', $this->outstandingAmount, true);
        Model::validateRequired('deductedByCashCoupons', $this->deductedByCashCoupons, true);
        Model::validateRequired('deductedByPrepaidCard', $this->deductedByPrepaidCard, true);
        Model::validateRequired('pretaxAmountLocal', $this->pretaxAmountLocal, true);
        Model::validateRequired('tax', $this->tax, true);
        Model::validateRequired('afterTaxAmount', $this->afterTaxAmount, true);
        Model::validateRequired('paymentCurrency', $this->paymentCurrency, true);
        Model::validateRequired('roundDownDiscount', $this->roundDownDiscount, true);
    }
    public function toMap() {
        $res = [];
        $res['Item'] = $this->item;
        $res['ProductCode'] = $this->productCode;
        $res['ProductType'] = $this->productType;
        $res['SubscriptionType'] = $this->subscriptionType;
        $res['ProductName'] = $this->productName;
        $res['ProductDetail'] = $this->productDetail;
        $res['PretaxGrossAmount'] = $this->pretaxGrossAmount;
        $res['InvoiceDiscount'] = $this->invoiceDiscount;
        $res['DeductedByCoupons'] = $this->deductedByCoupons;
        $res['PretaxAmount'] = $this->pretaxAmount;
        $res['Currency'] = $this->currency;
        $res['PaymentAmount'] = $this->paymentAmount;
        $res['OutstandingAmount'] = $this->outstandingAmount;
        $res['DeductedByCashCoupons'] = $this->deductedByCashCoupons;
        $res['DeductedByPrepaidCard'] = $this->deductedByPrepaidCard;
        $res['PretaxAmountLocal'] = $this->pretaxAmountLocal;
        $res['Tax'] = $this->tax;
        $res['AfterTaxAmount'] = $this->afterTaxAmount;
        $res['PaymentCurrency'] = $this->paymentCurrency;
        $res['RoundDownDiscount'] = $this->roundDownDiscount;
        return $res;
    }
    /**
     * @param array $map
     * @return item
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Item'])){
            $model->item = $map['Item'];
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
        if(isset($map['ProductName'])){
            $model->productName = $map['ProductName'];
        }
        if(isset($map['ProductDetail'])){
            $model->productDetail = $map['ProductDetail'];
        }
        if(isset($map['PretaxGrossAmount'])){
            $model->pretaxGrossAmount = $map['PretaxGrossAmount'];
        }
        if(isset($map['InvoiceDiscount'])){
            $model->invoiceDiscount = $map['InvoiceDiscount'];
        }
        if(isset($map['DeductedByCoupons'])){
            $model->deductedByCoupons = $map['DeductedByCoupons'];
        }
        if(isset($map['PretaxAmount'])){
            $model->pretaxAmount = $map['PretaxAmount'];
        }
        if(isset($map['Currency'])){
            $model->currency = $map['Currency'];
        }
        if(isset($map['PaymentAmount'])){
            $model->paymentAmount = $map['PaymentAmount'];
        }
        if(isset($map['OutstandingAmount'])){
            $model->outstandingAmount = $map['OutstandingAmount'];
        }
        if(isset($map['DeductedByCashCoupons'])){
            $model->deductedByCashCoupons = $map['DeductedByCashCoupons'];
        }
        if(isset($map['DeductedByPrepaidCard'])){
            $model->deductedByPrepaidCard = $map['DeductedByPrepaidCard'];
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
        if(isset($map['RoundDownDiscount'])){
            $model->roundDownDiscount = $map['RoundDownDiscount'];
        }
        return $model;
    }
    /**
     * @description item
     * @var string
     */
    public $item;

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
     * @description productName
     * @var string
     */
    public $productName;

    /**
     * @description productDetail
     * @var string
     */
    public $productDetail;

    /**
     * @description pretaxGrossAmount
     * @var float
     */
    public $pretaxGrossAmount;

    /**
     * @description invoiceDiscount
     * @var float
     */
    public $invoiceDiscount;

    /**
     * @description deductedByCoupons
     * @var float
     */
    public $deductedByCoupons;

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
     * @description paymentAmount
     * @var float
     */
    public $paymentAmount;

    /**
     * @description outstandingAmount
     * @var float
     */
    public $outstandingAmount;

    /**
     * @description deductedByCashCoupons
     * @var float
     */
    public $deductedByCashCoupons;

    /**
     * @description deductedByPrepaidCard
     * @var float
     */
    public $deductedByPrepaidCard;

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

    /**
     * @description roundDownDiscount
     * @var string
     */
    public $roundDownDiscount;

}

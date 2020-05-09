<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryResourcePackageInstancesResponse\data\instances;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\instance\applicableProducts;

class instance extends Model {
    protected $_name = [
        'instanceId' => 'InstanceId',
        'region' => 'Region',
        'totalAmount' => 'TotalAmount',
        'totalAmountUnit' => 'TotalAmountUnit',
        'remainingAmount' => 'RemainingAmount',
        'remainingAmountUnit' => 'RemainingAmountUnit',
        'effectiveTime' => 'EffectiveTime',
        'expiryTime' => 'ExpiryTime',
        'remark' => 'Remark',
        'packageType' => 'PackageType',
        'status' => 'Status',
        'deductType' => 'DeductType',
        'applicableProducts' => 'ApplicableProducts',
    ];
    public function validate() {
        Model::validateRequired('instanceId', $this->instanceId, true);
        Model::validateRequired('region', $this->region, true);
        Model::validateRequired('totalAmount', $this->totalAmount, true);
        Model::validateRequired('totalAmountUnit', $this->totalAmountUnit, true);
        Model::validateRequired('remainingAmount', $this->remainingAmount, true);
        Model::validateRequired('remainingAmountUnit', $this->remainingAmountUnit, true);
        Model::validateRequired('effectiveTime', $this->effectiveTime, true);
        Model::validateRequired('expiryTime', $this->expiryTime, true);
        Model::validateRequired('remark', $this->remark, true);
        Model::validateRequired('packageType', $this->packageType, true);
        Model::validateRequired('status', $this->status, true);
        Model::validateRequired('deductType', $this->deductType, true);
        Model::validateRequired('applicableProducts', $this->applicableProducts, true);
    }
    public function toMap() {
        $res = [];
        $res['InstanceId'] = $this->instanceId;
        $res['Region'] = $this->region;
        $res['TotalAmount'] = $this->totalAmount;
        $res['TotalAmountUnit'] = $this->totalAmountUnit;
        $res['RemainingAmount'] = $this->remainingAmount;
        $res['RemainingAmountUnit'] = $this->remainingAmountUnit;
        $res['EffectiveTime'] = $this->effectiveTime;
        $res['ExpiryTime'] = $this->expiryTime;
        $res['Remark'] = $this->remark;
        $res['PackageType'] = $this->packageType;
        $res['Status'] = $this->status;
        $res['DeductType'] = $this->deductType;
        $res['ApplicableProducts'] = null !== $this->applicableProducts ? $this->applicableProducts->toMap() : null;
        return $res;
    }
    /**
     * @param array $map
     * @return instance
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['InstanceId'])){
            $model->instanceId = $map['InstanceId'];
        }
        if(isset($map['Region'])){
            $model->region = $map['Region'];
        }
        if(isset($map['TotalAmount'])){
            $model->totalAmount = $map['TotalAmount'];
        }
        if(isset($map['TotalAmountUnit'])){
            $model->totalAmountUnit = $map['TotalAmountUnit'];
        }
        if(isset($map['RemainingAmount'])){
            $model->remainingAmount = $map['RemainingAmount'];
        }
        if(isset($map['RemainingAmountUnit'])){
            $model->remainingAmountUnit = $map['RemainingAmountUnit'];
        }
        if(isset($map['EffectiveTime'])){
            $model->effectiveTime = $map['EffectiveTime'];
        }
        if(isset($map['ExpiryTime'])){
            $model->expiryTime = $map['ExpiryTime'];
        }
        if(isset($map['Remark'])){
            $model->remark = $map['Remark'];
        }
        if(isset($map['PackageType'])){
            $model->packageType = $map['PackageType'];
        }
        if(isset($map['Status'])){
            $model->status = $map['Status'];
        }
        if(isset($map['DeductType'])){
            $model->deductType = $map['DeductType'];
        }
        if(isset($map['ApplicableProducts'])){
            $model->applicableProducts = applicableProducts::fromMap($map['ApplicableProducts']);
        }
        return $model;
    }
    /**
     * @description instanceId
     * @var string
     */
    public $instanceId;

    /**
     * @description region
     * @var string
     */
    public $region;

    /**
     * @description totalAmount
     * @var string
     */
    public $totalAmount;

    /**
     * @description totalAmountUnit
     * @var string
     */
    public $totalAmountUnit;

    /**
     * @description remainingAmount
     * @var string
     */
    public $remainingAmount;

    /**
     * @description remainingAmountUnit
     * @var string
     */
    public $remainingAmountUnit;

    /**
     * @description effectiveFrom
     * @var string
     */
    public $effectiveTime;

    /**
     * @description effectiveTo
     * @var string
     */
    public $expiryTime;

    /**
     * @description remark
     * @var string
     */
    public $remark;

    /**
     * @description templateCode
     * @var string
     */
    public $packageType;

    /**
     * @description status
     * @var string
     */
    public $status;

    /**
     * @description deductType
     * @var string
     */
    public $deductType;

    /**
     * @description resourceTypes
     * @var instance.applicableProducts
     */
    public $applicableProducts;

}

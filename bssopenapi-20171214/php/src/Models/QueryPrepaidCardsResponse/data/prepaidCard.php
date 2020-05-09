<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryPrepaidCardsResponse\data;

use AlibabaCloud\Tea\Model;

class prepaidCard extends Model {
    protected $_name = [
        'prepaidCardId' => 'PrepaidCardId',
        'prepaidCardNo' => 'PrepaidCardNo',
        'grantedTime' => 'GrantedTime',
        'effectiveTime' => 'EffectiveTime',
        'expiryTime' => 'ExpiryTime',
        'applicableProducts' => 'ApplicableProducts',
        'applicableScenarios' => 'ApplicableScenarios',
        'nominalValue' => 'NominalValue',
        'balance' => 'Balance',
        'status' => 'Status',
    ];
    public function validate() {
        Model::validateRequired('prepaidCardId', $this->prepaidCardId, true);
        Model::validateRequired('prepaidCardNo', $this->prepaidCardNo, true);
        Model::validateRequired('grantedTime', $this->grantedTime, true);
        Model::validateRequired('effectiveTime', $this->effectiveTime, true);
        Model::validateRequired('expiryTime', $this->expiryTime, true);
        Model::validateRequired('applicableProducts', $this->applicableProducts, true);
        Model::validateRequired('applicableScenarios', $this->applicableScenarios, true);
        Model::validateRequired('nominalValue', $this->nominalValue, true);
        Model::validateRequired('balance', $this->balance, true);
        Model::validateRequired('status', $this->status, true);
    }
    public function toMap() {
        $res = [];
        $res['PrepaidCardId'] = $this->prepaidCardId;
        $res['PrepaidCardNo'] = $this->prepaidCardNo;
        $res['GrantedTime'] = $this->grantedTime;
        $res['EffectiveTime'] = $this->effectiveTime;
        $res['ExpiryTime'] = $this->expiryTime;
        $res['ApplicableProducts'] = $this->applicableProducts;
        $res['ApplicableScenarios'] = $this->applicableScenarios;
        $res['NominalValue'] = $this->nominalValue;
        $res['Balance'] = $this->balance;
        $res['Status'] = $this->status;
        return $res;
    }
    /**
     * @param array $map
     * @return prepaidCard
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['PrepaidCardId'])){
            $model->prepaidCardId = $map['PrepaidCardId'];
        }
        if(isset($map['PrepaidCardNo'])){
            $model->prepaidCardNo = $map['PrepaidCardNo'];
        }
        if(isset($map['GrantedTime'])){
            $model->grantedTime = $map['GrantedTime'];
        }
        if(isset($map['EffectiveTime'])){
            $model->effectiveTime = $map['EffectiveTime'];
        }
        if(isset($map['ExpiryTime'])){
            $model->expiryTime = $map['ExpiryTime'];
        }
        if(isset($map['ApplicableProducts'])){
            $model->applicableProducts = $map['ApplicableProducts'];
        }
        if(isset($map['ApplicableScenarios'])){
            $model->applicableScenarios = $map['ApplicableScenarios'];
        }
        if(isset($map['NominalValue'])){
            $model->nominalValue = $map['NominalValue'];
        }
        if(isset($map['Balance'])){
            $model->balance = $map['Balance'];
        }
        if(isset($map['Status'])){
            $model->status = $map['Status'];
        }
        return $model;
    }
    /**
     * @description prepaidCardId
     * @var integer
     */
    public $prepaidCardId;

    /**
     * @description prepaidCardNo
     * @var string
     */
    public $prepaidCardNo;

    /**
     * @description grantedTime
     * @var string
     */
    public $grantedTime;

    /**
     * @description effectiveTime
     * @var string
     */
    public $effectiveTime;

    /**
     * @description expiryTime
     * @var string
     */
    public $expiryTime;

    /**
     * @description applicableProducts
     * @var string
     */
    public $applicableProducts;

    /**
     * @description applicableScenarios
     * @var string
     */
    public $applicableScenarios;

    /**
     * @description nominalValue
     * @var string
     */
    public $nominalValue;

    /**
     * @description balance
     * @var string
     */
    public $balance;

    /**
     * @description status
     * @var string
     */
    public $status;

}

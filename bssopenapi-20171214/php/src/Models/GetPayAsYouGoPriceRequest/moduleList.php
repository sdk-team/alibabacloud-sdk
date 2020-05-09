<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\GetPayAsYouGoPriceRequest;

use AlibabaCloud\Tea\Model;

class moduleList extends Model {
    protected $_name = [
        'moduleCode' => 'ModuleCode',
        'config' => 'Config',
        'priceType' => 'PriceType',
    ];
    public function validate() {
        Model::validateRequired('moduleCode', $this->moduleCode, true);
        Model::validateRequired('config', $this->config, true);
        Model::validateRequired('priceType', $this->priceType, true);
    }
    public function toMap() {
        $res = [];
        $res['ModuleCode'] = $this->moduleCode;
        $res['Config'] = $this->config;
        $res['PriceType'] = $this->priceType;
        return $res;
    }
    /**
     * @param array $map
     * @return moduleList
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['ModuleCode'])){
            $model->moduleCode = $map['ModuleCode'];
        }
        if(isset($map['Config'])){
            $model->config = $map['Config'];
        }
        if(isset($map['PriceType'])){
            $model->priceType = $map['PriceType'];
        }
        return $model;
    }
    /**
     * @description moduleCode
     * @var string
     */
    public $moduleCode;

    /**
     * @description config
     * @var string
     */
    public $config;

    /**
     * @description priceType
     * @var string
     */
    public $priceType;

}

<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\GetSubscriptionPriceRequest;

use AlibabaCloud\Tea\Model;

class moduleList extends Model {
    protected $_name = [
        'moduleCode' => 'ModuleCode',
        'config' => 'Config',
        'moduleStatus' => 'ModuleStatus',
        'tag' => 'Tag',
    ];
    public function validate() {
        Model::validateRequired('moduleCode', $this->moduleCode, true);
        Model::validateRequired('config', $this->config, true);
    }
    public function toMap() {
        $res = [];
        $res['ModuleCode'] = $this->moduleCode;
        $res['Config'] = $this->config;
        $res['ModuleStatus'] = $this->moduleStatus;
        $res['Tag'] = $this->tag;
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
        if(isset($map['ModuleStatus'])){
            $model->moduleStatus = $map['ModuleStatus'];
        }
        if(isset($map['Tag'])){
            $model->tag = $map['Tag'];
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
     * @description moduleStatus
     * @var integer
     */
    public $moduleStatus;

    /**
     * @description tag
     * @var string
     */
    public $tag;

}

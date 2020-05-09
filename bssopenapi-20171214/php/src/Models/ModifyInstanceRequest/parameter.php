<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\ModifyInstanceRequest;

use AlibabaCloud\Tea\Model;

class parameter extends Model {
    protected $_name = [
        'code' => 'Code',
        'value' => 'Value',
    ];
    public function validate() {
        Model::validateRequired('code', $this->code, true);
        Model::validateRequired('value', $this->value, true);
    }
    public function toMap() {
        $res = [];
        $res['Code'] = $this->code;
        $res['Value'] = $this->value;
        return $res;
    }
    /**
     * @param array $map
     * @return parameter
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Code'])){
            $model->code = $map['Code'];
        }
        if(isset($map['Value'])){
            $model->value = $map['Value'];
        }
        return $model;
    }
    /**
     * @description code
     * @var string
     */
    public $code;

    /**
     * @description value
     * @var string
     */
    public $value;

}

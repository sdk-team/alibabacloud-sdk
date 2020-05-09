<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models;

use AlibabaCloud\Tea\Model;

class DescribeResourcePackageProductRequest extends Model {
    protected $_name = [
        'productCode' => 'ProductCode',
    ];
    public function validate() {}
    public function toMap() {
        $res = [];
        $res['ProductCode'] = $this->productCode;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeResourcePackageProductRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['ProductCode'])){
            $model->productCode = $map['ProductCode'];
        }
        return $model;
    }
    /**
     * @description productCode
     * @var string
     */
    public $productCode;

}

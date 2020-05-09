<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\GetResourcePackagePriceResponse\data\promotions;

use AlibabaCloud\Tea\Model;

class promotion extends Model {
    protected $_name = [
        'id' => 'Id',
        'name' => 'Name',
    ];
    public function validate() {
        Model::validateRequired('id', $this->id, true);
        Model::validateRequired('name', $this->name, true);
    }
    public function toMap() {
        $res = [];
        $res['Id'] = $this->id;
        $res['Name'] = $this->name;
        return $res;
    }
    /**
     * @param array $map
     * @return promotion
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Id'])){
            $model->id = $map['Id'];
        }
        if(isset($map['Name'])){
            $model->name = $map['Name'];
        }
        return $model;
    }
    /**
     * @description id
     * @var integer
     */
    public $id;

    /**
     * @description name
     * @var string
     */
    public $name;

}

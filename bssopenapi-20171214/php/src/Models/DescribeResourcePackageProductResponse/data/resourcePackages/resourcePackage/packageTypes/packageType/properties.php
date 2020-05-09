<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\DescribeResourcePackageProductResponse\data\resourcePackages\resourcePackage\packageTypes\packageType;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\DescribeResourcePackageProductResponse\data\resourcePackages\resourcePackage\packageTypes\packageType\properties\property;

class properties extends Model {
    protected $_name = [
        'property' => 'Property',
    ];
    public function validate() {
        Model::validateRequired('property', $this->property, true);
    }
    public function toMap() {
        $res = [];
        $res['Property'] = [];
        if(null !== $this->property && is_array($this->property)){
            $n = 0;
            foreach($this->property as $item){
                $res['Property'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        return $res;
    }
    /**
     * @param array $map
     * @return properties
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Property'])){
            if(!empty($map['Property'])){
                $model->property = [];
                $n = 0;
                foreach($map['Property'] as $item) {
                    $model->property[$n++] = null !== $item ? property::fromMap($item) : $item;
                }
            }
        }
        return $model;
    }
    /**
     * @description Property
     * @var array
     */
    public $property;

}

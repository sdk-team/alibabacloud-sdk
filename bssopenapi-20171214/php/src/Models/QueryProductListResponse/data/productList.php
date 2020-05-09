<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryProductListResponse\data;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\QueryProductListResponse\data\productList\product;

class productList extends Model {
    protected $_name = [
        'product' => 'Product',
    ];
    public function validate() {
        Model::validateRequired('product', $this->product, true);
    }
    public function toMap() {
        $res = [];
        $res['Product'] = [];
        if(null !== $this->product && is_array($this->product)){
            $n = 0;
            foreach($this->product as $item){
                $res['Product'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        return $res;
    }
    /**
     * @param array $map
     * @return productList
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Product'])){
            if(!empty($map['Product'])){
                $model->product = [];
                $n = 0;
                foreach($map['Product'] as $item) {
                    $model->product[$n++] = null !== $item ? product::fromMap($item) : $item;
                }
            }
        }
        return $model;
    }
    /**
     * @description Product
     * @var array
     */
    public $product;

}

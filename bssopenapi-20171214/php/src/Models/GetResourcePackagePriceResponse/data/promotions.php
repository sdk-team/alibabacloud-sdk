<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models\GetResourcePackagePriceResponse\data;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\GetResourcePackagePriceResponse\data\promotions\promotion;

class promotions extends Model {
    protected $_name = [
        'promotion' => 'Promotion',
    ];
    public function validate() {
        Model::validateRequired('promotion', $this->promotion, true);
    }
    public function toMap() {
        $res = [];
        $res['Promotion'] = [];
        if(null !== $this->promotion && is_array($this->promotion)){
            $n = 0;
            foreach($this->promotion as $item){
                $res['Promotion'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        return $res;
    }
    /**
     * @param array $map
     * @return promotions
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['Promotion'])){
            if(!empty($map['Promotion'])){
                $model->promotion = [];
                $n = 0;
                foreach($map['Promotion'] as $item) {
                    $model->promotion[$n++] = null !== $item ? promotion::fromMap($item) : $item;
                }
            }
        }
        return $model;
    }
    /**
     * @description Promotion
     * @var array
     */
    public $promotion;

}

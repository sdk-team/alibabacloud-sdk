<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\BssOpenApi\V20171214\Models;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\BssOpenApi\V20171214\Models\AllocateCostUnitResourceRequest\resourceInstanceList;

class AllocateCostUnitResourceRequest extends Model {
    protected $_name = [
        'fromUnitUserId' => 'FromUnitUserId',
        'fromUnitId' => 'FromUnitId',
        'resourceInstanceList' => 'ResourceInstanceList',
        'toUnitUserId' => 'ToUnitUserId',
        'toUnitId' => 'ToUnitId',
    ];
    public function validate() {
        Model::validateRequired('fromUnitUserId', $this->fromUnitUserId, true);
        Model::validateRequired('fromUnitId', $this->fromUnitId, true);
        Model::validateRequired('resourceInstanceList', $this->resourceInstanceList, true);
        Model::validateRequired('toUnitUserId', $this->toUnitUserId, true);
        Model::validateRequired('toUnitId', $this->toUnitId, true);
    }
    public function toMap() {
        $res = [];
        $res['FromUnitUserId'] = $this->fromUnitUserId;
        $res['FromUnitId'] = $this->fromUnitId;
        $res['ResourceInstanceList'] = [];
        if(null !== $this->resourceInstanceList && is_array($this->resourceInstanceList)){
            $n = 0;
            foreach($this->resourceInstanceList as $item){
                $res['ResourceInstanceList'][$n++] = null !== $item ? $item->toMap() : $item;
            }
        }
        $res['ToUnitUserId'] = $this->toUnitUserId;
        $res['ToUnitId'] = $this->toUnitId;
        return $res;
    }
    /**
     * @param array $map
     * @return AllocateCostUnitResourceRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['FromUnitUserId'])){
            $model->fromUnitUserId = $map['FromUnitUserId'];
        }
        if(isset($map['FromUnitId'])){
            $model->fromUnitId = $map['FromUnitId'];
        }
        if(isset($map['ResourceInstanceList'])){
            if(!empty($map['ResourceInstanceList'])){
                $model->resourceInstanceList = [];
                $n = 0;
                foreach($map['ResourceInstanceList'] as $item) {
                    $model->resourceInstanceList[$n++] = null !== $item ? resourceInstanceList::fromMap($item) : $item;
                }
            }
        }
        if(isset($map['ToUnitUserId'])){
            $model->toUnitUserId = $map['ToUnitUserId'];
        }
        if(isset($map['ToUnitId'])){
            $model->toUnitId = $map['ToUnitId'];
        }
        return $model;
    }
    /**
     * @description fromUnitUserId
     * @var integer
     */
    public $fromUnitUserId;

    /**
     * @description fromUnitId
     * @var integer
     */
    public $fromUnitId;

    /**
     * @description resourceInstanceList
     * @var array
     */
    public $resourceInstanceList;

    /**
     * @description toUnitUserId
     * @var integer
     */
    public $toUnitUserId;

    /**
     * @description toUnitId
     * @var integer
     */
    public $toUnitId;

}

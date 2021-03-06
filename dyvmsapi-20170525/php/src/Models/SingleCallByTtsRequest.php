<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\Dyvmsapi\V20170525\Models;

use AlibabaCloud\Tea\Model;

class SingleCallByTtsRequest extends Model {
    protected $_name = [
        'accessKeyId' => 'AccessKeyId',
        'ownerId' => 'OwnerId',
        'resourceOwnerAccount' => 'ResourceOwnerAccount',
        'resourceOwnerId' => 'ResourceOwnerId',
        'calledShowNumber' => 'CalledShowNumber',
        'calledNumber' => 'CalledNumber',
        'ttsCode' => 'TtsCode',
        'ttsParam' => 'TtsParam',
        'playTimes' => 'PlayTimes',
        'volume' => 'Volume',
        'speed' => 'Speed',
        'outId' => 'OutId',
    ];
    public function validate() {
        Model::validateRequired('calledShowNumber', $this->calledShowNumber, true);
        Model::validateRequired('calledNumber', $this->calledNumber, true);
        Model::validateRequired('ttsCode', $this->ttsCode, true);
    }
    public function toMap() {
        $res = [];
        $res['AccessKeyId'] = $this->accessKeyId;
        $res['OwnerId'] = $this->ownerId;
        $res['ResourceOwnerAccount'] = $this->resourceOwnerAccount;
        $res['ResourceOwnerId'] = $this->resourceOwnerId;
        $res['CalledShowNumber'] = $this->calledShowNumber;
        $res['CalledNumber'] = $this->calledNumber;
        $res['TtsCode'] = $this->ttsCode;
        $res['TtsParam'] = $this->ttsParam;
        $res['PlayTimes'] = $this->playTimes;
        $res['Volume'] = $this->volume;
        $res['Speed'] = $this->speed;
        $res['OutId'] = $this->outId;
        return $res;
    }
    /**
     * @param array $map
     * @return SingleCallByTtsRequest
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['AccessKeyId'])){
            $model->accessKeyId = $map['AccessKeyId'];
        }
        if(isset($map['OwnerId'])){
            $model->ownerId = $map['OwnerId'];
        }
        if(isset($map['ResourceOwnerAccount'])){
            $model->resourceOwnerAccount = $map['ResourceOwnerAccount'];
        }
        if(isset($map['ResourceOwnerId'])){
            $model->resourceOwnerId = $map['ResourceOwnerId'];
        }
        if(isset($map['CalledShowNumber'])){
            $model->calledShowNumber = $map['CalledShowNumber'];
        }
        if(isset($map['CalledNumber'])){
            $model->calledNumber = $map['CalledNumber'];
        }
        if(isset($map['TtsCode'])){
            $model->ttsCode = $map['TtsCode'];
        }
        if(isset($map['TtsParam'])){
            $model->ttsParam = $map['TtsParam'];
        }
        if(isset($map['PlayTimes'])){
            $model->playTimes = $map['PlayTimes'];
        }
        if(isset($map['Volume'])){
            $model->volume = $map['Volume'];
        }
        if(isset($map['Speed'])){
            $model->speed = $map['Speed'];
        }
        if(isset($map['OutId'])){
            $model->outId = $map['OutId'];
        }
        return $model;
    }
    /**
     * @description appKey
     * @var string
     */
    public $accessKeyId;

    /**
     * @description ownerId
     * @var integer
     */
    public $ownerId;

    /**
     * @description resourceOwnerAccount
     * @var string
     */
    public $resourceOwnerAccount;

    /**
     * @description resourceOwnerId
     * @var integer
     */
    public $resourceOwnerId;

    /**
     * @description calledShowNumber
     * @var string
     */
    public $calledShowNumber;

    /**
     * @description calledNumber
     * @var string
     */
    public $calledNumber;

    /**
     * @description ttsCode
     * @var string
     */
    public $ttsCode;

    /**
     * @description ttsParam
     * @var string
     */
    public $ttsParam;

    /**
     * @description playTimes
     * @var integer
     */
    public $playTimes;

    /**
     * @description volume
     * @var integer
     */
    public $volume;

    /**
     * @description speed
     * @var integer
     */
    public $speed;

    /**
     * @description outId
     * @var string
     */
    public $outId;

}

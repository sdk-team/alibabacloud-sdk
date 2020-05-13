<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\vod\V20170314\Models\GetMessageCallbackResponse;

use AlibabaCloud\Tea\Model;

class messageCallback extends Model {
    protected $_name = [
        'callbackType' => 'CallbackType',
        'callbackSwitch' => 'CallbackSwitch',
        'callbackURL' => 'CallbackURL',
        'eventTypeList' => 'EventTypeList',
        'authSwitch' => 'AuthSwitch',
        'authKey' => 'AuthKey',
        'mnsEndpoint' => 'MnsEndpoint',
        'mnsQueueName' => 'MnsQueueName',
    ];
    public function validate() {
        Model::validateRequired('callbackType', $this->callbackType, true);
        Model::validateRequired('callbackSwitch', $this->callbackSwitch, true);
        Model::validateRequired('callbackURL', $this->callbackURL, true);
        Model::validateRequired('eventTypeList', $this->eventTypeList, true);
        Model::validateRequired('authSwitch', $this->authSwitch, true);
        Model::validateRequired('authKey', $this->authKey, true);
        Model::validateRequired('mnsEndpoint', $this->mnsEndpoint, true);
        Model::validateRequired('mnsQueueName', $this->mnsQueueName, true);
    }
    public function toMap() {
        $res = [];
        $res['CallbackType'] = $this->callbackType;
        $res['CallbackSwitch'] = $this->callbackSwitch;
        $res['CallbackURL'] = $this->callbackURL;
        $res['EventTypeList'] = $this->eventTypeList;
        $res['AuthSwitch'] = $this->authSwitch;
        $res['AuthKey'] = $this->authKey;
        $res['MnsEndpoint'] = $this->mnsEndpoint;
        $res['MnsQueueName'] = $this->mnsQueueName;
        return $res;
    }
    /**
     * @param array $map
     * @return messageCallback
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['CallbackType'])){
            $model->callbackType = $map['CallbackType'];
        }
        if(isset($map['CallbackSwitch'])){
            $model->callbackSwitch = $map['CallbackSwitch'];
        }
        if(isset($map['CallbackURL'])){
            $model->callbackURL = $map['CallbackURL'];
        }
        if(isset($map['EventTypeList'])){
            $model->eventTypeList = $map['EventTypeList'];
        }
        if(isset($map['AuthSwitch'])){
            $model->authSwitch = $map['AuthSwitch'];
        }
        if(isset($map['AuthKey'])){
            $model->authKey = $map['AuthKey'];
        }
        if(isset($map['MnsEndpoint'])){
            $model->mnsEndpoint = $map['MnsEndpoint'];
        }
        if(isset($map['MnsQueueName'])){
            $model->mnsQueueName = $map['MnsQueueName'];
        }
        return $model;
    }
    /**
     * @description callbackType
     * @var string
     */
    public $callbackType;

    /**
     * @description callbackSwitch
     * @var string
     */
    public $callbackSwitch;

    /**
     * @description uri
     * @var string
     */
    public $callbackURL;

    /**
     * @description eventTypeList
     * @var string
     */
    public $eventTypeList;

    /**
     * @description authSwitch
     * @var string
     */
    public $authSwitch;

    /**
     * @description httpAuthKey
     * @var string
     */
    public $authKey;

    /**
     * @description mnsEndpoint
     * @var string
     */
    public $mnsEndpoint;

    /**
     * @description mnsQueueName
     * @var string
     */
    public $mnsQueueName;

}

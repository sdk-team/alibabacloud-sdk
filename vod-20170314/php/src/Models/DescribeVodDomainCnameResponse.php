<?php

// This file is auto-generated, don't edit it. Thanks.
namespace AlibabaCloud\SDK\vod\V20170314\Models;

use AlibabaCloud\Tea\Model;

use AlibabaCloud\SDK\vod\V20170314\Models\DescribeVodDomainCnameResponse\cnameDatas;

class DescribeVodDomainCnameResponse extends Model {
    protected $_name = [
        'requestId' => 'RequestId',
        'cnameDatas' => 'CnameDatas',
    ];
    public function validate() {
        Model::validateRequired('requestId', $this->requestId, true);
        Model::validateRequired('cnameDatas', $this->cnameDatas, true);
    }
    public function toMap() {
        $res = [];
        $res['RequestId'] = $this->requestId;
        $res['CnameDatas'] = null !== $this->cnameDatas ? $this->cnameDatas->toMap() : null;
        return $res;
    }
    /**
     * @param array $map
     * @return DescribeVodDomainCnameResponse
     */
    public static function fromMap($map = []) {
        $model = new self();
        if(isset($map['RequestId'])){
            $model->requestId = $map['RequestId'];
        }
        if(isset($map['CnameDatas'])){
            $model->cnameDatas = cnameDatas::fromMap($map['CnameDatas']);
        }
        return $model;
    }
    /**
     * @description requestId
     * @var string
     */
    public $requestId;

    /**
     * @description data.content
     * @var DescribeVodDomainCnameResponse.cnameDatas
     */
    public $cnameDatas;

}

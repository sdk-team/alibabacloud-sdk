<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\ECI\V20180808\ECI\DescribeMultiContainerGroupMetricResponse\monitorDatas\containerInfos\containerStats\diskIoStats;

use AlibabaCloud\Tea\Model;

class ioWaitTime extends Model
{
    public $device;

    public $major;

    public $minor;

    public $stats;
    protected $_name = [
        'device' => 'Device',
        'major'  => 'Major',
        'minor'  => 'Minor',
        'stats'  => 'Stats',
    ];
}
// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20190610.models;

import com.aliyun.tea.*;

public class BatchSetVodDomainConfigsResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static BatchSetVodDomainConfigsResponse build(java.util.Map<String, ?> map) throws Exception {
        BatchSetVodDomainConfigsResponse self = new BatchSetVodDomainConfigsResponse();
        return TeaModel.build(map, self);
    }

}

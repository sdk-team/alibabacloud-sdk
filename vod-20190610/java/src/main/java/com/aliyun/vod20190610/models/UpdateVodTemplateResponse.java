// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20190610.models;

import com.aliyun.tea.*;

public class UpdateVodTemplateResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("VodTemplateId")
    @Validation(required = true)
    public String vodTemplateId;

    public static UpdateVodTemplateResponse build(java.util.Map<String, ?> map) throws Exception {
        UpdateVodTemplateResponse self = new UpdateVodTemplateResponse();
        return TeaModel.build(map, self);
    }

}
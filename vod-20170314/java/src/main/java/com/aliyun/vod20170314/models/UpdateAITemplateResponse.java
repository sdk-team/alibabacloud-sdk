// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170314.models;

import com.aliyun.tea.*;

public class UpdateAITemplateResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("TemplateId")
    @Validation(required = true)
    public String templateId;

    public static UpdateAITemplateResponse build(java.util.Map<String, ?> map) throws Exception {
        UpdateAITemplateResponse self = new UpdateAITemplateResponse();
        return TeaModel.build(map, self);
    }

}
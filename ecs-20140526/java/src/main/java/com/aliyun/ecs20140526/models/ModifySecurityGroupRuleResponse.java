// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ecs20140526.models;

import com.aliyun.tea.*;

public class ModifySecurityGroupRuleResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    public static ModifySecurityGroupRuleResponse build(java.util.Map<String, ?> map) throws Exception {
        ModifySecurityGroupRuleResponse self = new ModifySecurityGroupRuleResponse();
        return TeaModel.build(map, self);
    }

}

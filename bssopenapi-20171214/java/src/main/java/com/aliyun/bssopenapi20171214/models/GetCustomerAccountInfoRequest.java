// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.bssopenapi20171214.models;

import com.aliyun.tea.*;

public class GetCustomerAccountInfoRequest extends TeaModel {
    @NameInMap("OwnerId")
    @Validation(required = true)
    public Long ownerId;

    public static GetCustomerAccountInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        GetCustomerAccountInfoRequest self = new GetCustomerAccountInfoRequest();
        return TeaModel.build(map, self);
    }

}
// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.bss20140714.models;

import com.aliyun.tea.*;

public class vnoBatchRefundOrderResponse extends TeaModel {
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    @NameInMap("Success")
    @Validation(required = true)
    public Boolean success;

    @NameInMap("Code")
    @Validation(required = true)
    public String code;

    @NameInMap("Message")
    @Validation(required = true)
    public String message;

    @NameInMap("Data")
    @Validation(required = true)
    public String data;

    public static vnoBatchRefundOrderResponse build(java.util.Map<String, ?> map) throws Exception {
        vnoBatchRefundOrderResponse self = new vnoBatchRefundOrderResponse();
        return TeaModel.build(map, self);
    }

}

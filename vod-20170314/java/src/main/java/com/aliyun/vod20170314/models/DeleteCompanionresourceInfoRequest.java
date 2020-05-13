// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170314.models;

import com.aliyun.tea.*;

public class DeleteCompanionresourceInfoRequest extends TeaModel {
    @NameInMap("AccessKeyId")
    public String accessKeyId;

    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("CompanionResourceIds")
    @Validation(required = true)
    public String companionResourceIds;

    public static DeleteCompanionresourceInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteCompanionresourceInfoRequest self = new DeleteCompanionresourceInfoRequest();
        return TeaModel.build(map, self);
    }

}
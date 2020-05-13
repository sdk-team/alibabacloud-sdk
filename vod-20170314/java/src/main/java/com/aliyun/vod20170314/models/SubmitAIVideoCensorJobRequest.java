// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.vod20170314.models;

import com.aliyun.tea.*;

public class SubmitAIVideoCensorJobRequest extends TeaModel {
    @NameInMap("AccessKeyId")
    public String accessKeyId;

    @NameInMap("OwnerId")
    public String ownerId;

    @NameInMap("ResourceOwnerId")
    public String resourceOwnerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    @NameInMap("MediaId")
    @Validation(required = true)
    public String mediaId;

    @NameInMap("AIVideoCensorConfig")
    public String AIVideoCensorConfig;

    @NameInMap("UserData")
    public String userData;

    public static SubmitAIVideoCensorJobRequest build(java.util.Map<String, ?> map) throws Exception {
        SubmitAIVideoCensorJobRequest self = new SubmitAIVideoCensorJobRequest();
        return TeaModel.build(map, self);
    }

}
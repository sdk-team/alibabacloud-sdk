// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ecs20140526.models;

import com.aliyun.tea.*;

public class ModifySecurityGroupAttributeRequest extends TeaModel {
    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("SecurityGroupId")
    @Validation(required = true)
    public String securityGroupId;

    @NameInMap("Description")
    public String description;

    @NameInMap("SecurityGroupName")
    public String securityGroupName;

    @NameInMap("RegionId")
    @Validation(required = true)
    public String regionId;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    public static ModifySecurityGroupAttributeRequest build(java.util.Map<String, ?> map) throws Exception {
        ModifySecurityGroupAttributeRequest self = new ModifySecurityGroupAttributeRequest();
        return TeaModel.build(map, self);
    }

}

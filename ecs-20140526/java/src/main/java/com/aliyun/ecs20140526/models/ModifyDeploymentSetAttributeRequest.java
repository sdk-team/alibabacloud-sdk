// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ecs20140526.models;

import com.aliyun.tea.*;

public class ModifyDeploymentSetAttributeRequest extends TeaModel {
    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("DeploymentSetId")
    @Validation(required = true)
    public String deploymentSetId;

    @NameInMap("Description")
    public String description;

    @NameInMap("DeploymentSetName")
    public String deploymentSetName;

    @NameInMap("RegionId")
    @Validation(required = true)
    public String regionId;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    public static ModifyDeploymentSetAttributeRequest build(java.util.Map<String, ?> map) throws Exception {
        ModifyDeploymentSetAttributeRequest self = new ModifyDeploymentSetAttributeRequest();
        return TeaModel.build(map, self);
    }

}

// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ecs20140526.models;

import com.aliyun.tea.*;

public class ModifyBandwidthPackageSpecRequest extends TeaModel {
    @NameInMap("OwnerId")
    public Long ownerId;

    @NameInMap("ResourceOwnerAccount")
    public String resourceOwnerAccount;

    @NameInMap("ResourceOwnerId")
    public Long resourceOwnerId;

    @NameInMap("OwnerAccount")
    public String ownerAccount;

    @NameInMap("RegionId")
    @Validation(required = true)
    public String regionId;

    @NameInMap("BandwidthPackageId")
    @Validation(required = true)
    public String bandwidthPackageId;

    @NameInMap("Bandwidth")
    @Validation(required = true)
    public String bandwidth;

    public static ModifyBandwidthPackageSpecRequest build(java.util.Map<String, ?> map) throws Exception {
        ModifyBandwidthPackageSpecRequest self = new ModifyBandwidthPackageSpecRequest();
        return TeaModel.build(map, self);
    }

}

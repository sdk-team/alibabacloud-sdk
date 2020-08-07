// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.imagesearch20180111.models;

import com.aliyun.tea.*;

public class ListTagResourcesResponse extends TeaModel {
    // requestId
    @NameInMap("RequestId")
    @Validation(required = true)
    public String requestId;

    // data.nextToken
    @NameInMap("NextToken")
    @Validation(required = true)
    public String nextToken;

    // data.tagResources
    @NameInMap("TagResources")
    @Validation(required = true)
    public ListTagResourcesResponseTagResources tagResources;

    public static ListTagResourcesResponse build(java.util.Map<String, ?> map) throws Exception {
        ListTagResourcesResponse self = new ListTagResourcesResponse();
        return TeaModel.build(map, self);
    }

    public static class ListTagResourcesResponseTagResourcesTagResource extends TeaModel {
        // resourceId
        @NameInMap("ResourceId")
        @Validation(required = true)
        public String resourceId;

        // resourceType
        @NameInMap("ResourceType")
        @Validation(required = true)
        public String resourceType;

        // tagKey
        @NameInMap("TagKey")
        @Validation(required = true)
        public String tagKey;

        // tagValue
        @NameInMap("TagValue")
        @Validation(required = true)
        public String tagValue;

        public static ListTagResourcesResponseTagResourcesTagResource build(java.util.Map<String, ?> map) throws Exception {
            ListTagResourcesResponseTagResourcesTagResource self = new ListTagResourcesResponseTagResourcesTagResource();
            return TeaModel.build(map, self);
        }

    }

    public static class ListTagResourcesResponseTagResources extends TeaModel {
        // TagResource
        @NameInMap("TagResource")
        @Validation(required = true)
        public java.util.List<ListTagResourcesResponseTagResourcesTagResource> tagResource;

        public static ListTagResourcesResponseTagResources build(java.util.Map<String, ?> map) throws Exception {
            ListTagResourcesResponseTagResources self = new ListTagResourcesResponseTagResources();
            return TeaModel.build(map, self);
        }

    }

}

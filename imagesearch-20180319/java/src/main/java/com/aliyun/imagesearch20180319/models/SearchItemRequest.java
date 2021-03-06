// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.imagesearch20180319.models;

import com.aliyun.tea.*;

public class SearchItemRequest extends TeaModel {
    @NameInMap("headers")
    public java.util.Map<String, String> headers;

    @NameInMap("query")
    @Validation(required = true)
    public SearchItemQuery query;

    public static SearchItemRequest build(java.util.Map<String, ?> map) throws Exception {
        SearchItemRequest self = new SearchItemRequest();
        return TeaModel.build(map, self);
    }

    public SearchItemRequest setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public SearchItemRequest setQuery(SearchItemQuery query) {
        this.query = query;
        return this;
    }
    public SearchItemQuery getQuery() {
        return this.query;
    }

}

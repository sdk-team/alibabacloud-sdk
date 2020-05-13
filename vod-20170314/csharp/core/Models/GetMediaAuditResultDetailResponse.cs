// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.vod20170314.Models
{
    public class GetMediaAuditResultDetailResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("MediaAuditResultDetail")]
        [Validation(Required=true)]
        public GetMediaAuditResultDetailResponseMediaAuditResultDetail MediaAuditResultDetail { get; set; }
        public class GetMediaAuditResultDetailResponseMediaAuditResultDetail : TeaModel {
            [NameInMap("Total")]
            [Validation(Required=true)]
            public int? Total { get; set; }
            [NameInMap("List")]
            [Validation(Required=true)]
            public List<GetMediaAuditResultDetailResponseMediaAuditResultDetailList> List { get; set; }
            public class GetMediaAuditResultDetailResponseMediaAuditResultDetailList : TeaModel {
                    public string PornLabel { get; set; }
                    public string PornScore { get; set; }
                    public string TerrorismLabel { get; set; }
                    public string TerrorismScore { get; set; }
                    public string Timestamp { get; set; }
                    public string Url { get; set; }
            }
        };

    }

}

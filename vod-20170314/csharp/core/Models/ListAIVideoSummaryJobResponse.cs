// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.vod20170314.Models
{
    public class ListAIVideoSummaryJobResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("AIVideoSummaryJobList")]
        [Validation(Required=true)]
        public ListAIVideoSummaryJobResponseAIVideoSummaryJobList AIVideoSummaryJobList { get; set; }
        public class ListAIVideoSummaryJobResponseAIVideoSummaryJobList : TeaModel {
            [NameInMap("AIVideoSummaryJob")]
            [Validation(Required=true)]
            public List<ListAIVideoSummaryJobResponseAIVideoSummaryJobListAIVideoSummaryJob> AIVideoSummaryJob { get; set; }
            public class ListAIVideoSummaryJobResponseAIVideoSummaryJobListAIVideoSummaryJob : TeaModel {
                    public string JobId { get; set; }
                    public string MediaId { get; set; }
                    public string Status { get; set; }
                    public string Code { get; set; }
                    public string Message { get; set; }
                    public string CreationTime { get; set; }
                    public string Data { get; set; }
            }
        };

        [NameInMap("NonExistAIVideoSummaryJobIds")]
        [Validation(Required=true)]
        public ListAIVideoSummaryJobResponseNonExistAIVideoSummaryJobIds NonExistAIVideoSummaryJobIds { get; set; }
        public class ListAIVideoSummaryJobResponseNonExistAIVideoSummaryJobIds : TeaModel {
            [NameInMap("String")]
            [Validation(Required=true)]
            public List<string> String { get; set; }
        };

    }

}
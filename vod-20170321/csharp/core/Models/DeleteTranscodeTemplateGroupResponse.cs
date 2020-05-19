// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.vod20170321.Models
{
    public class DeleteTranscodeTemplateGroupResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("NonExistTranscodeTemplateIds")]
        [Validation(Required=true)]
        public List<string> NonExistTranscodeTemplateIds { get; set; }

    }

}

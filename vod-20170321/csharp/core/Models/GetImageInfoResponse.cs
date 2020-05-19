// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.vod20170321.Models
{
    public class GetImageInfoResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("ImageInfo")]
        [Validation(Required=true)]
        public GetImageInfoResponseImageInfo ImageInfo { get; set; }
        public class GetImageInfoResponseImageInfo : TeaModel {
            [NameInMap("ImageId")]
            [Validation(Required=true)]
            public string ImageId { get; set; }
            [NameInMap("Title")]
            [Validation(Required=true)]
            public string Title { get; set; }
            [NameInMap("CreationTime")]
            [Validation(Required=true)]
            public string CreationTime { get; set; }
            [NameInMap("ImageType")]
            [Validation(Required=true)]
            public string ImageType { get; set; }
            [NameInMap("Tags")]
            [Validation(Required=true)]
            public string Tags { get; set; }
            [NameInMap("URL")]
            [Validation(Required=true)]
            public string URL { get; set; }
            [NameInMap("CateId")]
            [Validation(Required=true)]
            public long CateId { get; set; }
            [NameInMap("CateName")]
            [Validation(Required=true)]
            public string CateName { get; set; }
            [NameInMap("Description")]
            [Validation(Required=true)]
            public string Description { get; set; }
            [NameInMap("StorageLocation")]
            [Validation(Required=true)]
            public string StorageLocation { get; set; }
            [NameInMap("Status")]
            [Validation(Required=true)]
            public string Status { get; set; }
            [NameInMap("AppId")]
            [Validation(Required=true)]
            public string AppId { get; set; }
            [NameInMap("RegionId")]
            [Validation(Required=true)]
            public string RegionId { get; set; }
            [NameInMap("AuditStatus")]
            [Validation(Required=true)]
            public string AuditStatus { get; set; }
            [NameInMap("Mezzanine")]
            [Validation(Required=true)]
            public GetImageInfoResponseImageInfoMezzanine Mezzanine { get; set; }
            public class GetImageInfoResponseImageInfoMezzanine : TeaModel {
                [NameInMap("OriginalFileName")]
                [Validation(Required=true)]
                public string OriginalFileName { get; set; }

                [NameInMap("FileSize")]
                [Validation(Required=true)]
                public string FileSize { get; set; }

                [NameInMap("Width")]
                [Validation(Required=true)]
                public int? Width { get; set; }

                [NameInMap("Height")]
                [Validation(Required=true)]
                public int? Height { get; set; }

                [NameInMap("FileURL")]
                [Validation(Required=true)]
                public string FileURL { get; set; }

            }
        };

    }

}

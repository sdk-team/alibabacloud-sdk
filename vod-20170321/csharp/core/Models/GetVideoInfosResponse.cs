// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.vod20170321.Models
{
    public class GetVideoInfosResponse : TeaModel {
        [NameInMap("RequestId")]
        [Validation(Required=true)]
        public string RequestId { get; set; }

        [NameInMap("VideoList")]
        [Validation(Required=true)]
        public List<GetVideoInfosResponseVideoList> VideoList { get; set; }
        public class GetVideoInfosResponseVideoList : TeaModel {
            [NameInMap("VideoId")]
            [Validation(Required=true)]
            public string VideoId { get; set; }

            [NameInMap("Title")]
            [Validation(Required=true)]
            public string Title { get; set; }

            [NameInMap("Tags")]
            [Validation(Required=true)]
            public string Tags { get; set; }

            [NameInMap("Status")]
            [Validation(Required=true)]
            public string Status { get; set; }

            [NameInMap("Size")]
            [Validation(Required=true)]
            public long Size { get; set; }

            [NameInMap("Duration")]
            [Validation(Required=true)]
            public float? Duration { get; set; }

            [NameInMap("Description")]
            [Validation(Required=true)]
            public string Description { get; set; }

            [NameInMap("ModificationTime")]
            [Validation(Required=true)]
            public string ModificationTime { get; set; }

            [NameInMap("CreationTime")]
            [Validation(Required=true)]
            public string CreationTime { get; set; }

            [NameInMap("CoverURL")]
            [Validation(Required=true)]
            public string CoverURL { get; set; }

            [NameInMap("CateId")]
            [Validation(Required=true)]
            public long CateId { get; set; }

            [NameInMap("CateName")]
            [Validation(Required=true)]
            public string CateName { get; set; }

            [NameInMap("DownloadSwitch")]
            [Validation(Required=true)]
            public string DownloadSwitch { get; set; }

            [NameInMap("TemplateGroupId")]
            [Validation(Required=true)]
            public string TemplateGroupId { get; set; }

            [NameInMap("PreprocessStatus")]
            [Validation(Required=true)]
            public string PreprocessStatus { get; set; }

            [NameInMap("StorageLocation")]
            [Validation(Required=true)]
            public string StorageLocation { get; set; }

            [NameInMap("RegionId")]
            [Validation(Required=true)]
            public string RegionId { get; set; }

            [NameInMap("CustomMediaInfo")]
            [Validation(Required=true)]
            public string CustomMediaInfo { get; set; }

            [NameInMap("AppId")]
            [Validation(Required=true)]
            public string AppId { get; set; }

            [NameInMap("ThumbnailList")]
            [Validation(Required=true)]
            public List<GetVideoInfosResponseVideoListThumbnailList> ThumbnailList { get; set; }
            public class GetVideoInfosResponseVideoListThumbnailList : TeaModel {
                [NameInMap("URL")]
                [Validation(Required=true)]
                public string URL { get; set; }

            }

            [NameInMap("Snapshots")]
            [Validation(Required=true)]
            public List<string> Snapshots { get; set; }

        }

        [NameInMap("NonExistVideoIds")]
        [Validation(Required=true)]
        public List<string> NonExistVideoIds { get; set; }

    }

}

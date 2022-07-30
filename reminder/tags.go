package reminder

import "fmt"

type TagInfo struct {
	Tag     string
	TagName string
}

type TagInfos struct {
	Tags []*TagInfo
	Code int64
}

func GetTags() {
	infos := TagInfos{
		Tags: []*TagInfo{
			{Tag: "a", TagName: "apple"},
			{Tag: "b", TagName: "banana"},
		},
		Code: 64,
	}

	tags := infos.Tags

	tagInfos := make([]*TagInfo, 0, len(tags))
	for _, tag := range tags {
		tagInfos = append(tagInfos, &TagInfo{
			Tag:     tag.Tag,
			TagName: tag.TagName,
		})
	}
	res := &TagInfos{
		Tags: tagInfos,
	}

	for _, v := range res.Tags {
		fmt.Println(v.Tag, v.TagName)
	}
}

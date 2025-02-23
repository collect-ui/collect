package collect

func GetFilters() map[string]any {
	type register map[string]any
	return register{
		"uuid":                Uuid,
		"is_empty":            IsEmpty,
		"must":                Must,
		"current_date_time":   CurrentDateTime,
		"current_date_format": CurrentDateFormat,
		"replace":             Replace,
		"md5":                 Md5,
		"sub_str":             SubStr,
		"get_key":             GetKey,
		"pinyin":              Pinyin,
		"hash_sha":            HashSha,
		"snow_id":             SnowID,
		"index":               Index,
		"unix_time":           UnixTime,
		"unix_time2datetime":  UnixTime2Datetime,
		"contains":            Contains,
		"to_json":             ToJSON,
		"cast":                Cast,
		"multiply":            Multiply,
		"divide":              Divide,
		"sub_arr":             SubArr,
		"range_number":        RangeNumber,
		"sub_arr_attr":        SubArrAttr,
		"str_contains":        StrContains,
		"random_int":          RandomInt,
		"first_item":          FirstItem,
		"concat":              Concat,
		"genId":               GenId,
		"join":                Join,
		"date_format":         DateFormat,
	}

}

package e

var MsgFlags = map[int]string{
	Success:            "成功",
	InternalError:      "内部错误",
	InvalidParams:      "参数错误",
	Unauthorized:       "认证失败",
	NotFound:           "未找到指定的数据",
	JWTNotAToken:       "令牌非法",
	JWTInvalid:         "令牌无效",
	JWTOutOfTime:       "超出认证时间",
	PermissionDenied:   "无权限",
	RegistrationFailed: "注册失败",
	RegisterDuplicated: "该用户名已被注册",
	NoMoreLocker:       "该储物柜已满",
	UploadFailed:       "上传失败",
	FileTypeMismatch:   "文件类型不符",
	RedisError:         "缓存错误",
	JsonUnmarshalError: "Json解析错误",
	CacheNotFound:      "缓存不存在",
	UsernameInvalid:    "用户名不能以数字开头",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InternalError]
}

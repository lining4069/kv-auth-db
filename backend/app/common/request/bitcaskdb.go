package request

type BitcaskPutRequest struct {
	Key   string `form:"key" json:"Key" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
}

func (b *BitcaskPutRequest) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"key.required":   "key值不能为空",
		"value.required": "value值不能为空",
	}
}

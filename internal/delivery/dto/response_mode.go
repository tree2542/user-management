package dto

type ResponseModel struct {
	StatusCode int         `json:"status_code"`
	MessageTh  string      `json:"message_th"`
	MessageEn  string      `json:"message_en"`
	Data       interface{} `json:"data"`
}

func (r *ResponseModel) SetSuccess(data interface{}) {
	r.StatusCode = 200
	r.MessageTh = "สำเร็จ"
	r.MessageEn = "Success"
	r.Data = data
}

func (r *ResponseModel) SetInternalServerError(data interface{}) {
	r.StatusCode = 500
	r.MessageTh = "เซิฟเวอร์มีปัญหา"
	r.MessageEn = "InternalServerError"
	r.Data = data
}


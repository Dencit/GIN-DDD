package api_trans

type ApiTransStruct struct {
}

func Instance() *ApiTransStruct {
	instance := &ApiTransStruct{}
	return instance
}

func (receiver ApiTransStruct) Index(obj any, meta any, trans any) any {

	return nil
}

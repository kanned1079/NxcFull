package handler

import "errors"

func failOnRpcError(err error, response interface{}) error {
	if err != nil {
		return errors.New("rpc调用失败" + err.Error())
	}
	if response == nil {
		return errors.New("rpc调用失败无返回值")
	}
	return nil
}

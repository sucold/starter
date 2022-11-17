package mail

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

func CreateClient1(accessKeyId *string, accessKeySecret *string) (_result *dm20151123.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dm.aliyuncs.com")
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:       tea.String("as"),
		AddressType:       tea.Int32(1),
		TagName:           tea.String("asd"),
		ReplyToAddress:    tea.Bool(true),
		ToAddress:         tea.String("asdasd"),
		Subject:           tea.String("sadasd"),
		HtmlBody:          tea.String("asdasd"),
		TextBody:          tea.String("csacsac"),
		FromAlias:         tea.String("ascasc"),
		ReplyAddress:      tea.String("ascasc"),
		ReplyAddressAlias: tea.String("ascasc"),
		ClickTrace:        tea.String("sacasc"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SingleSendMailWithOptions(singleSendMailRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		if _err != nil {
			return _err
		}
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}

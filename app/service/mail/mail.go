package mail

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
)

type Config struct {
	Endpoint          string
	AccessKeyId       string
	AccessKeySecret   string
	AccountName       string
	AddressType       int32
	ReplyToAddress    bool
	FromAlias         string
	ReplyAddress      string
	ReplyAddressAlias string
	Expire            int64
	ErrorTimes        int64
	DiffTime          int64
}

var (
	Conf    *Config
	options = &util.RuntimeOptions{}
	client  *dm20151123.Client
)

func Init(config *Config) (err error) {
	Conf = config
	if client, err = CreateClient(tea.String(Conf.AccessKeyId), tea.String(Conf.AccessKeySecret)); err != nil {
		return err
	}
	options = &util.RuntimeOptions{}
	return nil
}
func CreateClient(accessKeyId *string, accessKeySecret *string) (result *dm20151123.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String(Conf.Endpoint)
	result = &dm20151123.Client{}
	result, err = dm20151123.NewClient(config)
	return result, err
}

func Send(ctx context.Context, to string, template *Template, params ...map[string]any) (err error) {
	return nil
	var content string
	var subject string
	if content, err = g.View().ParseContent(ctx, template.Body, params...); err != nil {
		return err
	}
	if subject, err = g.View().ParseContent(ctx, template.Subject, params...); err != nil {
		return err
	}
	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:       tea.String("admin@goant.xyz"),
		AddressType:       tea.Int32(Conf.AddressType),
		ReplyToAddress:    tea.Bool(Conf.ReplyToAddress),
		ToAddress:         tea.String(to),
		Subject:           tea.String(subject),
		HtmlBody:          tea.String(content),
		FromAlias:         tea.String(Conf.FromAlias),
		ReplyAddress:      tea.String(Conf.ReplyAddress),
		ReplyAddressAlias: tea.String(Conf.ReplyAddressAlias),
	}
	if _, err := client.SingleSendMailWithOptions(singleSendMailRequest, options); err != nil {
		return err
	}
	return nil
}

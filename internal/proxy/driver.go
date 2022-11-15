package proxy

import (
	"github.com/litsoftware/litmedia/internal/gateway/core"
	"github.com/litsoftware/litmedia/internal/gateway/oss"
	"github.com/litsoftware/litmedia/internal/gateway/s3"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"strings"
)

func GetClient(confName string) interface{} {
	//supportDrivers := []string{
	//	"oss",
	//	"s3",
	//}

	//storageConfs := []string{
	//	"sensitive",
	//	"general",
	//}

	driver := "oss"

	conf := config.GetMap(driver + "." + confName)
	var client core.ObjectLayer
	var err error

	if conf != nil {
		if strings.Contains(conf["endpoint"].(string), "aliyuncs.com") {
			client, err = (&oss.OssLayer{}).NewGatewayLayer(core.Credentials{
				Endpoint:        conf["endpoint"].(string),
				AccessKeyId:     conf["access_key_id"].(string),
				AccessKeySecret: conf["access_key_secret"].(string),
				Ssl:             true,
				Bucket:          conf["bucket"].(string),
			})

			if err != nil {
				return nil
			}
		} else {
			client, err = (&s3.GatewayS3Layer{}).NewGatewayLayer(core.Credentials{
				Endpoint:        conf["endpoint"].(string),
				AccessKeyId:     conf["access_key_id"].(string),
				AccessKeySecret: conf["access_key_secret"].(string),
				Ssl:             true,
				Bucket:          conf["bucket"].(string),
			})
			if err != nil {
				return nil
			}
		}
	}

	return client
}

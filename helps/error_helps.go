package helps

import (
	"github.com/OpenStars/EtcdBackendService/Int64BigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"strings"
)

func IsError(err error) bool {
	return err != nil && !strings.Contains(err.Error(), generic.TErrorCode_EGood.String())
}

package swagger

import (
	"fmt"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
)

func MergeSwaggers(swaggers ...[]byte) ([]byte, error) {
	specs := make([]*spec.Swagger, 0, len(swaggers))

	for _, swagger := range swaggers {
		s := new(spec.Swagger)
		if err := s.UnmarshalJSON(swagger); err != nil {
			return nil, fmt.Errorf(`unmarshal swagger: %w`, err)
		}
		specs = append(specs, s)
	}

	if len(specs) > 1 {
		_ = analysis.Mixin(specs[0], specs[1:]...)
	}

	return specs[0].MarshalJSON()
}

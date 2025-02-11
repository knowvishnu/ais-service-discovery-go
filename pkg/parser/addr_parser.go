package parser

import (
	"github.com/peak-ai/ais-service-discovery-go/pkg/types"
	"strings"
)

// ParseAddr -
// @todo - this is fairly crude, could be improved.
func ParseAddr(signature string) (*types.Signature, error) {
	namespace := ""
	service := ""
	instance := ""
	parts := strings.Split(signature, ".")
	if len(parts) == 2 {
		namespace = parts[0]
		service = parts[1]
	} else {
		namespace = "default"
		service = parts[0]
	}
	svcParts := strings.Split(service, "->")
	service = svcParts[0]
	instance = svcParts[1]
	return &types.Signature{
		Namespace: namespace,
		Service:   service,
		Instance:   instance,
	}, nil
}

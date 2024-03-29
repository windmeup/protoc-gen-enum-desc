package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/windmeup/protoc-gen-enum-desc/internal/module"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(pgs.DebugEnv("DEBUG_PGED"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.EnumdescForLanguage("go")).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}

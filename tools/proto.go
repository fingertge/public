// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 16:02:17
// * Proj: utils
// * Pack: tools
// * File: proto.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package tools

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func GetMessageName(m proto.Message) string {

	return string(protoimpl.X.MessageDescriptorOf(m).FullName())
}

func NewProtoMessage(name string) (proto.Message, error) {
	m, e := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(name))
	if e != nil {
		return nil, nil
	}
	return m.New().Interface(), nil

}

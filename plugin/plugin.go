package plugin

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/kazegusuri/go-proto-anonymizer"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	anon generator.Single
}

func NewPlugin() generator.Plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "goanonymizer"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.anon = p.NewImport("github.com/kazegusuri/go-proto-anonymizer")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}

		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		}
	}
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	typeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (m *`, typeName, `) Anonymize() {`)
	p.In()
	p.P(`if m == nil { return  }`)

	for _, field := range message.Field {
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "m." + fieldName
		rule := p.getExtention(field)

		repeated := field.IsRepeated() && !p.IsMap(field)
		if repeated {
			if p.isMethodNull(rule) {
				p.P(variableName, ` = nil`)
				continue
			}
			p.P(`for i := range `, variableName, `{`)
			variableName = variableName + "[i]"
			p.In()
		}

		if p.isOneofType(field) {
			oneofType := p.OneOfTypeName(message, field)
			oneofName := p.GetFieldName(message, field)

			p.P(`if ov, ok := m`, `.Get`, oneofName+`().(* `+oneofType+`); ok {`)
			variableName = "ov." + fieldName
			p.In()
		}

		p.generateForField(file, message, field, rule, variableName, repeated)

		if p.isOneofType(field) {
			p.Out()
			p.P(`}`)
		}

		if repeated {
			p.Out()
			p.P(`}`)
		}
	}

	p.Out()
	p.P(`}`)
	p.P()
}

func (p *plugin) isOneofType(field *descriptor.FieldDescriptorProto) bool {
	return field.OneofIndex != nil
}

func (p *plugin) getExtention(field *descriptor.FieldDescriptorProto) *anonymizer.AnonymizeRule {
	var rule *anonymizer.AnonymizeRule
	if field.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(field.Options, anonymizer.E_Anon)
	if err != nil {
		return nil
	}
	rule, ok := v.(*anonymizer.AnonymizeRule)
	if !ok {
		return nil
	}
	return rule
}

func (p *plugin) isMethodNull(rule *anonymizer.AnonymizeRule) bool {
	return rule != nil && rule.Method == anonymizer.AnonymizeMethod_NULL
}

func (p *plugin) generateForField(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto, rule *anonymizer.AnonymizeRule, variableName string, repeated bool) {
	// TODO: support well known type to log pretty message
	// TODO: marshal unknown type

	switch field.GetTypeName() {
	case ".google.protobuf.Duration":
		return

	case ".google.protobuf.Timestamp":
		return
	}

	if p.IsMap(field) {
		mapDesc := p.GoMapType(nil, field)
		if mapDesc == nil {
			p.P("// unavaiable map type")
			return
		}
		valField := mapDesc.ValueField

		// sanity check to avoid unexpected loop
		if p.IsMap(valField) {
			p.P("// unavaible map type: nested map")
			return
		}

		if p.isMethodNull(rule) {
			p.P(variableName, ` = nil`)
			return
		}

		p.P("for mk := range ", variableName, " {")
		p.In()
		variableName = variableName + `[mk]`
		p.generateForField(file, message, valField, p.getExtention(valField), variableName, false)

		p.Out()
		p.P("}")
		return
	}

	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = 0`)
		} else {
			p.P(`_ = `, variableName)
		}

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = false`)
		} else {
			p.P(`_ = `, variableName)
		}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = ""`)
		} else {
			p.P(`_ = `, variableName)
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = nil`)
		} else {
			p.P(`_ = `, variableName)
		}

	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = 0`)
		} else {
			p.P(`_ = `, variableName)
		}

	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if p.isMethodNull(rule) {
			p.P(variableName, ` = nil`)
		} else {
			p.P(p.anon.Use(), `.Anonymize(`, variableName, `)`)
		}

	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// ?
		p.P("// group type not supported")
	}
}

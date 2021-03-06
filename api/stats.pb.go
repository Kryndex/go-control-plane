// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/stats.proto

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf4 "github.com/gogo/protobuf/types"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration for pluggable stats sinks.
type StatsSink struct {
	// The name of the stats sink to instantiate. The name must match a supported
	// stats sink. *envoy.statsd* is a built-in sink suitable for emitting to
	// `statsd <https://github.com/etsy/statsd>`_. Any other built-in stats sink
	// can be found in `well_known_names.h
	// <https://github.com/envoyproxy/envoy/blob/master/source/common/config/well_known_names.h>`_
	// in the Envoy repository.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Stats sink specific configuration which depends on the sink being
	// instantiated. See :ref:`StatsdSink <envoy_api_msg_StatsdSink>` for an
	// example.
	Config *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *StatsSink) Reset()                    { *m = StatsSink{} }
func (m *StatsSink) String() string            { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()               {}
func (*StatsSink) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{0} }

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StatsSink) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Statistics :ref:`architecture overview <arch_overview_statistics>`.
type StatsConfig struct {
	// Each stat name is iteratively processed through these tag specifiers.
	// When a tag is matched, the first capture group is removed from the name so
	// later :ref:`TagSpecifiers <envoy_api_msg_TagSpecifier>` cannot match that
	// same portion of the match.
	StatsTags []*TagSpecifier `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags" json:"stats_tags,omitempty"`
	// Use all default tag regexes specified in Envoy. These can be combined with
	// custom tags specified in :ref:`stats_tags
	// <envoy_api_field_StatsConfig.stats_tags>`. They will be processed before
	// the custom tags.
	//
	// .. note::
	//
	//   If any default tags are specified twice, the config will be considered
	//   invalid.
	//
	// See `well_known_names.h
	// <https://github.com/envoyproxy/envoy/blob/master/source/common/config/well_known_names.h>`_
	// for a list of the default tags in Envoy.
	//
	// If not provided, the value is assumed to be true.
	UseAllDefaultTags *google_protobuf.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags" json:"use_all_default_tags,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{1} }

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

// Designates a tag to strip from the tag extracted name and provide as a named
// tag value for all statistics. This will only occur if any part of the name
// matches the regex provided with one or more capture groups.
type TagSpecifier struct {
	// Attaches an identifier to the tag values to identify the tag being in the
	// sink. Envoy has a set of default names and regexes to extract dynamic
	// portions of existing stats, which can be found in `well_known_names.h
	// <https://github.com/envoyproxy/envoy/blob/master/source/common/config/well_known_names.h>`_
	// in the Envoy repository. If a :ref:`tag_name
	// <envoy_api_field_TagSpecifier.tag_name>` is provided in the config with an
	// empty regex, Envoy will attempt to find that name in its set of defaults
	// and use the accompanying regex.
	//
	// .. note::
	//
	//   If any default tags are specified twice, the config will be considered
	//   invalid.
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// The first capture group identifies the portion of the name to remove. The
	// second capture group (which will normally be nested inside the first) will
	// designate the value of the tag for the statistic. If no second capture
	// group is provided, the first will also be used to set the value of the tag.
	// All other capture groups will be ignored.
	//
	// Take for example, with a stat name ``cluster.foo_cluster.upstream_rq_timeout``
	// and
	//
	// .. code-block:: json
	//
	//   {
	//     "tag_name": "envoy.cluster_name",
	//     "regex": "^cluster\.((.+?)\.)"
	//   }
	//
	// Note that the regex will remove ``foo_cluster.`` making the tag extracted
	// name ``cluster.upstream_rq_timeout`` and the tag value for
	// ``envoy.cluster_name`` will be ``foo_cluster`` (note: there will be no
	// ``.`` character because of the second capture group).
	//
	// An example with two regexes and stat name
	// ``http.connection_manager_1.user_agent.ios.downstream_cx_total``:
	//
	// .. code-block:: json
	//
	//   [
	//     {
	//       "tag_name": "envoy.http_user_agent",
	//       "regex": "^http(?=\.).*?\.user_agent\.((.+?)\.)\w+?$"
	//     },
	//     {
	//       "tag_name": "envoy.http_conn_manager_prefix",
	//       "regex": "^http\.((.*?)\.)"
	//     }
	//   ]
	//
	// The first regex will remove ``ios.``, leaving the tag extracted name
	// ``http.connection_manager_1.user_agent.downstream_cx_total``. The tag
	// ``envoy.http_user_agent`` will be added with tag value ``ios``.
	//
	// The second regex will remove ``connection_manager_1.`` from the tag
	// extracted name produced by the first regex
	// ``http.connection_manager_1.user_agent.downstream_cx_total``, leaving
	// ``http.user_agent.downstream_cx_total`` as the tag extracted name. The tag
	// ``envoy.http_conn_manager_prefix`` will be added with the tag value
	// ``connection_manager_1``.
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (m *TagSpecifier) Reset()                    { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string            { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()               {}
func (*TagSpecifier) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{2} }

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func (m *TagSpecifier) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// Stats configuration proto schema for built-in *envoy.statsd* sink.
type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
}

func (m *StatsdSink) Reset()                    { *m = StatsdSink{} }
func (m *StatsdSink) String() string            { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()               {}
func (*StatsdSink) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{3} }

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StatsdSink_Address struct {
	Address *Address `protobuf:"bytes,1,opt,name=address,oneof"`
}
type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier()        {}
func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StatsdSink_OneofMarshaler, _StatsdSink_OneofUnmarshaler, _StatsdSink_OneofSizer, []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

func _StatsdSink_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StatsdSink)
	// statsd_specifier
	switch x := m.StatsdSpecifier.(type) {
	case *StatsdSink_Address:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Address); err != nil {
			return err
		}
	case *StatsdSink_TcpClusterName:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.TcpClusterName)
	case nil:
	default:
		return fmt.Errorf("StatsdSink.StatsdSpecifier has unexpected type %T", x)
	}
	return nil
}

func _StatsdSink_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StatsdSink)
	switch tag {
	case 1: // statsd_specifier.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Address)
		err := b.DecodeMessage(msg)
		m.StatsdSpecifier = &StatsdSink_Address{msg}
		return true, err
	case 2: // statsd_specifier.tcp_cluster_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.StatsdSpecifier = &StatsdSink_TcpClusterName{x}
		return true, err
	default:
		return false, nil
	}
}

func _StatsdSink_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StatsdSink)
	// statsd_specifier
	switch x := m.StatsdSpecifier.(type) {
	case *StatsdSink_Address:
		s := proto.Size(x.Address)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StatsdSink_TcpClusterName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TcpClusterName)))
		n += len(x.TcpClusterName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Stats configuration proto schema for built-in *envoy.dog_statsd* sink.
// The sink emits stats with `DogStatsD <https://docs.datadoghq.com/guides/dogstatsd/>`_
// compatible tags. Tags are configurable via :ref:`StatsConfig <envoy_api_msg_StatsConfig>`.
type DogStatsdSink struct {
	// Types that are valid to be assigned to DogStatsdSpecifier:
	//	*DogStatsdSink_Address
	//	*DogStatsdSink_TcpClusterName
	DogStatsdSpecifier isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
}

func (m *DogStatsdSink) Reset()                    { *m = DogStatsdSink{} }
func (m *DogStatsdSink) String() string            { return proto.CompactTextString(m) }
func (*DogStatsdSink) ProtoMessage()               {}
func (*DogStatsdSink) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{4} }

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DogStatsdSink_Address struct {
	Address *Address `protobuf:"bytes,1,opt,name=address,oneof"`
}
type DogStatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier()        {}
func (*DogStatsdSink_TcpClusterName) isDogStatsdSink_DogStatsdSpecifier() {}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (m *DogStatsdSink) GetAddress() *Address {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *DogStatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DogStatsdSink) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DogStatsdSink_OneofMarshaler, _DogStatsdSink_OneofUnmarshaler, _DogStatsdSink_OneofSizer, []interface{}{
		(*DogStatsdSink_Address)(nil),
		(*DogStatsdSink_TcpClusterName)(nil),
	}
}

func _DogStatsdSink_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DogStatsdSink)
	// dog_statsd_specifier
	switch x := m.DogStatsdSpecifier.(type) {
	case *DogStatsdSink_Address:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Address); err != nil {
			return err
		}
	case *DogStatsdSink_TcpClusterName:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.TcpClusterName)
	case nil:
	default:
		return fmt.Errorf("DogStatsdSink.DogStatsdSpecifier has unexpected type %T", x)
	}
	return nil
}

func _DogStatsdSink_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DogStatsdSink)
	switch tag {
	case 1: // dog_statsd_specifier.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Address)
		err := b.DecodeMessage(msg)
		m.DogStatsdSpecifier = &DogStatsdSink_Address{msg}
		return true, err
	case 2: // dog_statsd_specifier.tcp_cluster_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DogStatsdSpecifier = &DogStatsdSink_TcpClusterName{x}
		return true, err
	default:
		return false, nil
	}
}

func _DogStatsdSink_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DogStatsdSink)
	// dog_statsd_specifier
	switch x := m.DogStatsdSpecifier.(type) {
	case *DogStatsdSink_Address:
		s := proto.Size(x.Address)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DogStatsdSink_TcpClusterName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TcpClusterName)))
		n += len(x.TcpClusterName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*StatsSink)(nil), "envoy.api.v2.StatsSink")
	proto.RegisterType((*StatsConfig)(nil), "envoy.api.v2.StatsConfig")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.api.v2.TagSpecifier")
	proto.RegisterType((*StatsdSink)(nil), "envoy.api.v2.StatsdSink")
	proto.RegisterType((*DogStatsdSink)(nil), "envoy.api.v2.DogStatsdSink")
}
func (m *StatsSink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatsSink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStats(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.Config.Size()))
		n1, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *StatsConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatsConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatsTags) > 0 {
		for _, msg := range m.StatsTags {
			dAtA[i] = 0xa
			i++
			i = encodeVarintStats(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.UseAllDefaultTags != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.UseAllDefaultTags.Size()))
		n2, err := m.UseAllDefaultTags.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *TagSpecifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TagSpecifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TagName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStats(dAtA, i, uint64(len(m.TagName)))
		i += copy(dAtA[i:], m.TagName)
	}
	if len(m.Regex) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStats(dAtA, i, uint64(len(m.Regex)))
		i += copy(dAtA[i:], m.Regex)
	}
	return i, nil
}

func (m *StatsdSink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatsdSink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StatsdSpecifier != nil {
		nn3, err := m.StatsdSpecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *StatsdSink_Address) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.Address.Size()))
		n4, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *StatsdSink_TcpClusterName) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintStats(dAtA, i, uint64(len(m.TcpClusterName)))
	i += copy(dAtA[i:], m.TcpClusterName)
	return i, nil
}
func (m *DogStatsdSink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DogStatsdSink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DogStatsdSpecifier != nil {
		nn5, err := m.DogStatsdSpecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn5
	}
	return i, nil
}

func (m *DogStatsdSink_Address) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.Address.Size()))
		n6, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *DogStatsdSink_TcpClusterName) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintStats(dAtA, i, uint64(len(m.TcpClusterName)))
	i += copy(dAtA[i:], m.TcpClusterName)
	return i, nil
}
func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StatsSink) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func (m *StatsConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.StatsTags) > 0 {
		for _, e := range m.StatsTags {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	if m.UseAllDefaultTags != nil {
		l = m.UseAllDefaultTags.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func (m *TagSpecifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.TagName)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Regex)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func (m *StatsdSink) Size() (n int) {
	var l int
	_ = l
	if m.StatsdSpecifier != nil {
		n += m.StatsdSpecifier.Size()
	}
	return n
}

func (m *StatsdSink_Address) Size() (n int) {
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}
func (m *StatsdSink_TcpClusterName) Size() (n int) {
	var l int
	_ = l
	l = len(m.TcpClusterName)
	n += 1 + l + sovStats(uint64(l))
	return n
}
func (m *DogStatsdSink) Size() (n int) {
	var l int
	_ = l
	if m.DogStatsdSpecifier != nil {
		n += m.DogStatsdSpecifier.Size()
	}
	return n
}

func (m *DogStatsdSink_Address) Size() (n int) {
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}
func (m *DogStatsdSink_TcpClusterName) Size() (n int) {
	var l int
	_ = l
	l = len(m.TcpClusterName)
	n += 1 + l + sovStats(uint64(l))
	return n
}

func sovStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatsSink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsSink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsSink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &google_protobuf4.Struct{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatsConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatsTags = append(m.StatsTags, &TagSpecifier{})
			if err := m.StatsTags[len(m.StatsTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseAllDefaultTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UseAllDefaultTags == nil {
				m.UseAllDefaultTags = &google_protobuf.BoolValue{}
			}
			if err := m.UseAllDefaultTags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TagSpecifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TagSpecifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagSpecifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Regex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatsdSink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsdSink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsdSink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Address{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.StatsdSpecifier = &StatsdSink_Address{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatsdSpecifier = &StatsdSink_TcpClusterName{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DogStatsdSink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DogStatsdSink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DogStatsdSink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Address{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.DogStatsdSpecifier = &DogStatsdSink_Address{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TcpClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DogStatsdSpecifier = &DogStatsdSink_TcpClusterName{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStats(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/stats.proto", fileDescriptorStats) }

var fileDescriptorStats = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x6b, 0x4a, 0x5b, 0x32, 0x09, 0xd0, 0x5a, 0x41, 0x4d, 0x23, 0x88, 0xa2, 0x3d, 0x45,
	0x1c, 0xbc, 0x22, 0x9c, 0x38, 0xa1, 0xa6, 0x3d, 0x54, 0x42, 0x42, 0x68, 0xb7, 0xe2, 0xba, 0x9a,
	0xec, 0x7a, 0xad, 0x15, 0x66, 0x6d, 0xd9, 0xde, 0x00, 0x2f, 0xc0, 0x8d, 0x23, 0xef, 0xc2, 0x91,
	0x23, 0x47, 0x1e, 0x01, 0xe5, 0xc6, 0x5b, 0xa0, 0xb5, 0x77, 0x51, 0x00, 0x71, 0xed, 0xcd, 0xa3,
	0x7f, 0x66, 0xfe, 0xef, 0xf7, 0xc0, 0x7d, 0xd4, 0x55, 0x6c, 0x1d, 0x3a, 0xcb, 0xb4, 0x51, 0x4e,
	0xd1, 0x11, 0xaf, 0x37, 0xea, 0x03, 0x43, 0x5d, 0xb1, 0xcd, 0x72, 0x7a, 0xd2, 0xca, 0x58, 0x14,
	0x86, 0xdb, 0xae, 0x61, 0xfa, 0x50, 0x28, 0x25, 0x24, 0x8f, 0x7d, 0xb5, 0x6e, 0xca, 0xd8, 0x3a,
	0xd3, 0xe4, 0xae, 0x53, 0x67, 0x7f, 0xab, 0xef, 0x0c, 0x6a, 0xcd, 0x4d, 0x3f, 0x7d, 0xba, 0x41,
	0x59, 0x15, 0xe8, 0x78, 0xdc, 0x3f, 0x82, 0x10, 0xbd, 0x82, 0x41, 0xda, 0x62, 0xa4, 0x55, 0xfd,
	0x86, 0x52, 0xb8, 0x5d, 0xe3, 0x5b, 0x3e, 0x21, 0x73, 0xb2, 0x18, 0x24, 0xfe, 0x4d, 0x63, 0x38,
	0xcc, 0x55, 0x5d, 0x56, 0x62, 0x72, 0x6b, 0x4e, 0x16, 0xc3, 0xe5, 0x29, 0x0b, 0x56, 0xac, 0xb7,
	0x62, 0xa9, 0x07, 0x49, 0xba, 0xb6, 0xe8, 0x33, 0x81, 0xa1, 0x5f, 0x79, 0xe1, 0x6b, 0xfa, 0x0c,
	0xc0, 0x07, 0xcd, 0x1c, 0x0a, 0x3b, 0x21, 0xf3, 0xfd, 0xc5, 0x70, 0x39, 0x65, 0xbb, 0x71, 0xd9,
	0x35, 0x8a, 0x54, 0xf3, 0xbc, 0x2a, 0x2b, 0x6e, 0x92, 0x81, 0xef, 0xbe, 0x46, 0x61, 0xe9, 0x0b,
	0x18, 0x37, 0x96, 0x67, 0x28, 0x65, 0x56, 0xf0, 0x12, 0x1b, 0xe9, 0xc2, 0x92, 0x40, 0x32, 0xfd,
	0x87, 0x64, 0xa5, 0x94, 0x7c, 0x8d, 0xb2, 0xe1, 0xc9, 0x49, 0x63, 0xf9, 0xb9, 0x94, 0x97, 0x61,
	0xaa, 0x5d, 0x16, 0x3d, 0x87, 0xd1, 0xae, 0x0f, 0x3d, 0x83, 0x3b, 0x0e, 0x45, 0xb6, 0x13, 0xf8,
	0xc8, 0xa1, 0x78, 0xd9, 0x66, 0x1e, 0xc3, 0x81, 0xe1, 0x82, 0xbf, 0xf7, 0x46, 0x83, 0x24, 0x14,
	0xd1, 0x47, 0x02, 0xe0, 0x83, 0x15, 0xfe, 0xb3, 0x9e, 0xc0, 0x51, 0x77, 0x21, 0x3f, 0x3e, 0x5c,
	0x3e, 0xf8, 0x33, 0xd4, 0x79, 0x10, 0xaf, 0xf6, 0x92, 0xbe, 0x8f, 0x3e, 0x86, 0x63, 0x97, 0xeb,
	0x2c, 0x97, 0x8d, 0x75, 0xdc, 0x04, 0x6b, 0x6f, 0x71, 0xb5, 0x97, 0xdc, 0x73, 0xb9, 0xbe, 0x08,
	0x42, 0xcb, 0xb0, 0x3a, 0x83, 0x63, 0xff, 0x11, 0x45, 0x66, 0x7f, 0x23, 0x1f, 0x7c, 0xf9, 0xf9,
	0x75, 0x9f, 0x44, 0x9f, 0x08, 0xdc, 0xbd, 0x54, 0xe2, 0xe6, 0x58, 0x1e, 0xc1, 0xb8, 0x50, 0x22,
	0xfb, 0x0f, 0xcf, 0x6a, 0xf4, 0x6d, 0x3b, 0x23, 0xdf, 0xb7, 0x33, 0xf2, 0x63, 0x3b, 0x23, 0xeb,
	0x43, 0x7f, 0x8e, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x73, 0x64, 0x76, 0xe3, 0x02,
	0x00, 0x00,
}

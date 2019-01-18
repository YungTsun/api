# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/resource/v1alpha2/resource.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.resource.metadata.v1alpha2 import metadata_pb2 as datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2
from datahub.recommendation.v1alpha2 import recommendation_pb2 as datahub_dot_recommendation_dot_v1alpha2_dot_recommendation__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from datahub.metric.v1alpha2 import metric_pb2 as datahub_dot_metric_dot_v1alpha2_dot_metric__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/resource/v1alpha2/resource.proto',
  package='containersai.datahub.resource.v1alpha2',
  syntax='proto3',
  serialized_options=_b('Z6github.com/containers-ai/api/datahub/resource/v1alpha2'),
  serialized_pb=_b('\n(datahub/resource/v1alpha2/resource.proto\x12&containersai.datahub.resource.v1alpha2\x1a\x31\x64\x61tahub/resource/metadata/v1alpha2/metadata.proto\x1a\x34\x64\x61tahub/recommendation/v1alpha2/recommendation.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a$datahub/metric/v1alpha2/metric.proto\"\xab\x03\n\tContainer\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\\\n\x0elimit_resource\x18\x02 \x03(\x0b\x32\x44.containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry\x12`\n\x10request_resource\x18\x03 \x03(\x0b\x32\x46.containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry\x1a\x66\n\x12LimitResourceEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12?\n\x05value\x18\x02 \x01(\x0b\x32\x30.containersai.datahub.metric.v1alpha2.MetricData:\x02\x38\x01\x1ah\n\x14RequestResourceEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12?\n\x05value\x18\x02 \x01(\x0b\x32\x30.containersai.datahub.metric.v1alpha2.MetricData:\x02\x38\x01\"\xbb\x03\n\x03Pod\x12X\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32?.containersai.datahub.resource.metadata.v1alpha2.NamespacedName\x12\x15\n\rresource_link\x18\x02 \x01(\t\x12\x45\n\ncontainers\x18\x03 \x03(\x0b\x32\x31.containersai.datahub.resource.v1alpha2.Container\x12\x14\n\x0cis_predicted\x18\x04 \x01(\x08\x12O\n\x06scaler\x18\x05 \x01(\x0b\x32?.containersai.datahub.resource.metadata.v1alpha2.NamespacedName\x12\x11\n\tnode_name\x18\x06 \x01(\t\x12.\n\nstart_time\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12R\n\x06policy\x18\x08 \x01(\x0e\x32\x42.containersai.datahub.recommendation.v1alpha2.RecommendationPolicy\"\x14\n\x04Node\x12\x0c\n\x04name\x18\x01 \x01(\tB8Z6github.com/containers-ai/api/datahub/resource/v1alpha2b\x06proto3')
  ,
  dependencies=[datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2.DESCRIPTOR,datahub_dot_recommendation_dot_v1alpha2_dot_recommendation__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,datahub_dot_metric_dot_v1alpha2_dot_metric__pb2.DESCRIPTOR,])




_CONTAINER_LIMITRESOURCEENTRY = _descriptor.Descriptor(
  name='LimitResourceEntry',
  full_name='containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=480,
  serialized_end=582,
)

_CONTAINER_REQUESTRESOURCEENTRY = _descriptor.Descriptor(
  name='RequestResourceEntry',
  full_name='containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=584,
  serialized_end=688,
)

_CONTAINER = _descriptor.Descriptor(
  name='Container',
  full_name='containersai.datahub.resource.v1alpha2.Container',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.resource.v1alpha2.Container.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit_resource', full_name='containersai.datahub.resource.v1alpha2.Container.limit_resource', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='request_resource', full_name='containersai.datahub.resource.v1alpha2.Container.request_resource', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_CONTAINER_LIMITRESOURCEENTRY, _CONTAINER_REQUESTRESOURCEENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=261,
  serialized_end=688,
)


_POD = _descriptor.Descriptor(
  name='Pod',
  full_name='containersai.datahub.resource.v1alpha2.Pod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containersai.datahub.resource.v1alpha2.Pod.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_link', full_name='containersai.datahub.resource.v1alpha2.Pod.resource_link', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='containers', full_name='containersai.datahub.resource.v1alpha2.Pod.containers', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_predicted', full_name='containersai.datahub.resource.v1alpha2.Pod.is_predicted', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scaler', full_name='containersai.datahub.resource.v1alpha2.Pod.scaler', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_name', full_name='containersai.datahub.resource.v1alpha2.Pod.node_name', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.datahub.resource.v1alpha2.Pod.start_time', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='policy', full_name='containersai.datahub.resource.v1alpha2.Pod.policy', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=691,
  serialized_end=1134,
)


_NODE = _descriptor.Descriptor(
  name='Node',
  full_name='containersai.datahub.resource.v1alpha2.Node',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.resource.v1alpha2.Node.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1136,
  serialized_end=1156,
)

_CONTAINER_LIMITRESOURCEENTRY.fields_by_name['value'].message_type = datahub_dot_metric_dot_v1alpha2_dot_metric__pb2._METRICDATA
_CONTAINER_LIMITRESOURCEENTRY.containing_type = _CONTAINER
_CONTAINER_REQUESTRESOURCEENTRY.fields_by_name['value'].message_type = datahub_dot_metric_dot_v1alpha2_dot_metric__pb2._METRICDATA
_CONTAINER_REQUESTRESOURCEENTRY.containing_type = _CONTAINER
_CONTAINER.fields_by_name['limit_resource'].message_type = _CONTAINER_LIMITRESOURCEENTRY
_CONTAINER.fields_by_name['request_resource'].message_type = _CONTAINER_REQUESTRESOURCEENTRY
_POD.fields_by_name['namespaced_name'].message_type = datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2._NAMESPACEDNAME
_POD.fields_by_name['containers'].message_type = _CONTAINER
_POD.fields_by_name['scaler'].message_type = datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2._NAMESPACEDNAME
_POD.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POD.fields_by_name['policy'].enum_type = datahub_dot_recommendation_dot_v1alpha2_dot_recommendation__pb2._RECOMMENDATIONPOLICY
DESCRIPTOR.message_types_by_name['Container'] = _CONTAINER
DESCRIPTOR.message_types_by_name['Pod'] = _POD
DESCRIPTOR.message_types_by_name['Node'] = _NODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Container = _reflection.GeneratedProtocolMessageType('Container', (_message.Message,), dict(

  LimitResourceEntry = _reflection.GeneratedProtocolMessageType('LimitResourceEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINER_LIMITRESOURCEENTRY,
    __module__ = 'datahub.resource.v1alpha2.resource_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.resource.v1alpha2.Container.LimitResourceEntry)
    ))
  ,

  RequestResourceEntry = _reflection.GeneratedProtocolMessageType('RequestResourceEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINER_REQUESTRESOURCEENTRY,
    __module__ = 'datahub.resource.v1alpha2.resource_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.resource.v1alpha2.Container.RequestResourceEntry)
    ))
  ,
  DESCRIPTOR = _CONTAINER,
  __module__ = 'datahub.resource.v1alpha2.resource_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.resource.v1alpha2.Container)
  ))
_sym_db.RegisterMessage(Container)
_sym_db.RegisterMessage(Container.LimitResourceEntry)
_sym_db.RegisterMessage(Container.RequestResourceEntry)

Pod = _reflection.GeneratedProtocolMessageType('Pod', (_message.Message,), dict(
  DESCRIPTOR = _POD,
  __module__ = 'datahub.resource.v1alpha2.resource_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.resource.v1alpha2.Pod)
  ))
_sym_db.RegisterMessage(Pod)

Node = _reflection.GeneratedProtocolMessageType('Node', (_message.Message,), dict(
  DESCRIPTOR = _NODE,
  __module__ = 'datahub.resource.v1alpha2.resource_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.resource.v1alpha2.Node)
  ))
_sym_db.RegisterMessage(Node)


DESCRIPTOR._options = None
_CONTAINER_LIMITRESOURCEENTRY._options = None
_CONTAINER_REQUESTRESOURCEENTRY._options = None
# @@protoc_insertion_point(module_scope)

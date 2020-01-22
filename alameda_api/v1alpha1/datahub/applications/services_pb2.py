# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/applications/services.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.applications import applications_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_applications_dot_applications__pb2
from alameda_api.v1alpha1.datahub.common import queries_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2
from alameda_api.v1alpha1.datahub.schemas import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/applications/services.proto',
  package='containersai.alameda.v1alpha1.datahub.applications',
  syntax='proto3',
  serialized_options=_b('ZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/applications'),
  serialized_pb=_b('\n8alameda_api/v1alpha1/datahub/applications/services.proto\x12\x32\x63ontainersai.alameda.v1alpha1.datahub.applications\x1a<alameda_api/v1alpha1/datahub/applications/applications.proto\x1a\x31\x61lameda_api/v1alpha1/datahub/common/queries.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/schemas/types.proto\x1a\x17google/rpc/status.proto\"w\n\x19\x43reateApplicationsRequest\x12Z\n\x0c\x61pplications\x18\x01 \x03(\x0b\x32\x44.containersai.alameda.v1alpha1.datahub.applications.WriteApplication\"\xc0\x01\n\x17ListApplicationsRequest\x12U\n\x0fquery_condition\x18\x01 \x01(\x0b\x32<.containersai.alameda.v1alpha1.datahub.common.QueryCondition\x12N\n\x0bschema_meta\x18\x02 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\"\x99\x01\n\x18ListApplicationsResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12Y\n\x0c\x61pplications\x18\x02 \x03(\x0b\x32\x43.containersai.alameda.v1alpha1.datahub.applications.ReadApplication\"\xc2\x01\n\x19\x44\x65leteApplicationsRequest\x12U\n\x0fquery_condition\x18\x01 \x01(\x0b\x32<.containersai.alameda.v1alpha1.datahub.common.QueryCondition\x12N\n\x0bschema_meta\x18\x02 \x03(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMetaBHZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/applicationsb\x06proto3')
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_applications_dot_applications__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_CREATEAPPLICATIONSREQUEST = _descriptor.Descriptor(
  name='CreateApplicationsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='applications', full_name='containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest.applications', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=300,
  serialized_end=419,
)


_LISTAPPLICATIONSREQUEST = _descriptor.Descriptor(
  name='ListApplicationsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest.schema_meta', index=1,
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
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=422,
  serialized_end=614,
)


_LISTAPPLICATIONSRESPONSE = _descriptor.Descriptor(
  name='ListApplicationsResponse',
  full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='applications', full_name='containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse.applications', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=617,
  serialized_end=770,
)


_DELETEAPPLICATIONSREQUEST = _descriptor.Descriptor(
  name='DeleteApplicationsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest.schema_meta', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=773,
  serialized_end=967,
)

_CREATEAPPLICATIONSREQUEST.fields_by_name['applications'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_applications_dot_applications__pb2._WRITEAPPLICATION
_LISTAPPLICATIONSREQUEST.fields_by_name['query_condition'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTAPPLICATIONSREQUEST.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_LISTAPPLICATIONSRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTAPPLICATIONSRESPONSE.fields_by_name['applications'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_applications_dot_applications__pb2._READAPPLICATION
_DELETEAPPLICATIONSREQUEST.fields_by_name['query_condition'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_DELETEAPPLICATIONSREQUEST.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
DESCRIPTOR.message_types_by_name['CreateApplicationsRequest'] = _CREATEAPPLICATIONSREQUEST
DESCRIPTOR.message_types_by_name['ListApplicationsRequest'] = _LISTAPPLICATIONSREQUEST
DESCRIPTOR.message_types_by_name['ListApplicationsResponse'] = _LISTAPPLICATIONSRESPONSE
DESCRIPTOR.message_types_by_name['DeleteApplicationsRequest'] = _DELETEAPPLICATIONSREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateApplicationsRequest = _reflection.GeneratedProtocolMessageType('CreateApplicationsRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEAPPLICATIONSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.CreateApplicationsRequest)
  })
_sym_db.RegisterMessage(CreateApplicationsRequest)

ListApplicationsRequest = _reflection.GeneratedProtocolMessageType('ListApplicationsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTAPPLICATIONSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.ListApplicationsRequest)
  })
_sym_db.RegisterMessage(ListApplicationsRequest)

ListApplicationsResponse = _reflection.GeneratedProtocolMessageType('ListApplicationsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTAPPLICATIONSRESPONSE,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.ListApplicationsResponse)
  })
_sym_db.RegisterMessage(ListApplicationsResponse)

DeleteApplicationsRequest = _reflection.GeneratedProtocolMessageType('DeleteApplicationsRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEAPPLICATIONSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.DeleteApplicationsRequest)
  })
_sym_db.RegisterMessage(DeleteApplicationsRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

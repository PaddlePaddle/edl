# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: compute_server.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

import common_pb2 as common__pb2

DESCRIPTOR = _descriptor.FileDescriptor(
    name='compute_server.proto',
    package='compute_server',
    syntax='proto3',
    serialized_options=None,
    serialized_pb=b'\n\x14\x63ompute_server.proto\x12\x0e\x63ompute_server\x1a\x0c\x63ommon.proto\"N\n\x10\x46ilesMetaRequest\x12\x13\n\x0b\x64\x61ta_reader\x18\x01 \x01(\t\x12%\n\nfiles_meta\x18\x02 \x01(\x0b\x32\x11.common.FilesMeta2\x92\x01\n\rComputeServer\x12H\n\x12\x44oTasksByFilesMeta\x12 .compute_server.FilesMetaRequest\x1a\x0e.common.RPCRet\"\x00\x12\x37\n\x12SetTaskProgramDesc\x12\x0f.common.Program\x1a\x0e.common.RPCRet\"\x00\x62\x06proto3',
    dependencies=[common__pb2.DESCRIPTOR, ])

_FILESMETAREQUEST = _descriptor.Descriptor(
    name='FilesMetaRequest',
    full_name='compute_server.FilesMetaRequest',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='data_reader',
            full_name='compute_server.FilesMetaRequest.data_reader',
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"".decode('utf-8'),
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='files_meta',
            full_name='compute_server.FilesMetaRequest.files_meta',
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[],
    serialized_start=54,
    serialized_end=132, )

_FILESMETAREQUEST.fields_by_name[
    'files_meta'].message_type = common__pb2._FILESMETA
DESCRIPTOR.message_types_by_name['FilesMetaRequest'] = _FILESMETAREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FilesMetaRequest = _reflection.GeneratedProtocolMessageType(
    'FilesMetaRequest',
    (_message.Message, ),
    {
        'DESCRIPTOR': _FILESMETAREQUEST,
        '__module__': 'compute_server_pb2'
        # @@protoc_insertion_point(class_scope:compute_server.FilesMetaRequest)
    })
_sym_db.RegisterMessage(FilesMetaRequest)

_COMPUTESERVER = _descriptor.ServiceDescriptor(
    name='ComputeServer',
    full_name='compute_server.ComputeServer',
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    serialized_start=135,
    serialized_end=281,
    methods=[
        _descriptor.MethodDescriptor(
            name='DoTasksByFilesMeta',
            full_name='compute_server.ComputeServer.DoTasksByFilesMeta',
            index=0,
            containing_service=None,
            input_type=_FILESMETAREQUEST,
            output_type=common__pb2._RPCRET,
            serialized_options=None, ),
        _descriptor.MethodDescriptor(
            name='SetTaskProgramDesc',
            full_name='compute_server.ComputeServer.SetTaskProgramDesc',
            index=1,
            containing_service=None,
            input_type=common__pb2._PROGRAM,
            output_type=common__pb2._RPCRET,
            serialized_options=None, ),
    ])
_sym_db.RegisterServiceDescriptor(_COMPUTESERVER)

DESCRIPTOR.services_by_name['ComputeServer'] = _COMPUTESERVER

# @@protoc_insertion_point(module_scope)

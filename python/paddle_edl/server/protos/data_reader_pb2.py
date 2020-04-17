# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: data_reader.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

import common_pb2 as common__pb2

DESCRIPTOR = _descriptor.FileDescriptor(
    name='data_reader.proto',
    package='data_reader',
    syntax='proto3',
    serialized_options=None,
    serialized_pb=b'\n\x11\x64\x61ta_reader.proto\x12\x0b\x64\x61ta_reader\x1a\x0c\x63ommon.proto\")\n\x06Record\x12\x11\n\trecord_no\x18\x01 \x01(\x03\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"?\n\x04\x46ile\x12\x11\n\tfile_path\x18\x01 \x01(\t\x12$\n\x07records\x18\x02 \x03(\x0b\x32\x13.data_reader.Record\")\n\x05\x46iles\x12 \n\x05\x66iles\x18\x01 \x03(\x0b\x32\x11.data_reader.File\"N\n\x0c\x44\x61taResponse\x12\x1b\n\x03ret\x18\x01 \x01(\x0b\x32\x0e.common.RPCRet\x12!\n\x05\x66iles\x18\x02 \x01(\x0b\x32\x12.data_reader.Files\"B\n\x10\x44\x61taSourceStatus\x12\x0c\n\x04path\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\x05\x12\x10\n\x08\x65rr_info\x18\x03 \x01(\t\"N\n\x18\x44\x61taSourceStatusResponse\x12\x32\n\x0bstatus_list\x18\x01 \x03(\x0b\x32\x1d.data_reader.DataSourceStatus2\x9b\x01\n\nDataServer\x12\x39\n\x07GetData\x12\x11.common.FilesMeta\x1a\x19.data_reader.DataResponse\"\x00\x12R\n\x13GetDataSourceStatus\x12\x12.common.DataSource\x1a%.data_reader.DataSourceStatusResponse\"\x00\x62\x06proto3',
    dependencies=[common__pb2.DESCRIPTOR, ])

_RECORD = _descriptor.Descriptor(
    name='Record',
    full_name='data_reader.Record',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='record_no',
            full_name='data_reader.Record.record_no',
            index=0,
            number=1,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='data',
            full_name='data_reader.Record.data',
            index=1,
            number=2,
            type=12,
            cpp_type=9,
            label=1,
            has_default_value=False,
            default_value=b"",
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
    serialized_start=48,
    serialized_end=89, )

_FILE = _descriptor.Descriptor(
    name='File',
    full_name='data_reader.File',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='file_path',
            full_name='data_reader.File.file_path',
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
            name='records',
            full_name='data_reader.File.records',
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
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
    serialized_start=91,
    serialized_end=154, )

_FILES = _descriptor.Descriptor(
    name='Files',
    full_name='data_reader.Files',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='files',
            full_name='data_reader.Files.files',
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
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
    serialized_start=156,
    serialized_end=197, )

_DATARESPONSE = _descriptor.Descriptor(
    name='DataResponse',
    full_name='data_reader.DataResponse',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='ret',
            full_name='data_reader.DataResponse.ret',
            index=0,
            number=1,
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
        _descriptor.FieldDescriptor(
            name='files',
            full_name='data_reader.DataResponse.files',
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
    serialized_start=199,
    serialized_end=277, )

_DATASOURCESTATUS = _descriptor.Descriptor(
    name='DataSourceStatus',
    full_name='data_reader.DataSourceStatus',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='path',
            full_name='data_reader.DataSourceStatus.path',
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
            name='status',
            full_name='data_reader.DataSourceStatus.status',
            index=1,
            number=2,
            type=5,
            cpp_type=1,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='err_info',
            full_name='data_reader.DataSourceStatus.err_info',
            index=2,
            number=3,
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
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[],
    serialized_start=279,
    serialized_end=345, )

_DATASOURCESTATUSRESPONSE = _descriptor.Descriptor(
    name='DataSourceStatusResponse',
    full_name='data_reader.DataSourceStatusResponse',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='status_list',
            full_name='data_reader.DataSourceStatusResponse.status_list',
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=3,
            has_default_value=False,
            default_value=[],
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
    serialized_start=347,
    serialized_end=425, )

_FILE.fields_by_name['records'].message_type = _RECORD
_FILES.fields_by_name['files'].message_type = _FILE
_DATARESPONSE.fields_by_name['ret'].message_type = common__pb2._RPCRET
_DATARESPONSE.fields_by_name['files'].message_type = _FILES
_DATASOURCESTATUSRESPONSE.fields_by_name[
    'status_list'].message_type = _DATASOURCESTATUS
DESCRIPTOR.message_types_by_name['Record'] = _RECORD
DESCRIPTOR.message_types_by_name['File'] = _FILE
DESCRIPTOR.message_types_by_name['Files'] = _FILES
DESCRIPTOR.message_types_by_name['DataResponse'] = _DATARESPONSE
DESCRIPTOR.message_types_by_name['DataSourceStatus'] = _DATASOURCESTATUS
DESCRIPTOR.message_types_by_name[
    'DataSourceStatusResponse'] = _DATASOURCESTATUSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Record = _reflection.GeneratedProtocolMessageType(
    'Record',
    (_message.Message, ),
    {
        'DESCRIPTOR': _RECORD,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.Record)
    })
_sym_db.RegisterMessage(Record)

File = _reflection.GeneratedProtocolMessageType(
    'File',
    (_message.Message, ),
    {
        'DESCRIPTOR': _FILE,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.File)
    })
_sym_db.RegisterMessage(File)

Files = _reflection.GeneratedProtocolMessageType(
    'Files',
    (_message.Message, ),
    {
        'DESCRIPTOR': _FILES,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.Files)
    })
_sym_db.RegisterMessage(Files)

DataResponse = _reflection.GeneratedProtocolMessageType(
    'DataResponse',
    (_message.Message, ),
    {
        'DESCRIPTOR': _DATARESPONSE,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.DataResponse)
    })
_sym_db.RegisterMessage(DataResponse)

DataSourceStatus = _reflection.GeneratedProtocolMessageType(
    'DataSourceStatus',
    (_message.Message, ),
    {
        'DESCRIPTOR': _DATASOURCESTATUS,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.DataSourceStatus)
    })
_sym_db.RegisterMessage(DataSourceStatus)

DataSourceStatusResponse = _reflection.GeneratedProtocolMessageType(
    'DataSourceStatusResponse',
    (_message.Message, ),
    {
        'DESCRIPTOR': _DATASOURCESTATUSRESPONSE,
        '__module__': 'data_reader_pb2'
        # @@protoc_insertion_point(class_scope:data_reader.DataSourceStatusResponse)
    })
_sym_db.RegisterMessage(DataSourceStatusResponse)

_DATASERVER = _descriptor.ServiceDescriptor(
    name='DataServer',
    full_name='data_reader.DataServer',
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    serialized_start=428,
    serialized_end=583,
    methods=[
        _descriptor.MethodDescriptor(
            name='GetData',
            full_name='data_reader.DataServer.GetData',
            index=0,
            containing_service=None,
            input_type=common__pb2._FILESMETA,
            output_type=_DATARESPONSE,
            serialized_options=None, ),
        _descriptor.MethodDescriptor(
            name='GetDataSourceStatus',
            full_name='data_reader.DataServer.GetDataSourceStatus',
            index=1,
            containing_service=None,
            input_type=common__pb2._DATASOURCE,
            output_type=_DATASOURCESTATUSRESPONSE,
            serialized_options=None, ),
    ])
_sym_db.RegisterServiceDescriptor(_DATASERVER)

DESCRIPTOR.services_by_name['DataServer'] = _DATASERVER

# @@protoc_insertion_point(module_scope)

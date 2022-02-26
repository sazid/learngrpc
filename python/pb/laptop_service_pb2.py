# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: laptop_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import laptop_message_pb2 as laptop__message__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14laptop_service.proto\x12\x03rpc\x1a\x14laptop_message.proto\"2\n\x13\x43reateLaptopRequest\x12\x1b\n\x06laptop\x18\x01 \x01(\x0b\x32\x0b.rpc.Laptop\"\"\n\x14\x43reateLaptopResponse\x12\n\n\x02id\x18\x01 \x01(\t2P\n\rLaptopService\x12?\n\x06\x43reate\x12\x18.rpc.CreateLaptopRequest\x1a\x19.rpc.CreateLaptopResponse\"\x00\x42\x08Z\x06\x61pi/v1b\x06proto3')



_CREATELAPTOPREQUEST = DESCRIPTOR.message_types_by_name['CreateLaptopRequest']
_CREATELAPTOPRESPONSE = DESCRIPTOR.message_types_by_name['CreateLaptopResponse']
CreateLaptopRequest = _reflection.GeneratedProtocolMessageType('CreateLaptopRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATELAPTOPREQUEST,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.CreateLaptopRequest)
  })
_sym_db.RegisterMessage(CreateLaptopRequest)

CreateLaptopResponse = _reflection.GeneratedProtocolMessageType('CreateLaptopResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATELAPTOPRESPONSE,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.CreateLaptopResponse)
  })
_sym_db.RegisterMessage(CreateLaptopResponse)

_LAPTOPSERVICE = DESCRIPTOR.services_by_name['LaptopService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006api/v1'
  _CREATELAPTOPREQUEST._serialized_start=51
  _CREATELAPTOPREQUEST._serialized_end=101
  _CREATELAPTOPRESPONSE._serialized_start=103
  _CREATELAPTOPRESPONSE._serialized_end=137
  _LAPTOPSERVICE._serialized_start=139
  _LAPTOPSERVICE._serialized_end=219
# @@protoc_insertion_point(module_scope)

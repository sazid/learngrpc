# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: keyboard_message.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16keyboard_message.proto\x12\x03rpc\"|\n\x08Keyboard\x12$\n\x06layout\x18\x01 \x01(\x0e\x32\x14.rpc.Keyboard.Layout\x12\x0f\n\x07\x62\x61\x63klit\x18\x02 \x01(\x08\"9\n\x06Layout\x12\x0b\n\x07UNKNOWN\x10\x00\x12\n\n\x06QWERTY\x10\x01\x12\n\n\x06QWERTZ\x10\x02\x12\n\n\x06\x41ZERTY\x10\x03\x42\x08Z\x06\x61pi/v1b\x06proto3')



_KEYBOARD = DESCRIPTOR.message_types_by_name['Keyboard']
_KEYBOARD_LAYOUT = _KEYBOARD.enum_types_by_name['Layout']
Keyboard = _reflection.GeneratedProtocolMessageType('Keyboard', (_message.Message,), {
  'DESCRIPTOR' : _KEYBOARD,
  '__module__' : 'keyboard_message_pb2'
  # @@protoc_insertion_point(class_scope:rpc.Keyboard)
  })
_sym_db.RegisterMessage(Keyboard)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006api/v1'
  _KEYBOARD._serialized_start=31
  _KEYBOARD._serialized_end=155
  _KEYBOARD_LAYOUT._serialized_start=98
  _KEYBOARD_LAYOUT._serialized_end=155
# @@protoc_insertion_point(module_scope)

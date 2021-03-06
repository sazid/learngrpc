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
import filter_message_pb2 as filter__message__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14laptop_service.proto\x12\x03rpc\x1a\x14laptop_message.proto\x1a\x14\x66ilter_message.proto\"2\n\x13\x43reateLaptopRequest\x12\x1b\n\x06laptop\x18\x01 \x01(\x0b\x32\x0b.rpc.Laptop\"\"\n\x14\x43reateLaptopResponse\x12\n\n\x02id\x18\x01 \x01(\t\"2\n\x13SearchLaptopRequest\x12\x1b\n\x06\x66ilter\x18\x01 \x01(\x0b\x32\x0b.rpc.Filter\"3\n\x14SearchLaptopResponse\x12\x1b\n\x06laptop\x18\x01 \x01(\x0b\x32\x0b.rpc.Laptop\"R\n\x12UploadImageRequest\x12\x1e\n\x04info\x18\x01 \x01(\x0b\x32\x0e.rpc.ImageInfoH\x00\x12\x14\n\nchunk_data\x18\x02 \x01(\x0cH\x00\x42\x06\n\x04\x64\x61ta\"2\n\tImageInfo\x12\x11\n\tlaptop_id\x18\x01 \x01(\t\x12\x12\n\nimage_type\x18\x02 \x01(\t\"/\n\x13UploadImageResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\r\"5\n\x11RateLaptopRequest\x12\x11\n\tlaptop_id\x18\x01 \x01(\t\x12\r\n\x05score\x18\x02 \x01(\x01\"S\n\x12RateLaptopResponse\x12\x11\n\tlaptop_id\x18\x01 \x01(\t\x12\x13\n\x0brated_count\x18\x02 \x01(\r\x12\x15\n\raverage_score\x18\x03 \x01(\x01\x32\xa4\x02\n\rLaptopService\x12?\n\x06\x43reate\x12\x18.rpc.CreateLaptopRequest\x1a\x19.rpc.CreateLaptopResponse\"\x00\x12G\n\x0cSearchLaptop\x12\x18.rpc.SearchLaptopRequest\x1a\x19.rpc.SearchLaptopResponse\"\x00\x30\x01\x12\x44\n\x0bUploadImage\x12\x17.rpc.UploadImageRequest\x1a\x18.rpc.UploadImageResponse\"\x00(\x01\x12\x43\n\nRateLaptop\x12\x16.rpc.RateLaptopRequest\x1a\x17.rpc.RateLaptopResponse\"\x00(\x01\x30\x01\x42\x08Z\x06\x61pi/v1b\x06proto3')



_CREATELAPTOPREQUEST = DESCRIPTOR.message_types_by_name['CreateLaptopRequest']
_CREATELAPTOPRESPONSE = DESCRIPTOR.message_types_by_name['CreateLaptopResponse']
_SEARCHLAPTOPREQUEST = DESCRIPTOR.message_types_by_name['SearchLaptopRequest']
_SEARCHLAPTOPRESPONSE = DESCRIPTOR.message_types_by_name['SearchLaptopResponse']
_UPLOADIMAGEREQUEST = DESCRIPTOR.message_types_by_name['UploadImageRequest']
_IMAGEINFO = DESCRIPTOR.message_types_by_name['ImageInfo']
_UPLOADIMAGERESPONSE = DESCRIPTOR.message_types_by_name['UploadImageResponse']
_RATELAPTOPREQUEST = DESCRIPTOR.message_types_by_name['RateLaptopRequest']
_RATELAPTOPRESPONSE = DESCRIPTOR.message_types_by_name['RateLaptopResponse']
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

SearchLaptopRequest = _reflection.GeneratedProtocolMessageType('SearchLaptopRequest', (_message.Message,), {
  'DESCRIPTOR' : _SEARCHLAPTOPREQUEST,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.SearchLaptopRequest)
  })
_sym_db.RegisterMessage(SearchLaptopRequest)

SearchLaptopResponse = _reflection.GeneratedProtocolMessageType('SearchLaptopResponse', (_message.Message,), {
  'DESCRIPTOR' : _SEARCHLAPTOPRESPONSE,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.SearchLaptopResponse)
  })
_sym_db.RegisterMessage(SearchLaptopResponse)

UploadImageRequest = _reflection.GeneratedProtocolMessageType('UploadImageRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADIMAGEREQUEST,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.UploadImageRequest)
  })
_sym_db.RegisterMessage(UploadImageRequest)

ImageInfo = _reflection.GeneratedProtocolMessageType('ImageInfo', (_message.Message,), {
  'DESCRIPTOR' : _IMAGEINFO,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.ImageInfo)
  })
_sym_db.RegisterMessage(ImageInfo)

UploadImageResponse = _reflection.GeneratedProtocolMessageType('UploadImageResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADIMAGERESPONSE,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.UploadImageResponse)
  })
_sym_db.RegisterMessage(UploadImageResponse)

RateLaptopRequest = _reflection.GeneratedProtocolMessageType('RateLaptopRequest', (_message.Message,), {
  'DESCRIPTOR' : _RATELAPTOPREQUEST,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.RateLaptopRequest)
  })
_sym_db.RegisterMessage(RateLaptopRequest)

RateLaptopResponse = _reflection.GeneratedProtocolMessageType('RateLaptopResponse', (_message.Message,), {
  'DESCRIPTOR' : _RATELAPTOPRESPONSE,
  '__module__' : 'laptop_service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.RateLaptopResponse)
  })
_sym_db.RegisterMessage(RateLaptopResponse)

_LAPTOPSERVICE = DESCRIPTOR.services_by_name['LaptopService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\006api/v1'
  _CREATELAPTOPREQUEST._serialized_start=73
  _CREATELAPTOPREQUEST._serialized_end=123
  _CREATELAPTOPRESPONSE._serialized_start=125
  _CREATELAPTOPRESPONSE._serialized_end=159
  _SEARCHLAPTOPREQUEST._serialized_start=161
  _SEARCHLAPTOPREQUEST._serialized_end=211
  _SEARCHLAPTOPRESPONSE._serialized_start=213
  _SEARCHLAPTOPRESPONSE._serialized_end=264
  _UPLOADIMAGEREQUEST._serialized_start=266
  _UPLOADIMAGEREQUEST._serialized_end=348
  _IMAGEINFO._serialized_start=350
  _IMAGEINFO._serialized_end=400
  _UPLOADIMAGERESPONSE._serialized_start=402
  _UPLOADIMAGERESPONSE._serialized_end=449
  _RATELAPTOPREQUEST._serialized_start=451
  _RATELAPTOPREQUEST._serialized_end=504
  _RATELAPTOPRESPONSE._serialized_start=506
  _RATELAPTOPRESPONSE._serialized_end=589
  _LAPTOPSERVICE._serialized_start=592
  _LAPTOPSERVICE._serialized_end=884
# @@protoc_insertion_point(module_scope)

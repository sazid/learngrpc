# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import auth_service_pb2 as auth__service__pb2


class AuthServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Login = channel.unary_unary(
                '/rpc.AuthService/Login',
                request_serializer=auth__service__pb2.LoginRequest.SerializeToString,
                response_deserializer=auth__service__pb2.LoginResponse.FromString,
                )


class AuthServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Login(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Login': grpc.unary_unary_rpc_method_handler(
                    servicer.Login,
                    request_deserializer=auth__service__pb2.LoginRequest.FromString,
                    response_serializer=auth__service__pb2.LoginResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'rpc.AuthService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AuthService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Login(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/rpc.AuthService/Login',
            auth__service__pb2.LoginRequest.SerializeToString,
            auth__service__pb2.LoginResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

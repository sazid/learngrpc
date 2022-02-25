import traceback
import grpc
import generator

import laptop_service_pb2_grpc as laptop_service_rpc
import laptop_service_pb2 as laptop_service


def main():
    channel = grpc.insecure_channel("localhost:8500")
    stub = laptop_service_rpc.LaptopServiceStub(channel)

    laptop = generator.new_laptop()

    try:
        req = laptop_service.CreateLaptopRequest(
            laptop=laptop,
        )
        res = stub.Create(req)
        print(res)
    except:
        traceback.print_exc()


if __name__ == '__main__':
    main()

import traceback
import grpc
import generator

import laptop_service_pb2_grpc as laptop_service_rpc
import laptop_service_pb2 as laptop_service
import filter_message_pb2
import memory_message_pb2


def create_laptop(stub: laptop_service_rpc.LaptopServiceStub):
    laptop = generator.new_laptop()
    req = laptop_service.CreateLaptopRequest(
        laptop=laptop,
    )
    res = stub.Create(req)
    print("created laptop with id:", res.id)


def search_laptop(stub: laptop_service_rpc.LaptopServiceStub, filter: filter_message_pb2.Filter):
    print("search filter:", filter)

    req = laptop_service.SearchLaptopRequest(
        filter=filter,
    )
    stream = stub.SearchLaptop(req)

    for res in stream:
        laptop = res.laptop
        print(f"""found {laptop.id}:
+ RAM: {laptop.ram}
+ CPU cores: {laptop.cpu.number_cores}
+ CPU freq: {laptop.cpu.min_ghz} - {laptop.cpu.max_ghz}
+ Price: {laptop.price_usd}""")


def main():
    #with open("/home/szxo3/Downloads/roots.pem", "rb") as f:
    #creds = grpc.ssl_channel_credentials()
    #channel = grpc.secure_channel("qa.automationsolutionz.com:20001", creds)
    channel = grpc.insecure_channel("qa.automationsolutionz.com:20001")
    stub = laptop_service_rpc.LaptopServiceStub(channel)

    try:
        for i in range(10):
            create_laptop(stub)
    
        search_laptop(stub, filter_message_pb2.Filter(
            max_price_usd=1500,
            min_cpu_cores=8,
            min_cpu_ghz=2.2,
            min_ram=memory_message_pb2.Memory(
                value=8,
                unit=memory_message_pb2.Memory.GIGABYTE,
            ),
        ))
    except:
        traceback.print_exc()


if __name__ == '__main__':
    main()

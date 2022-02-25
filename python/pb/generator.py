from google.protobuf.timestamp_pb2 import Timestamp

import laptop_message_pb2 as laptop
import keyboard_message_pb2 as keyboard
import processor_message_pb2 as processor
import memory_message_pb2 as memory
import storage_message_pb2 as storage
import screen_message_pb2 as screen


import uuid


def new_keyboard() -> keyboard.Keyboard:
    return keyboard.Keyboard(
        layout=keyboard.Keyboard.QWERTY,
        backlit=True,
    )


def new_memory(value, unit) -> memory.Memory:
    return memory.Memory(
        value=value,
        unit=unit,
    )


def new_gpu() -> processor.GPU:
    return processor.GPU(
        brand="Nvidia",
        name="RTX 2060",
        min_ghz=1.0,
        max_ghz=2.0,
        memory=new_memory(6, memory.Memory.GIGABYTE),
    )


def new_cpu() -> processor.CPU:
    return processor.CPU(
        brand="AMD",
        name="Ryzen 5 1600",
        min_ghz=2.2,
        max_ghz=3.5,
        number_cores=6,
        number_threads=12,
    )


def new_storage() -> storage.Storage:
    return storage.Storage(
        driver=storage.Storage.SSD,
        memory=new_memory(512, memory.Memory.GIGABYTE)
    )


def new_screen() -> screen.Screen:
    return screen.Screen(
        size_inch=24.5,
        resolution=screen.Screen.Resolution(
            width=24,
            height=12,
        ),
        panel=screen.Screen.IPS,
        multitouch=True,
    )


def new_laptop() -> laptop.Laptop:
    timestamp = Timestamp()
    timestamp.GetCurrentTime()

    return laptop.Laptop(
        id=str(uuid.uuid4()),
        brand="Generic Brand",
        name="Generic Laptop",
        cpu=new_cpu(),
        gpus=[new_gpu()],
        screen=new_screen(),
        storages=[new_storage()],
        weight_kg=2.5,
        price_usd=600,
        release_year=2022,
        updated_at=timestamp,
    )

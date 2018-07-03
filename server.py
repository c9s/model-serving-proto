from concurrent import futures
import time
import math
import click
import json
import base64
import logging
import scipy.misc
import numpy as np
import io

import grpc
import serving_pb2
import serving_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

def load_image_ndarray(data, mode = 'RGB'):
    f = io.BytesIO(data)
    return scipy.misc.imread(f, mode=mode)

def load_base64_image_ndarray(data, mode = 'RGB'):
    imgbs = base64.b64decode(data)
    f = io.BytesIO(imgbs)
    return scipy.misc.imread(f, mode=mode)


class ObjectDetectionServicer(serving_pb2_grpc.ObjectDetectionServicer):
    def __init__(self):
        pass
    def Detect(self, request, context):
        img = load_image_ndarray(request.image)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    serving_pb2_grpc.add_ObjectDetectionServicer_to_server(ObjectDetectionServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()

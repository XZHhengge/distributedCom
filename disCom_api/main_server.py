import tornado.ioloop
import tornado.web
# from tornado import ioloop
import asyncio
import logging
from proto import code_pb2_grpc
from proto import code_pb2
import grpc


class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")


def make_app():
    return tornado.web.Application([
        (r"/", MainHandler),
    ])


# class CodeServicer(code_pb2_grpc.CodeServicer):
#     def GetAns(self, request, context):
async def run() -> None:
    async with grpc.aio.insecure_channel("localhost:50051") as channel:
        stub = code_pb2_grpc.CodeStub(channel)
        response = await stub.GetAns(code_pb2.CalRequest(number=b"1000000"))
    print(response)


async def main():
    app = make_app()
    app.listen(8888)
    await asyncio.Event().wait()

if __name__ == "__main__":
    logging.basicConfig()
    # asyncio.run(run())
    # asyncio.run(main())
    # app = make_app()
    # app.listen(8888)
    io_loop = tornado.ioloop.IOLoop.current()
    io_loop.run_sync(run)
    io_loop.run_sync(main)

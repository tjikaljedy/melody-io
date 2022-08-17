import asyncio
import logging
import sys
from response_channel import LoggingSubscriber, resp_channel

from rsocket.extensions.authentication import Authentication, AuthenticationSimple
from rsocket.extensions.composite_metadata import CompositeMetadata
from rsocket.helpers import create_future
from rsocket.payload import Payload
from rsocket.routing.request_router import RequestRouter
from rsocket.routing.routing_request_handler import RoutingRequestHandler
from rsocket.rsocket_server import RSocketServer
from rsocket.transports.tcp import TransportTCP

router = RequestRouter()

@router.channel('channel')
async def channel_response(payload, composite_metadata):
    logging.info('Got channel request')
    subscriber = LoggingSubscriber()
    channel = resp_channel(local_subscriber=subscriber)
    return channel, subscriber

async def authenticator(route: str, authentication: Authentication):
    if isinstance(authentication, AuthenticationSimple):
        if authentication.password != b'12345':
            raise Exception('Authentication error')
    else:
        raise Exception('Unsupported authentication')
    
def handler_factory(socket):
    return RoutingRequestHandler(socket, router, authenticator)


def handle_client(reader, writer):
    RSocketServer(TransportTCP(reader, writer), handler_factory=handler_factory)

async def run_server(server_port):
    logging.info('Starting server at localhost:%s', server_port)

    server = await asyncio.start_server(handle_client, 'localhost', server_port)

    async with server:
        await server.serve_forever()
        
if __name__ == '__main__':
    port = sys.argv[1] if len(sys.argv) > 1 else 6565
    logging.basicConfig(level=logging.DEBUG)
    asyncio.run(run_server(port))
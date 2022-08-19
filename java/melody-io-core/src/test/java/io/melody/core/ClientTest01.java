package io.melody.core;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.List;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.messaging.rsocket.RSocketRequester;
import org.springframework.util.MimeTypeUtils;

import io.nats.client.Connection;
import io.nats.client.Nats;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufAllocator;
import io.netty.buffer.CompositeByteBuf;
import io.rsocket.Payload;
import io.rsocket.RSocket;
import io.rsocket.core.RSocketConnector;
import io.rsocket.metadata.AuthMetadataCodec;
import io.rsocket.metadata.CompositeMetadata;
import io.rsocket.metadata.CompositeMetadataCodec;
import io.rsocket.metadata.RoutingMetadata;
import io.rsocket.metadata.TaggingMetadataCodec;
import io.rsocket.metadata.WellKnownMimeType;
import io.rsocket.transport.netty.client.TcpClientTransport;
import io.rsocket.util.DefaultPayload;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@SpringBootTest
class ClientTest01 {
	@Autowired
	private RSocketRequester.Builder builder;
	private RSocketRequester client;
	
	@Test
	void testNATS() {
		try {
			Connection nc = Nats.connect("nats://localhost:4222");
			
			nc.publish("goes_command_dispatched", "auth.usersignin_task".getBytes(StandardCharsets.UTF_8));
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} finally {
		}
	}
	
	@Test
	void testRSocket() {
		try {
			RSocket rsocketClient = RSocketConnector.create()
			.metadataMimeType(WellKnownMimeType.MESSAGE_RSOCKET_COMPOSITE_METADATA.getString())
				    .connect(TcpClientTransport.create("127.0.0.1",7878))
				    .block();
			
			ByteBuf payloadData = ByteBufAllocator.DEFAULT.buffer().writeBytes("request msg".getBytes());
			RoutingMetadata routingMetadata = TaggingMetadataCodec.createRoutingMetadata(ByteBufAllocator.DEFAULT, List.of("initial2"));
			
			ByteBuf loginData =
			        AuthMetadataCodec.encodeSimpleMetadata(
			            ByteBufAllocator.DEFAULT, "tjikaljedy".toCharArray(), "password".toCharArray());
			
			ByteBuf loginData2 =
			        AuthMetadataCodec.encodeBearerMetadata(ByteBufAllocator.DEFAULT, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5LCJuYW1lIjoiSm9zZXBoIn0.OpOSSw7e485LOP5PrzScxHb7SR6sAOMRckfFwi4rp7o".toCharArray());
		
			CompositeByteBuf compositeMetadataBuffer = ByteBufAllocator.DEFAULT.compositeBuffer();
			
			CompositeMetadataCodec.encodeAndAddMetadata(
			        compositeMetadataBuffer, ByteBufAllocator.DEFAULT, WellKnownMimeType.MESSAGE_RSOCKET_ROUTING, routingMetadata.getContent());
			CompositeMetadataCodec.encodeAndAddMetadata(
			        compositeMetadataBuffer, ByteBufAllocator.DEFAULT, WellKnownMimeType.MESSAGE_RSOCKET_AUTHENTICATION, loginData2);
			
			
			
			rsocketClient.requestStream(DefaultPayload.create(payloadData, compositeMetadataBuffer))
		    .map(Payload::getDataUtf8)
		    .toIterable()
		    .forEach(System.out::println);
			
			/*client = builder
					.tcp("localhost", 7878);
			client.route("initial")
			.metadata("xx",  MimeTypeUtils.parseMimeType(WellKnownMimeType.MESSAGE_RSOCKET_AUTHENTICATION.getString()))
			.data("initial").retrieveFlux(String.class).subscribe();*/
		} catch (Exception e) {
			log.error(e.getMessage());
		}
	}

}

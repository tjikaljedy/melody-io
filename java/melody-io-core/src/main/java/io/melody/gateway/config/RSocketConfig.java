package io.melody.gateway.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.buffer.DefaultDataBufferFactory;
import org.springframework.http.codec.cbor.Jackson2CborDecoder;
import org.springframework.http.codec.cbor.Jackson2CborEncoder;
import org.springframework.messaging.rsocket.RSocketStrategies;

@Configuration
public class RSocketConfig {
	@Bean
	public RSocketStrategies rsocketStrategies() {
		return RSocketStrategies.builder().decoder(new Jackson2CborDecoder()).encoder(new Jackson2CborEncoder())
				.dataBufferFactory(new DefaultDataBufferFactory(true)).build();
	}
}

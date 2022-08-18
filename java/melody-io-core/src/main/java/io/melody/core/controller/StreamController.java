package io.melody.core.controller;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.messaging.rsocket.RSocketRequester;
import org.springframework.stereotype.Controller;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Flux;

@Slf4j
@Controller
public class StreamController implements ApplicationListener<ApplicationReadyEvent> {
	@Autowired
	private RSocketRequester.Builder builder;
	private RSocketRequester client;
	
	@Value("${core-config.rsocket-fs.enabled}")
	private boolean isEnabledRSocket;
	@Value("${core-config.rsocket-fs.host}")
	private String rsocketHost;
	@Value("${core-config.rsocket-fs.port}")
	private int rsocketPort;

	@MessageMapping("initial")
	public Flux<String> initial(@Payload String message) {
		log.info(">>> INITIAL-RSOCKET <<<");
		return Flux.empty();
	}

	@Override
	public void onApplicationEvent(ApplicationReadyEvent event) {
		try {
			client = builder.tcp(rsocketHost, rsocketPort);
			client.route("initial").data("initial").retrieveFlux(String.class).subscribe();
		} catch (Exception e) {
			log.error(e.getMessage());
		}
		
	}

}

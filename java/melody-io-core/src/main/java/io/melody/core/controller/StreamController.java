package io.melody.core.controller;


import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Controller;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Flux;

@Slf4j
@Controller
public class StreamController {


	@MessageMapping("initial")
	public Flux<String> initial(@Payload String message) {
		log.info(">>> INITIAL-RSOCKET <<<");
		return Flux.empty();
	}

}

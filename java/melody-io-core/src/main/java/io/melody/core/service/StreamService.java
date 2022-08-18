package io.melody.core.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.core.io.buffer.DataBufferFactory;
import org.springframework.core.io.buffer.DefaultDataBufferFactory;
import org.springframework.messaging.rsocket.RSocketRequester;
import org.springframework.stereotype.Service;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Service
public class StreamService implements ApplicationListener<ApplicationReadyEvent> {
	@Autowired
	private RSocketRequester.Builder builder;
	private DataBufferFactory dbf;
	private RSocketRequester requester;
	
	@Value("${core-config.rsocket-fs.enabled}")
	private boolean isEnabledRSocket;
	@Value("${core-config.rsocket-fs.host}")
	private String rsocketHost;
	@Value("${core-config.rsocket-fs.port}")
	private int rsocketPort;

	
	
	public DataBufferFactory getDbf() {
		return dbf;
	}


	@Override
	public void onApplicationEvent(ApplicationReadyEvent event) {
		try {
			requester = builder.tcp(rsocketHost, rsocketPort);
			requester.route("initial").data("initial").retrieveFlux(String.class).subscribe();
			
			dbf = new DefaultDataBufferFactory();
		} catch (Exception e) {
			log.error(e.getMessage());
		}
	}


	

}

package io.melody.core;

import java.util.concurrent.TimeUnit;

import javax.annotation.PreDestroy;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.boot.autoconfigure.security.reactive.ReactiveUserDetailsServiceAutoConfiguration;
import org.springframework.boot.autoconfigure.websocket.servlet.WebSocketServletAutoConfiguration;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.boot.web.servlet.support.SpringBootServletInitializer;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.event.EventListener;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@SpringBootApplication
@ComponentScan({ "io.melody.gateway" })
@EnableAutoConfiguration(exclude = { WebSocketServletAutoConfiguration.class,
		ReactiveUserDetailsServiceAutoConfiguration.class, DataSourceAutoConfiguration.class})
public class MelodyGateway extends SpringBootServletInitializer {

	public static void main(String[] args) {
		SpringApplication.run(MelodyGateway.class, args);
	}

	@Override
	protected SpringApplicationBuilder configure(SpringApplicationBuilder builder) {
		setRegisterErrorPageFilter(false);
		builder.sources(MelodyGateway.class);
		return builder;
	}
	
	
	@EventListener(ApplicationReadyEvent.class)
	public void readyToUse(ApplicationReadyEvent event) {
		log.info(">>> SERVER STARTUP <<<");
		try {
			TimeUnit.SECONDS.sleep(1);
			
		} catch (Exception e) {
			log.warn(">>> FAIL <<<");
		}
	}
	
	@PreDestroy
	public void onDestroy() throws Exception {
		log.info(">>> SERVER SHUTDOWN <<<");
		
	}

}

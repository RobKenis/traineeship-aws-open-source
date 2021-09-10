package com.robkenis.beanstalkdemo;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDB;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDBClient;
import com.amazonaws.services.dynamodbv2.document.DynamoDB;
import com.amazonaws.services.dynamodbv2.model.AttributeValue;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Date;
import java.util.Map;
import java.util.UUID;

@SpringBootApplication
@RestController
public class BeanstalkDemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(BeanstalkDemoApplication.class, args);
	}

	private final AmazonDynamoDB client = AmazonDynamoDBClient.builder().withRegion(Regions.EU_WEST_1).build();

	@PostMapping
	public void save(){
		client.putItem("table-van-rob", Map.ofEntries(
				Map.entry("id", new AttributeValue().withS(UUID.randomUUID().toString())),
				Map.entry("timestamp", new AttributeValue().withS(new Date().toString()))
		));
	}

}

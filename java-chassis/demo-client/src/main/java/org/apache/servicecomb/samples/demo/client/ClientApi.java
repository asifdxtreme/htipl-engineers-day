package org.apache.servicecomb.samples.demo.client;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import org.apache.servicecomb.provider.rest.common.RestSchema;
import org.apache.servicecomb.provider.springmvc.reference.RestTemplateBuilder;
import org.springframework.web.client.RestTemplate;

@RestSchema(schemaId = "helloclient")
@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class ClientApi {
	private static RestTemplate templateNew = RestTemplateBuilder.create();

	@GET
	@Path("/sayhello")
	public String sayHello() {
		String microserviceName = "DemoServer";
		String cseUrlPrefix = "cse://" + microserviceName + "/";
		String result = templateNew.getForObject(cseUrlPrefix + "sayhi", String.class);
		System.out.println(result);
		return result;
	}
}
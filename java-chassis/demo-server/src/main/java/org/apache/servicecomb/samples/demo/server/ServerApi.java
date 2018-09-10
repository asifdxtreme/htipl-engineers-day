
package org.apache.servicecomb.samples.demo.server;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import org.apache.servicecomb.provider.rest.common.RestSchema;

@RestSchema(schemaId = "hello")
@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class ServerApi {

	@Path("/sayhi")
	@GET
	public String sayHi() {
		return "Hello Guest: I am server 1.0";
	}
}

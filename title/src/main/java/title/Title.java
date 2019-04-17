package title;

import io.vertx.core.Vertx;
import io.vertx.core.http.HttpServer;
import io.vertx.core.http.HttpServerResponse;
import io.vertx.ext.web.Router;

public class Title {
    public static void main(String[] args) {
        Vertx vertx = Vertx.vertx();

        Router router = Router.router(vertx);
        router.route("/title").handler(routingContext -> {
            HttpServerResponse response = routingContext.response();

            response.putHeader("content-type", "text/plain");
            response.end("Google Cloud Quizz");
        });

        HttpServer server = vertx.createHttpServer();
        server.requestHandler(router).listen(8888);
    }
}

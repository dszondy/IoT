package main;
import com.sun.net.httpserver.HttpContext;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;
import dal.DBConnector;
import model.ParkPlace;

import java.io.BufferedOutputStream;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.InetSocketAddress;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Date;
import java.util.StringTokenizer;

//places/list?id=desc

// The tutorial can be found just here on the SSaurel's Blog :
// https://www.ssaurel.com/blog/create-a-simple-http-web-server-in-java
// Each Client Connection will be managed in a dedicated Thread
public class JavaHTTPServer {
    HttpServer server;
    private static JavaHTTPServer ourInstance = new JavaHTTPServer();
    public static JavaHTTPServer getInstance() {
        return ourInstance;
    }


    protected JavaHTTPServer(){
        System.out.println("The http server is starting.");
        try {
                server = HttpServer.create(new InetSocketAddress(9849), 0);
            } catch (IOException e) {
                e.printStackTrace();
            }
        HttpContext context = server.createContext("/places/list");
        context.setHandler(new placesListHandler());
        server.start();
        System.out.println("The http server is up.");
    }

    private class placesListHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange exchange) throws IOException {
            System.out.println("");
            String response = ParkPlace.convertArrayToJson(DBConnector.getInstance().QueryLots());
            exchange.sendResponseHeaders(200, response.getBytes().length);
            exchange.getResponseBody().write(response.getBytes());
            exchange.close();
        }
    }
}
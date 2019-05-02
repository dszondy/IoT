package main;

import dal.DBConnector;

import java.util.concurrent.TimeUnit;


public class Main {

    public static void main(String[] args) {
        JavaHTTPServer.getInstance();
        DBConnector.getInstance();
        while(true) {
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

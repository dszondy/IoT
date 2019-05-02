package dal;

import model.ParkPlace;

import java.sql.*;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;

public class DBConnector {
    private static DBConnector ourInstance = new DBConnector();

    public static DBConnector getInstance() {
        return ourInstance;
    }
    private Connection connection;
    public DBConnector() {
        try {
            Class.forName("org.postgresql.Driver");
        }catch (ClassNotFoundException exception){
            System.err.println("The jdbc driver is missing");
            return;
        }
        while (true) {
            try {
                connection = DriverManager.getConnection("jdbc:postgresql://parikng_database:5432/parking_database", "postgres", "admin");
                if (connection != null)
                    break;

            } catch (Exception e1) {
                System.err.println("Error while connectiong to database. Retrying in 10 seconds.");
                System.err.println(e1.getMessage());
                e1.printStackTrace();
                try {
                    wait(10000);
                } catch (InterruptedException e2) {
                    System.err.println("Interrupted while waiting after unsuccessful connection to database.");
                    e2.printStackTrace();
                }
            }
        }
        System.out.println("Successfully connected to PostgreSQL database!");
    }
    public Collection<ParkPlace> QueryLots(){
        try {
            Statement stmt = connection.createStatement();
            ResultSet rs;

            rs = stmt.executeQuery("SELECT id, state FROM lot_states;");
            List<ParkPlace> places = new ArrayList<>();
            while ( rs.next() ) {
                places.add(new ParkPlace(rs.getInt("id"), rs.getBoolean("state")));
            }
            return places;
        } catch (Exception e) {
            System.err.println("Unable to get parkingPlaces from database");
            System.err.println(e.getMessage());
        }
        return null;
    }
}

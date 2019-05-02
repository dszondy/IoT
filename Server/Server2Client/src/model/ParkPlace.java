package model;

import resources.JSONArray;
import resources.JSONObject;

import java.util.Collection;

public class ParkPlace {
    public ParkPlace(Integer id, Boolean occupied) {
        this.id = id;
        this.occupied = occupied;
    }

    public Integer id;
    public Boolean occupied;
    public String toJson(){
        JSONObject object = new JSONObject();
        object.put("id", id);
        object.put("occupied", occupied);
        return object.toString();
    }
    public static String convertArrayToJson(Collection<ParkPlace> places){
        JSONArray jsonPlaces = new JSONArray();
        for (ParkPlace p : places){
            jsonPlaces.put(new JSONObject(p.toJson()));
        }
        return jsonPlaces.toString();
    }
}

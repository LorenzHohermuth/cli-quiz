package org.quiz;

import java.util.HashMap;

public class Fetcher {

    public HashMap<String, String> fetch (String text) {
        HashMap<String, String> map = new HashMap<>();
        String[] keyValuePairs = text.split(";");
        for (String pair : keyValuePairs) {
            String[] valuePair = pair.split("=");
            map.put(valuePair[0], valuePair[1]);
        }
        return map;
    }


}

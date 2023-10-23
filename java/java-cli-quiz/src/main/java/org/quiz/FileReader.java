package org.quiz;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

public class FileReader {
    private InputStream is;
    FileReader(InputStream is) {
        this.is = is;
    }

    public String getContent() {
        try {
            return readFile(is);
        } catch (IOException e) {
            e.printStackTrace();
            return "";
        }
    }

    private String readFile(InputStream stream) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(stream));
        StringBuilder sb = new StringBuilder();
        String line;
        while ((line = br.readLine()) != null) {
            sb.append(line);
            sb.append("\n");
        }
        return sb.toString();
    }
}

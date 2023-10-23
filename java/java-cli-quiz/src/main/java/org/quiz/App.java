package org.quiz;

import java.util.Scanner;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        FileReader fr = new FileReader(App.class.getResourceAsStream("scrapingScript.js"));
        Scanner scanner = new Scanner(System.in);
        Fetcher f = new Fetcher();
        System.out.println("Q >_");
        System.out.println("Please enter the JS code in your Quizlet set");
        System.out.println("and paste the Output in the Command Line");
        System.out.println("JS :\n\n" + fr.getContent() + "\n");
        String inputQuizlet = scanner.nextLine();
        System.out.println(f.fetch(inputQuizlet));
    }
}

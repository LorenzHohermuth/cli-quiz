package org.quiz.learning.assets;

public class Card {

    static byte ratingReward = 10;
    static byte ratingPunishment = 20;

    private final String term;
    private final String definition;
    private byte rating = 25; // A value from 0 to 100
    // byte is from -128 to 127
    private Boolean isCardLearned = false;

    public Card(String term, String definition) {
        this.term = term;
        this.definition = definition;
    }

    public Card ratingUp () {
        if (rating + ratingReward >= 100) {
            rating = 100;
            isCardLearned = true;
        } else {
            rating += ratingReward;
        }
        return this;
    }

    public Card ratingDown () {
        if (rating - ratingPunishment <= 0) {
            rating = 0;
        } else {
            rating -= ratingPunishment;
        }
        return this;
    }


}

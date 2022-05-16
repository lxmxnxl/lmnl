package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func announceAdminCreds() {

	words := []string{
		"amped", "beef", "bear", "bee", "bendy", "berry", "best", "big", "bird", "blep", "bliss", "blob", "blop", "blue", "blup", "boo", "box", "brave", "bravo", "brup", "bubby", "burly", "calm", "cargo", "cat", "champ", "charm", "cheer", "clean", "clear", "cloud", "couch", "dab", "dear", "deep", "dog", "dream", "drum", "eagle", "easy", "enjoy", "epic", "excel", "fair", "fairy", "faith", "fate", "fancy", "fast", "first", "fish", "fit", "flip", "floop", "flute", "free", "fresh", "fruit", "fun", "funny", "fuzzy", "gift", "glam", "glee", "glob", "goal", "good", "gooey", "grape", "grace", "great", "green", "ham", "happy", "heart", "hello", "hero", "holy", "home", "honor", "hope", "huge", "humor", "icy", "inky", "jewel", "jolly", "joy", "juicy", "jumbo", "keen", "key", "kind", "kiss", "kiwi", "large", "laugh", "leaky", "lefty", "life", "lil", "look", "love", "loyal", "luck", "lucky", "macho", "magic", "major", "merry", "moose", "moral", "nacho", "neat", "newt", "new", "nice", "noble", "nurse", "open", "panda", "peace", "peach", "plane", "plop", "plug", "pog", "port", "post", "power", "pride", "prime", "prize", "proud", "pug", "pure", "quick", "quiet", "rad", "ready", "real", "red", "regal", "relax", "rich", "ripe", "rogue", "rosy", "safe", "scout", "sharp", "shine", "shiny", "shoe", "silky", "silly", "skip", "skup", "small", "smart", "smile", "sonic", "spark", "spicy", "star", "super", "sweet", "tacky", "tall", "tasty", "team", "tech", "terp", "testy", "that", "this", "time", "tippy", "true", "trust", "truth", "vast", "warm", "whole", "wind", "wise", "witty", "wow", "yay", "yeah", "yes", "young", "youth", "yummy", "zeal", "zen", "zest", "zesty", "zip", "zippy", "plane", "bear", "bird", "drum", "fish", "cat", "dog", "cloud", "square", "circle", "blep", "skup", "skip", "bubby", "bubbie", "couch", "orange", "blue", "cargo", "gopher", "plaster", "pastry", "muffin", "gooble", "gobble", "woadie", "fruit", "peach", "plop", "brup", "floop", "whistle", "wind", "breeze", "star", "blup", "blop", "kiwi", "berry", "laptop", "shoe", "cheese", "sprout", "pepper", "moose", "eagle", "terp", "scout", "brrup", "smoosh", "bee", "flower", "squirrel", "ham", "stick", "plug", "pug", "port", "flute", "portal", "castle", "box", "time", "trust", "look", "dolphin", "frinkle", "red", "orange", "yellow", "green", "blue", "indigo", "lil", "violet", "large", "small", "lovely", "lucky", "tasty", "northern", "medium", "wise", "joyful", "speedy", "quick", "fast", "post", "lifted", "sunset", "boutique", "shiny", "deep", "brushed", "smooth", "big", "tall", "double", "triple", "this", "that", "calm", "whimsical", "spunky", "witty", "gentle", "jolly", "lively", "gifted", "petite", "clean", "clear", "amiable", "brave", "bright", "cheeky", "cheer", "lucky", "macho", "static", "jittery", "prickly", "ripe", "zippy", "zip", "fuzzy", "silky", "silly", "vast", "zany", "rosy", "rogue", "burly", "icy", "obedient", "tacky", "stacked", "parched", "wilted", "crunk", "tech", "spiced", "tectonic", "floaty", "pointy", "testy", "zesty", "lefty", "ground", "leaky", "sonic", "woody", "inky", "slinky", "bendy", "fidgety", "tippy", "panda"}

	f1 := roll(0, len(words)-1)
	f2 := roll(0, len(words)-1)
	f3 := roll(0, len(words)-1)
	f4 := roll(0, len(words)-1)

	adminStr := strings.Join([]string{words[f1], words[f2], words[f3], words[f4]}, "-")

	log.Println("Your admin credential is: " + adminStr)

}

func roll(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

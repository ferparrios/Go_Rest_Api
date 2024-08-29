package quotes

import (
	"strings"
)

type Quote struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
	Source string `json:"source"`
}

var quotes = []Quote{
  {
   "With great power comes great responsibility.",
    "Uncle Ben",
    "Amazing Fantasy #15 (1962)",
  },
  {
   "Your friendly neighborhood Spider-Man.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "I'm Spider-Man, not Spider-Boy. Now, where's my coffee?",
    "Spider-Man",
    "Spider-Man: Homecoming (2017)",
  },
  {
   "I believe there's a hero in all of us, that keeps us honest, gives us strength, makes us noble, and finally allows us to die with pride, even though sometimes we have to be steady, and give up the thing we want the most. Even our dreams.",
    "Aunt May",
    "Spider-Man 2 (2004)",
  },
  {
   "You know, I guess one person can make a difference... nuff said.",
    "Spider-Man",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "Everybody needs help sometimes, Peter. Even Spider-Man.",
    "Mary Jane Watson",
    "Spider-Man: Homecoming (2017)",
  },
  {
   "I don't want to fight you, Flash. I wouldn't want to hurt your fist.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "I'm not like you. You're a murderer.",
    "Spider-Man",
    "Spider-Man: Homecoming (2017)",
  },
  {
   "I have a knack for getting into trouble.",
    "Spider-Man",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "I'm just your friendly neighborhood Spider-Man!",
    "Spider-Man",
    "Spider-Man: Into the Spider-Verse (2018)",
  },
  {
   "When you can do the things that I can, but you don't, and then the bad things happen? They happen because of you.",
    "Spider-Man",
    "Spider-Man: Homecoming (2017)",
  },
  {
   "Parker luck.",
    "Spider-Man",
    "The Amazing Spider-Man #500 (2003)",
  },
  {
   "Why do I get the feeling that whenever we meet, I'm going to be the one who ends up bleeding?",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "Not everyone is meant to make a difference, but for me, the choice to lead an ordinary life is no longer an option.",
    "Spider-Man",
    "Spider-Man (2002)",
  },
  {
   "I'm not alone. Not while I'm Spider-Man.",
    "Spider-Man",
    "The Amazing Spider-Man #573 (2008)",
  },
  {
   "I don't know if I'm a superhero or a menace. Probably a little bit of both.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (2014)",
  },
  {
   "I'll never stop trying to make a difference. As long as there's hope.",
    "Spider-Man",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "I'm the one who has to die when it's time for me to die, so let me live my life, the way I want to.",
    "Spider-Man",
    "The Spectacular Spider-Man #131 (1987)",
  },
  {
   "I'm not a hero, just a man who knows how to swing a web.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "The one thing that's clear to me, is that you can't fight fire with fire.",
    "Spider-Man",
    "Spider-Man: Homecoming (2017)",
  },
  {
   "With great power comes great responsibility.",
    "Uncle Ben",
    "Amazing Fantasy #15 (1962)",
  },
  {
   "My spider-sense is tingling.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "I have a responsibility to use my abilities to help people in need.",
    "Spider-Man",
    "Spider-Man: Into the Spider-Verse (2018)",
  },
  {
   "This is my gift, my curse. Who am I? I'm Spider-Man.",
    "Spider-Man",
    "Spider-Man (2002)",
  },
  {
   "The only way to get through life is to laugh your way through it. You either have to laugh or cry. I prefer to laugh. Crying gives me a headache.",
    "Spider-Man",
    "Spider-Man: Blue (2002)",
  },
  {
   "It's not about how hard you hit. It's about how hard you can get hit and keep moving forward.",
    "Spider-Man",
    "Amazing Spider-Man #537 (2007)",
  },
  {
   "I'm stronger than I thought. I can lift a car. If I had to, I could probably even lift a... bus!",
    "Spider-Man",
    "The Amazing Spider-Man #33 (1966)",
  },
  {
   "I've got you, I've got you. I swear on my spider-sense, I won't let you fall.",
    "Spider-Man",
    "Spider-Man: Blue (2002)",
  },
  {
   "The world needs heroes, and it's time for me to step up and be one.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (2014)",
  },
  {
   "I'm Peter Parker. I'm Spider-Man. And I'm not afraid of you.",
    "Spider-Man",
    "The Amazing Spider-Man #33 (1966)",
  },
  {
   "You know, I guess one person can make a difference. Nuff said.",
    "Stan Lee",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "You're the Spiderling. Crime-fighting spider. You're Spider-Boy?",
    "Tony Stark",
    "Captain America: Civil War (2016)",
  },
  {
   "It's not enough to be friendly. You have to be a friend.",
    "Spider-Man",
    "Spider-Man: Blue (2002)",
  },
  {
   "I have to stop him... because I created him.",
    "Spider-Man",
    "Spider-Man 3 (2007)",
  },
  {
   "I've always felt like I'm different from everybody else. Like I don't quite fit in. Maybe because I'm not supposed to.",
    "Miles Morales",
    "Spider-Man: Into the Spider-Verse (2018)",
  },
  {
   "It's easy to do good when bad things aren't happening to you. But when the bad things come, how are you gonna act then?",
    "Spider-Man",
    "The Amazing Spider-Man #537 (2007)",
  },
  {
   "I am the Ultimate Spider-Man. I'm not some guy in a bat suit.",
    "Miles Morales",
    "Ultimate Comics Spider-Man #1 (2011)",
  },
  {
   "If you keep something as complicated as love stored up inside, it could make you sick.",
    "Spider-Man",
    "Spider-Man: Blue (2002)",
  },
  {
   "I know that I have a responsibility to use my powers to help people. That's why I'm Spider-Man.",
    "Spider-Man",
    "Ultimate Spider-Man #1 (2000)",
  },
  {
   "With great power, comes great responsibility.",
    "Uncle Ben",
    "Spider-Man (2002)",
  },
  {
   "I'm not Superman, you know. I can't be everywhere at once.",
    "Spider-Man",
    "The Amazing Spider-Man #1 (1963)",
  },
  {
   "The one thing they love more than a hero is to see a hero fail, fall, die trying. In spite of everything you've done for them, eventually they will hate you.",
    "Green Goblin",
    "Spider-Man (2002)",
  },
  {
   "It's not about how hard you hit. It's about how hard you can get hit and keep moving forward.",
    "Spider-Man",
    "The Amazing Spider-Man #537 (2007)",
  },
  {
   "Some people are just born with tragedy in their blood.",
    "Spider-Man",
    "Spider-Man: Reign #1 (2006)",
  },
  {
   "This is my gift, my curse. Who am I? I'm Spider-Man.",
    "Spider-Man",
    "Spider-Man (2002)",
  },
  {
   "I can't turn my back on those in need. That's when I'll know I'm a true hero.",
    "Spider-Man",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "Things will never be the same. I'm a superhero. And I'm a married man.",
    "Spider-Man",
    "Amazing Spider-Man Annual #21 (1987)",
  },
  {
   "The bigger the responsibility, the bigger the target.",
    "Spider-Man",
    "Spider-Man: The Animated Series (1994)",
  },
  {
   "I don't know if I'm a superhero or not. But I don't really care.",
    "Miles Morales",
    "Spider-Man: Miles Morales (2018)",
  },
  {
   "No matter what I do, no matter how hard I try, the ones I love will always be the ones who pay.",
    "Spider-Man",
    "Spider-Man: Back in Black #1 (2007)",
  },
}

func GetQuotes()[]Quote{
	return quotes
}

func SearchQuotes(keyword string) []Quote{
	var result [] Quote
	for _, quote := range quotes {
		if strings.Contains(strings.ToLower(quote.Quote), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(quote.Author), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(quote.Source), strings.ToLower(keyword)){
				result = append(result, quote)
			}
	}
	return result
}
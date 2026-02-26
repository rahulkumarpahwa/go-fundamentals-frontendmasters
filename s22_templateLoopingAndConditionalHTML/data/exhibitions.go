package data

type Exhibition struct {
	Title           string
	Description     string
	Image           string
	Color           string
	CurrentlyOpened bool
}

var list = []Exhibition{ // Remember this is the constructor of a slice of exhibition.The constructor angle is always curly braces, soit's not actually a square brackets as an array in JSON, a reminder.We have the type first and then we open curly braces, andthat's the constructor of types in Go.
	// slice of exhibition struct.
	{
		Title:           "Life in Ancient Greek",
		Description:     "Uncover the world of ancient Greece through the sculptures, tools, and jewelry found in ruins from over 2000 years ago that have been unearthed through modern science and technology.",
		Image:           "ancient-greece.png",
		Color:           "red",
		CurrentlyOpened: true,
	},
	{
		Title:           "Aristotle: Life and Legacy",
		Description:     "Explore the life and legacy of the great philosopher Aristotle, one of the most influential thinkers of all time. Through rare artifacts and ancient texts, learn about his ideas on ethics, politics, and metaphysics that have shaped the world for centuries.",
		Image:           "aristotle.png",
		Color:           "purple",
		CurrentlyOpened: false,
	},
	{
		Title:           "Chameleon: Colorful Adaptations",
		Description:     "Discover the amazing world of chameleons and their incredible ability to change color. Through interactive displays and live chameleon exhibits, learn about the science behind their color changing and how they use it to communicate and camouflage in their environments.",
		Image:           "colorful-adaptations.png",
		Color:           "blue",
		CurrentlyOpened: false,
	},
	{
		Title:           "Sea Monsters: Myth and Reality",
		Description:     "Dive into the world of sea monsters and explore the myths and legends that have captured our imaginations for centuries. Through fossils, ancient maps, and interactive displays, discover the truth behind the stories and learn about the real-life creatures that inhabit our oceans.",
		Image:           "sea-monsters.png", // And also some weird thing that I'm not sure you will like isthat the trailing comma is mandatory.So if you try to remove the trailing commas,you can see that image has a trailing comma even if there is no other property.And also there is a trailing comma at the end of the collection(struct), it's mandatory.
		Color:           "green",
		CurrentlyOpened: true,
	},
}

// Setter
func Add(e Exhibition) {
	list = append(list, e)
}

// Getter
func GetAll() []Exhibition {
	return list
}

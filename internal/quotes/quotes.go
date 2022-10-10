package quotes

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Collection struct {
}

func NewCollection() *Collection {
	return &Collection{}
}

func (s Collection) GetRandQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

// It was grabbed out https://webcache.googleusercontent.com/search?q=cache:6FvWB6a8uQ4J:https://examples.yourdictionary.com/articles/words-of-wisdom.html&cd=2&hl=ru&ct=clnk&gl=ru&client=safari

var quotes = []string{
	"You are perfect because of your imperfections.",
	"Do what inspires you. Life is too short not to love the job you do every day.",
	"Complaining will not get anything done.",
	"At the end of your day, you’ve done your best. Even if you haven’t accomplished all that’s on your list. You’ve given it you’re all.",
	"You don’t need to have it figured all out. Taking the wrong path is part of the process.",
	"Never lose yourself because of someone else. You are perfect just the way you are.",
	"Trust your gut. If you ever feel it's not right, then it's not.",
	"A smile is a free way to brighten someone’s day.",
	"You can never take too many pictures. This is documentation of their life as a baby.",
	"Every baby is unique. Do what you think is best for your child.",
	"There isn’t any perfect parenting technique. You just have to use what works best for your family.",
	"No matter what anyone tells you, your instincts are always right.",
	"Patience is a learned skill that you have to practice in parenting every day. Adult time-outs are completely okay.",
	"Every moment is precious, even the bad ones. And typically the horrible moments are the ones you’ll laugh at down the road.",
	"Sleep is one of the hardest things to come by that you need the most. Therefore, it’s important to sleep when you can.",
	"Remember to laugh at the small stuff. Spilled cereal can quickly be cleaned up.",
	"Marriage takes work. But the payoff is years of happiness.",
	"When you dream, dream together. Commitment is the key to a long-lasting relationship.",
	"Marriage is two individuals moving together on a path. You can still be who you are and love someone else.",
	"It’s important to make time for each other. Little changes are hard to notice when you see someone every day.",
	"While it’s important to support each other, it’s also crucial to surround yourself with people that support you as a couple.",
	"Marriage is about compromise. Being a winner isn’t important, working together is.",
	"The key to a successful marriage is being strong when your partner is weak.",
	"Your partner is often the one that sees you at your worst, make sure they also see you at your best too.",
	"Aim for your dreams, but don’t lose yourself along the way. Sometimes the road to greatness takes creating your own path.",
	"Graduation is the end but also the beginning. It’s the start of a new chapter written entirely by you.",
	"The journey to enlightenment doesn’t end when you graduate, it is only just beginning. Learning is a lifelong process.",
	"Dreams are ever-evolving. Therefore, the dream you had at the beginning can very well be different than the one you have at the end.",
	"Failure is an option. Because you need to fail to grow. Each failure makes success that much closer.",
	"Celebrate every small success. The path to greatness is a rocky, but enlightened, journey.",
	"Your story is unique. It can only be forged by you.",
	"Graduation was the match that sparked your future. Light the night up with your potential.",
	"The teacher who is indeed wise does not bid you to enter the house of his wisdom but rather leads you to the threshold of your mind.",
	"I do not think much of a man who is not wiser today than he was yesterday.",
	"Good people are good because they've come to wisdom through failure. We get very little wisdom from success, you know.",
	"It is unwise to be too sure of one's own wisdom. It is healthy to be reminded that the strongest might weaken and the wisest might err.",
	"It's not that I'm so smart, it's just that I stay with problems longer.",
	"I am the wisest man alive, for I know one thing, and that is that I know nothing.",
	"A loving heart is the truest wisdom.",
	"It is characteristic of wisdom not to do desperate things.",
	"Be strong and bold; have no fear or dread of them, because it is the Lord your God who goes with you; he will not fail you or forsake you.",
	"When doubts filled my mind, your comfort gave me renewed hope and cheer.",
	"Don't love money; be satisfied with what you have.",
	"If any of you lacks wisdom, let him ask of God, who gives to all liberally and without reproach, and it will be given to him.",
	"Be very careful about what you think. Your thoughts run your life.",
	"Therefore do not worry about tomorrow, for tomorrow will worry about itself. Each day has enough trouble of its own.",
	"For we walk by faith, not by sight.",
	"Humble yourselves in the presence of the Lord, and He will exalt you.",
	"Each morning when I open my eyes I say to myself: I, not events, have the power to make me happy or unhappy today. I can choose which it shall be. Yesterday is dead, tomorrow hasn't arrived yet. I have just one day, today, and I'm going to be happy in it.",
	"If you want to be happy, set a goal that commands your thoughts, liberates your energy, and inspires your hopes.",
	"Happiness is a perfume you cannot pour on others without getting a few drops on yourself.",
	"Happiness cannot be traveled to, owned, earned, worn or consumed. Happiness is the spiritual experience of living every minute with love, grace and gratitude.",
	"Happiness comes when your work and words are of benefit to yourself and others.",
	"Happiness cannot come from without. It must come from within. It is not what we see and touch or that which others do for us which makes us happy; it is that which we think and feel and do, first for the other fellow and then for ourselves.",
	"Everything is a gift of the universe — even joy, anger, jealously, frustration, or separateness. Everything is perfect either for our growth or our enjoyment.",
	"Happiness cannot be traveled to, owned, earned, worn or consumed. Happiness is the spiritual experience of living every minute with love, grace, and gratitude.",
}

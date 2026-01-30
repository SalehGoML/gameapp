package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []uint
	startTime   time.Time
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
	Answer []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	Player     uint
	QuestionID uint
	Choice     PossibleAnswerChoice
}

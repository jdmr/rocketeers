package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// Question struct
type Question struct {
	ID       string    `json:"id"`
	Book     string    `json:"book"`
	Chapter  string    `json:"chapter"`
	Verses   string    `json:"verses"`
	Question string    `json:"question"`
	Answers  []*Answer `json:"answers"`
	Finished bool      `json:"finished"`
}

// Answer struct
type Answer struct {
	ID           string `json:"id"`
	Answer       string `json:"answer"`
	Status       bool   `json:"status"`
	TeamAnswerID string `json:"team_answer_id"`
	Checked      bool   `json:"checked"`
}

// Game struct
type Game struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Seconds    int            `json:"seconds"`
	Questions  int            `json:"questions"`
	Questions2 []*Question    `json:"questions2"`
	Status     string         `json:"status"`
	Created    time.Time      `json:"created"`
	Chapters   []*GameChapter `json:"chapters"`
	Teams      []*Team        `json:"teams"`
}

// GameChapter struct
type GameChapter struct {
	ID      string `json:"id"`
	Book    string `json:"book"`
	Chapter string `json:"chapter"`
}

// Team struct
type Team struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Answers []*Answer `json:"answers"`
	Points  int       `json:"points"`
}

func getQuestionsController(c echo.Context) error {
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select
			q.id
			, q.book
			, q.chapter
			, q.verses
			, q.question
			, COALESCE(a.id, '')
			, COALESCE(a.answer, '')
			, COALESCE(a.status, false)
		from pbe.questions q
		left join pbe.answers a on a.question_id = q.id
		order by q.book, q.chapter, q.verses, q.id, a.answer
	`)
	if err != nil {
		log.Error("Could not get question: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get question: "+err.Error())
	}

	questions := []*Question{}
	question := &Question{}
	for rows.Next() {
		var (
			id           string
			book         string
			chapter      string
			verses       string
			questionText string
			answerID     string
			answer       string
			status       bool
		)
		err = rows.Scan(&id, &book, &chapter, &verses, &questionText, &answerID, &answer, &status)
		if err != nil {
			log.Error("Could not get question: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get question: "+err.Error())
		}

		if question.ID != id {
			question = &Question{
				ID:       id,
				Book:     book,
				Chapter:  chapter,
				Verses:   verses,
				Question: questionText,
			}
			questions = append(questions, question)
		}
		if len(answerID) > 0 {
			a := &Answer{
				ID:     answerID,
				Answer: answer,
				Status: status,
			}
			question.Answers = append(question.Answers, a)
		}
	}

	return c.JSON(http.StatusOK, questions)
}

func addQuestionController(c echo.Context) error {
	question := &Question{}
	err := c.Bind(&question)
	if err != nil {
		log.Error("Could not parse question: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not parse question: "+err.Error())
	}

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not create database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create database transaction: "+err.Error())
	}

	question.ID, _ = UUID()
	_, err = tx.Exec(`
		insert into pbe.questions(id, book, chapter, verses, question) values(?,?,?,?,?)
	`, question.ID, question.Book, question.Chapter, question.Verses, question.Question)
	if err != nil {
		log.Error("Could not create question: ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not create question: "+err.Error())
	}

	for _, answer := range question.Answers {
		answer.ID, _ = UUID()
		_, err = tx.Exec(`
			insert into pbe.answers(id, answer, status, question_id)
			values(?,?,?,?)
		`, answer.ID, answer.Answer, answer.Status, question.ID)
		if err != nil {
			log.Error("Could not create answer: ", err)
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Coul dnot create answer: "+err.Error())
		}
	}

	tx.Commit()
	return c.JSON(http.StatusOK, question)
}

func deleteQuestionController(c echo.Context) error {
	questionID := c.Param("questionID")
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	_, err = conn.Exec(`
		delete from pbe.questions where id = ?
	`, questionID)
	if err != nil {
		log.Error("Could not delete question: ", questionID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not delete question: "+questionID+" : "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func deleteAnswerController(c echo.Context) error {
	answerID := c.Param("answerID")
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	_, err = conn.Exec(`
		delete from pbe.answers where id = ?
	`, answerID)
	if err != nil {
		log.Error("Could not delete answer: ", answerID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not delete answer: "+answerID+" : "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func getQuestionController(c echo.Context) error {
	questionID := c.Param("questionID")

	question, err := getQuestion(questionID)
	if err != nil {
		log.Error("Could not get question: ", questionID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get question: "+questionID+" : "+err.Error())
	}

	return c.JSON(http.StatusOK, question)
}

func getQuestion(questionID string) (*Question, error) {
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select 
			q.book
			, q.chapter
			, q.verses
			, q.question
			, COALESCE(a.id, '')
			, COALESCE(a.answer, '')
			, COALESCE(a.status, false)
		from pbe.questions q
		left join pbe.answers a on a.question_id = q.id
		where q.id = ?
		order by q.id, a.answer
	`, questionID)
	if err != nil {
		log.Error("Could not get question: ", err)
		return nil, err
	}

	question := &Question{
		ID: questionID,
	}
	for rows.Next() {
		var (
			book         string
			chapter      string
			verses       string
			questionText string
			answerID     string
			answer       string
			status       bool
		)
		err = rows.Scan(&book, &chapter, &verses, &questionText, &answerID, &answer, &status)
		if err != nil {
			log.Error("Could not get question: ", err)
			return nil, err
		}

		if len(question.Question) == 0 {
			question.Book = book
			question.Chapter = chapter
			question.Verses = verses
			question.Question = questionText
		}
		if len(answerID) > 0 {
			a := &Answer{
				ID:     answerID,
				Answer: answer,
				Status: status,
			}
			question.Answers = append(question.Answers, a)
		}
	}
	return question, nil
}

func addAnswerController(c echo.Context) error {
	questionID := c.Param("questionID")
	answer := &Answer{}
	err := c.Bind(&answer)
	if err != nil {
		log.Error("Could not parse answer: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not parse answer: "+err.Error())
	}

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	answer.ID, _ = UUID()

	_, err = conn.Exec(`
		insert into pbe.answers(id, answer, status, question_id)
		values(?,?,?,?)
	`, answer.ID, answer.Answer, answer.Status, questionID)
	if err != nil {
		log.Error("Could not create answer: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create answer: "+err.Error())
	}

	return c.JSON(http.StatusOK, answer)
}

func addGameController(c echo.Context) error {
	game := &Game{}
	err := c.Bind(&game)
	if err != nil {
		log.Error("Could not create game: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create game: "+err.Error())
	}

	game.ID, _ = UUID()

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not create database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create database transaction: "+err.Error())
	}

	_, err = tx.Exec(`
		insert into pbe.games(id, name, seconds, created, questions, status)
		values(?,?,?,NOW(),?,'OPEN')
	`, game.ID, game.Name, game.Seconds, game.Questions)
	if err != nil {
		log.Error("Could not create game: ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not create game: "+err.Error())
	}

	teamID, _ := UUID()
	_, err = tx.Exec(`
		insert into pbe.teams(id, name, game_id)
		values(?,'Home',?)
	`, teamID, game.ID)
	if err != nil {
		log.Error("Could not add team: ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not add team: "+err.Error())
	}

	for _, chapter := range game.Chapters {
		chapter.ID, _ = UUID()
		_, err = tx.Exec(`
			insert into pbe.game_chapters(id, book, chapter, game_id)
			values(?,?,?,?)
		`, chapter.ID, chapter.Book, chapter.Chapter, game.ID)
		if err != nil {
			log.Error("Could not create game chapter: ", err)
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Could not create game chapter: "+err.Error())
		}
	}

	tx.Commit()
	return c.JSON(http.StatusOK, game)
}

func getGamesController(c echo.Context) error {
	games, err := getGames()
	if err != nil {
		log.Error("Could not get games: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get games: "+err.Error())
	}
	return c.JSON(http.StatusOK, games)
}

func getGames() ([]*Game, error) {
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select g.id, g.name, g.seconds, g.created, g.questions, coalesce(g.status, 'OPEN'),
			coalesce(gc.id, ''), coalesce(gc.book, ''), coalesce(gc.chapter, ''), coalesce(t.id, ''), coalesce(t.name, '')
		from pbe.games g
		left join pbe.game_chapters gc on gc.game_id = g.id
		left join pbe.teams t on t.game_id = g.id 
		order by g.created, g.id, t.name, gc.book, gc.chapter
		limit 10
	`)
	if err != nil {
		log.Error("Could not get games: ", err)
		return nil, err
	}

	games := []*Game{}
	game := &Game{}
	team := &Team{}
	for rows.Next() {
		var (
			id        string
			name      string
			seconds   int
			created   string
			questions int
			status    string
			chapterID string
			book      string
			chapter   string
			teamID    string
			teamName  string
		)
		err = rows.Scan(&id, &name, &seconds, &created, &questions, &status, &chapterID, &book, &chapter, &teamID, &teamName)
		if err != nil {
			log.Error("Could not get game: ", err)
			return nil, err
		}

		date, _ := time.Parse("2006-01-02 15:04:05", created)
		if game.ID != id {
			game = &Game{
				ID:        id,
				Name:      name,
				Seconds:   seconds,
				Created:   date,
				Questions: questions,
				Status:    status,
			}
			games = append(games, game)
		}
		if len(book) > 0 {
			found := false
			for _, chp := range game.Chapters {
				if chp.ID == chapterID {
					found = true
					break
				}
			}
			if !found {
				gameChapter := &GameChapter{
					ID:      chapterID,
					Book:    book,
					Chapter: chapter,
				}
				game.Chapters = append(game.Chapters, gameChapter)
			}
		}
		if len(teamID) > 0 && team.ID != teamID {
			team = &Team{
				ID:   teamID,
				Name: teamName,
			}
			game.Teams = append(game.Teams, team)
		}
	}

	return games, nil
}

func getGameController(c echo.Context) error {
	gameID := c.Param("gameID")

	game, err := getGame(gameID)
	if err != nil {
		log.Error("Could not get game: ", gameID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get game: "+gameID+" : "+err.Error())
	}
	return c.JSON(http.StatusOK, game)
}

func getGame(gameID string) (*Game, error) {
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select g.name, g.seconds, g.created, g.questions, coalesce(g.status, 'OPEN'),
			coalesce(gc.id, ''), coalesce(gc.book, ''), coalesce(gc.chapter, ''), coalesce(t.id, ''), coalesce(t.name, '')
		from pbe.games g
		left join pbe.game_chapters gc on gc.game_id = g.id
		left join pbe.teams t on t.game_id = g.id 
		where g.id = ?
		order by t.name, gc.book, gc.chapter 
	`, gameID)
	if err != nil {
		log.Error("Could not get game: "+gameID, " : ", err)
		return nil, err
	}

	game := &Game{
		ID: gameID,
	}
	team := &Team{}
	for rows.Next() {
		var (
			name      string
			seconds   int
			created   string
			questions int
			status    string
			chapterID string
			book      string
			chapter   string
			teamID    string
			teamName  string
		)
		err = rows.Scan(&name, &seconds, &created, &questions, &status, &chapterID, &book, &chapter, &teamID, &teamName)
		if err != nil {
			log.Error("Could not get game: ", err)
			return nil, err
		}

		date, _ := time.Parse("2006-01-02 15:04:05", created)

		if len(game.Name) == 0 {
			game.Name = name
			game.Seconds = seconds
			game.Created = date
			game.Questions = questions
			game.Status = status
		}

		if len(book) > 0 {
			found := false
			for _, chp := range game.Chapters {
				if chp.ID == chapterID {
					found = true
					break
				}
			}
			if !found {
				gameChapter := &GameChapter{
					ID:      chapterID,
					Book:    book,
					Chapter: chapter,
				}
				game.Chapters = append(game.Chapters, gameChapter)
			}
		}

		if len(teamID) > 0 && team.ID != teamID {
			team = &Team{
				ID:   teamID,
				Name: teamName,
			}
			game.Teams = append(game.Teams, team)
		}
	}
	return game, nil
}

func deleteGameController(c echo.Context) error {
	gameID := c.Param("gameID")
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	_, err = conn.Exec(`
		delete from pbe.games where id = ?
	`, gameID)
	if err != nil {
		log.Error("Could not delete game: ", gameID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not delete game: "+gameID+" : "+err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func addTeamController(c echo.Context) error {
	gameID := c.Param("gameID")
	team := &Team{}
	err := c.Bind(&team)
	if err != nil {
		log.Error("Could not create team: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not create team: "+err.Error())
	}

	team.ID, _ = UUID()

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	_, err = conn.Exec(`
		insert into pbe.teams(id, name, game_id)
		values(?,?,?)
	`, team.ID, team.Name, gameID)
	if err != nil {
		log.Error("Could not add team: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not add team: "+err.Error())
	}

	return c.JSON(http.StatusOK, team)
}

func addTeamAnswerController(c echo.Context) error {
	gameID := c.Param("gameID")
	teamID := c.Param("teamID")
	answerID := c.Param("answerID")

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	id, _ := UUID()
	_, err = conn.Exec(`
		insert into pbe.team_answers(id, game_id, team_id, answer_id, created)
		values(?,?,?,?, NOW())
	`, id, gameID, teamID, answerID)
	if err != nil {
		log.Error("Could not add team answer: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not add team answer: "+err.Error())
	}

	return c.JSON(http.StatusOK, id)
}

func deleteTeamAnswerController(c echo.Context) error {
	gameID := c.Param("gameID")
	teamID := c.Param("teamID")
	answerID := c.Param("answerID")

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	_, err = conn.Exec(`
		delete from pbe.team_answers where game_id = ? and team_id = ? and answer_id = ?
	`, gameID, teamID, answerID)
	if err != nil {
		log.Error("Could not delete team answer: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not delete team answer: "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func getTeamController(c echo.Context) error {
	teamID := c.Param("teamID")

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select t.name, a.answer, a.status, coalesce(a.id, ''), coalesce(ta.id, '')
		from pbe.team_answers ta
		inner join pbe.teams t on t.id = ta.team_id
		left join pbe.answers a on a.id = ta.answer_id
		where t.id = ?
	`, teamID)
	if err != nil {
		log.Error("Could not get team: ", err)
		return c.JSON(http.StatusOK, "Could not get team: "+err.Error())
	}

	team := &Team{
		ID: teamID,
	}
	for rows.Next() {
		var (
			name         string
			answer       string
			status       bool
			answerID     string
			teamAnswerID string
		)

		err = rows.Scan(&name, &answer, &status, &answerID, &teamAnswerID)
		if err != nil {
			log.Error("Could not get team: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get team: "+err.Error())
		}

		if len(team.Name) == 0 {
			team.Name = name
		}

		if len(answer) > 0 {
			a := &Answer{
				ID:           answerID,
				TeamAnswerID: teamAnswerID,
				Answer:       answer,
				Status:       status,
			}
			team.Answers = append(team.Answers, a)
		}
	}

	return c.JSON(http.StatusOK, team)
}

func getHomeTeamController(c echo.Context) error {
	log.Info("Getting home team!!!")
	gameID := c.Param("gameID")

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select t.id, coalesce(a.answer, ''), coalesce(status, false), coalesce(a.id, ''), coalesce(ta.id, '')
		from pbe.teams t
		left join pbe.team_answers ta on t.id = ta.team_id
		left join pbe.answers a on a.id = ta.answer_id
		where t.name = 'Home' and t.game_id = ?
	`, gameID)
	if err != nil {
		log.Error("Could not get home team: ", err)
		return c.JSON(http.StatusOK, "Could not get home team: "+err.Error())
	}

	team := &Team{
		Name: "Home",
	}
	for rows.Next() {
		var (
			id           string
			answer       string
			status       bool
			answerID     string
			teamAnswerID string
		)

		err = rows.Scan(&id, &answer, &status, &answerID, &teamAnswerID)
		if err != nil {
			log.Error("Could not get home team: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get home team: "+err.Error())
		}

		log.Info("TeamID: ", team.ID, " : ", len(team.ID))
		if len(team.ID) == 0 {
			log.Info("Setting teamID: ", id)
			team.ID = id
		}

		if len(answer) > 0 {
			a := &Answer{
				ID:           answerID,
				TeamAnswerID: teamAnswerID,
				Answer:       answer,
				Status:       status,
			}
			team.Answers = append(team.Answers, a)
		}
	}

	return c.JSON(http.StatusOK, team)
}

func startGameController(c echo.Context) error {
	gameID := c.Param("gameID")
	game, err := getGame(gameID)
	if err != nil {
		log.Error("Could not get game: ", gameID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get game: "+gameID+" : "+err.Error())
	}

	if game.Status != "OPEN" {
		log.Error("Cannot start a game that is not open")
		return c.JSON(http.StatusInternalServerError, "Cannot start a game that is not open")
	}

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	questions := []*Question{}
	for _, chapter := range game.Chapters {
		rows, err := conn.Query(`
			select id
			from pbe.questions
			where book = ? and chapter = ?
			order by verses
		`, chapter.Book, chapter.Chapter)
		if err != nil {
			log.Error("Could not get questions: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get questions: "+err.Error())
		}

		for rows.Next() {
			var (
				id string
			)

			err = rows.Scan(&id)
			if err != nil {
				log.Error("Could not get question: ", err)
				return c.JSON(http.StatusInternalServerError, "Could not get question: "+err.Error())
			}

			question := &Question{
				ID: id,
			}
			questions = append(questions, question)
		}
	}

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not start database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not start database transaction: "+err.Error())
	}

	_, err = tx.Exec(`
		update pbe.games set status = ? where id = ?
	`, "STARTED", gameID)
	if err != nil {
		log.Error("Could not start game: ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not start game: "+err.Error())
	}

	for pos, question := range questions {
		gameQuestionID, _ := UUID()
		if pos == 0 {
			_, err = tx.Exec(`
				update pbe.games set question = ? where id = ?
			`, gameQuestionID, gameID)
			if err != nil {
				log.Error("Could not set current game question: ", err)
				tx.Rollback()
				return c.JSON(http.StatusInternalServerError, "Could not set current game question: "+err.Error())
			}
		}
		_, err = tx.Exec(`
			insert into pbe.game_questions(id, game_id, question_id, position)
			values(?,?,?,?)
		`, gameQuestionID, gameID, question.ID, pos)
		if err != nil {
			log.Error("Could not insert game question: ", err)
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Could not insert game question: "+err.Error())
		}
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

func finishGameController(c echo.Context) error {
	gameID := c.Param("gameID")
	game, err := getGame(gameID)
	if err != nil {
		log.Error("Could not get game: ", gameID, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get game: "+gameID+" : "+err.Error())
	}

	if game.Status != "STARTED" {
		log.Error("Cannot finish a game that is not started")
		return c.JSON(http.StatusInternalServerError, "Cannot finish a game that is not started")
	}

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not start database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not start database transaction: "+err.Error())
	}

	_, err = tx.Exec(`
		update pbe.games set status = ? where id = ?
	`, "FINISHED", gameID)
	if err != nil {
		log.Error("Could not finish game: ", err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

func getCurrentQuestionController(c echo.Context) error {
	gameID := c.Param("gameID")

	question, err := getCurrentQuestion(gameID)
	if err != nil {
		log.Error("Could not get current question: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get current question: "+err.Error())
	}
	return c.JSON(http.StatusOK, question)
}

func getCurrentQuestion(gameID string) (*Question, error) {
	log.Info("Getting current question")
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return nil, err
	}
	defer conn.Close()

	var (
		questionID string
		status     string
	)
	err = conn.QueryRow(`
		select status
		from pbe.games
		where id = ?
	`, gameID).Scan(&status)
	if err != nil {
		log.Error("Could not get game status: ", err)
		return nil, err
	}

	if status == "FINISHED" {
		question := &Question{
			Finished: true,
		}

		return question, nil
	}

	err = conn.QueryRow(`
		select gq.question_id
		from pbe.game_questions gq
		inner join pbe.games g on g.question = gq.id
		where gq.game_id = ?
	`, gameID).Scan(&questionID)
	if err != nil {
		log.Error("Could not get current question: ", err)
		question := &Question{
			Question: "Waiting for game to start",
		}

		return question, nil
	}

	question, err := getQuestion(questionID)
	if err != nil {
		log.Error("Could not get current question: ", err)
		return nil, err
	}

	return question, nil
}

func nextQuestionController(c echo.Context) error {
	gameID := c.Param("gameID")
	var (
		nextGameQuestionID string
		position           int
		questionID         string
	)

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not start database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not start database transaction: "+err.Error())
	}
	err = tx.QueryRow(`
		select gq.position
		from pbe.game_questions gq
		inner join pbe.games g on g.question = gq.id
		where gq.game_id = ?
		and g.status = 'STARTED'
	`, gameID).Scan(&position)
	if err != nil {
		log.Error("Could not get the current question position for game: ", gameID, " : ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not get the current question position for game: "+gameID+" : "+err.Error())
	}

	log.Info("Got position: ", position)

	position++

	log.Info("Looking for position: ", position, " : ", gameID)

	err = tx.QueryRow(`
		select gq.id, gq.question_id
		from pbe.game_questions gq
		inner join pbe.games g on g.id = gq.game_id
		where gq.game_id = ?
		and gq.position = ?
	`, gameID, position).Scan(&nextGameQuestionID, &questionID)
	if err != nil {
		log.Error("Could not set next question: ", err)
		_, err = tx.Exec(`
			update pbe.games set status = ? where id = ?
		`, "FINISHED", gameID)
		if err != nil {
			log.Error("Could not finish game: ", err)
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`
		update pbe.games set question = ? where id = ?
	`, nextGameQuestionID, gameID)
	if err != nil {
		log.Error("Could not set new current question: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not not set new current question: "+err.Error())
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

func previousQuestionController(c echo.Context) error {
	gameID := c.Param("gameID")
	var (
		previousGameQuestionID string
		position               int
		questionID             string
	)

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not start database transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not start database transaction: "+err.Error())
	}
	err = tx.QueryRow(`
		select gq.position
		from pbe.game_questions gq
		inner join pbe.games g on g.question = gq.id
		where gq.game_id = ?
		and g.status = 'STARTED'
	`, gameID).Scan(&position)
	if err != nil {
		log.Error("Could not get the current question position for game: ", gameID, " : ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not get the current question position for game: "+gameID+" : "+err.Error())
	}

	if position == 0 {
		// Already at position 0
		tx.Commit()
		return c.NoContent(http.StatusOK)
	}

	position--

	err = tx.QueryRow(`
		select gq.id, gq.question_id
		from pbe.game_questions gq
		inner join pbe.games g on g.id = gq.game_id
		where gq.game_id = ?
		and gq.position = ?
	`, gameID, position).Scan(&previousGameQuestionID, &questionID)
	if err != nil {
		log.Error("Could not get previous question: ", gameID, " : ", position, " : ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get previous question: "+err.Error())
	}

	_, err = tx.Exec(`
		update pbe.games set question = ? where id = ?
	`, previousGameQuestionID, gameID)
	if err != nil {
		log.Error("Could not set new current question: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not not set new current question: "+err.Error())
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

func getFinishedGameController(c echo.Context) error {
	gameID := c.Param("gameID")
	game, err := getGame(gameID)

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not connect to database: "+err.Error())
	}
	defer conn.Close()

	// rows, err := conn.Query(`
	// 	select
	// 		q.id
	// 		, q.book
	// 		, q.chapter
	// 		, q.verses
	// 		, q.question
	// 		, COALESCE(a.id, '')
	// 		, COALESCE(a.answer, '')
	// 		, COALESCE(a.status, false)
	// 	from pbe.questions q
	// 	inner join pbe.game_questions gq on gq.question_id = q.id
	// 	left join pbe.answers a on a.question_id = q.id
	// 	where gq.game_id = ?
	// 	order by q.id, a.answer
	// `, gameID)
	// if err != nil {
	// 	log.Error("Could not get questions: ", err)
	// 	return c.JSON(http.StatusInternalServerError, "Could not get questions: "+err.Error())
	// }

	// question := &Question{}
	// for rows.Next() {
	// 	var (
	// 		id           string
	// 		book         string
	// 		chapter      string
	// 		verses       string
	// 		questionText string
	// 		answerID     string
	// 		answer       string
	// 		status       bool
	// 	)
	// 	err = rows.Scan(&id, &book, &chapter, &verses, &questionText, &answerID, &answer, &status)
	// 	if err != nil {
	// 		log.Error("Could not get question: ", err)
	// 		return c.JSON(http.StatusInternalServerError, "Could not get question: "+err.Error())
	// 	}

	// 	if question.ID != id {
	// 		question := &Question{
	// 			ID:       id,
	// 			Book:     book,
	// 			Chapter:  chapter,
	// 			Verses:   verses,
	// 			Question: questionText,
	// 		}
	// 		game.Questions2 = append(game.Questions2, question)
	// 	}
	// 	if len(answerID) > 0 {
	// 		a := &Answer{
	// 			ID:     answerID,
	// 			Answer: answer,
	// 			Status: status,
	// 		}
	// 		question.Answers = append(question.Answers, a)
	// 	}
	// }

	for _, team := range game.Teams {
		log.Info("Getting results for: ", team.Name, " : ", team.ID)
		rows, err := conn.Query(`
			select a.answer, a.status, a.id, ta.id
			from pbe.team_answers ta
			inner join pbe.teams t on t.id = ta.team_id
			inner join pbe.answers a on a.id = ta.answer_id
			where t.id = ?
		`, team.ID)
		if err != nil {
			log.Error("Could not get team: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get team: "+err.Error())
		}

		for rows.Next() {
			var (
				answer       string
				status       bool
				answerID     string
				teamAnswerID string
			)

			err = rows.Scan(&answer, &status, &answerID, &teamAnswerID)
			if err != nil {
				log.Error("Could not get team answer: ", err)
				return c.JSON(http.StatusInternalServerError, "Could not get team answer: "+err.Error())
			}

			log.Info("Status: ", status)
			if len(answer) > 0 {
				a := &Answer{
					ID:           answerID,
					TeamAnswerID: teamAnswerID,
					Answer:       answer,
					Status:       status,
				}
				if a.Status {
					team.Points++
				} else if team.Points > 0 {
					team.Points--
				}
				team.Answers = append(team.Answers, a)
			}
		}
	}

	return c.JSON(http.StatusOK, game)
}
